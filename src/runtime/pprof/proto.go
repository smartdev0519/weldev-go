// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pprof

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
	"unsafe"

	"internal/pprof/profile"
)

// lostProfileEvent is the function to which lost profiling
// events are attributed.
// (The name shows up in the pprof graphs.)
func lostProfileEvent() { lostProfileEvent() }

// funcPC returns the PC for the func value f.
func funcPC(f interface{}) uintptr {
	return *(*[2]*uintptr)(unsafe.Pointer(&f))[1]
}

// A profileBuilder builds a profile.Profile incrementally from a
// stream of profile samples delivered by the runtime.
// TODO(rsc,matloob): In the long term, we'd like to avoid
// storing the entire profile.Profile in memory, instead streaming
// the encoded form out to an underlying writer.
// Even so, this one copy is a step forward from Go 1.8,
// which had two full copies of the data in memory.
type profileBuilder struct {
	p          *profile.Profile
	start      time.Time
	havePeriod bool
	locs       map[uintptr]*profile.Location
	m          profMap
}

// newProfileBuilder returns a new profileBuilder.
// CPU profiling data obtained from the runtime can be added
// by calling b.addCPUData, and then the eventual profile
// can be obtained by calling b.finish.
func newProfileBuilder() *profileBuilder {
	start := time.Now()
	p := &profile.Profile{
		PeriodType: &profile.ValueType{Type: "cpu", Unit: "nanoseconds"},
		SampleType: []*profile.ValueType{
			{Type: "samples", Unit: "count"},
			{Type: "cpu", Unit: "nanoseconds"},
		},
		TimeNanos: int64(start.UnixNano()),
	}
	return &profileBuilder{
		p:     p,
		start: start,
		locs:  make(map[uintptr]*profile.Location),
	}
}

// addCPUData adds the CPU profiling data to the profile.
// The data must be a whole number of records,
// as delivered by the runtime.
func (b *profileBuilder) addCPUData(data []uint64, tags []unsafe.Pointer) error {
	p := b.p
	if !b.havePeriod {
		// first record is period
		if len(data) < 3 {
			return fmt.Errorf("truncated profile")
		}
		if data[0] != 3 || data[2] == 0 {
			return fmt.Errorf("malformed profile")
		}
		period := int64(data[2])
		p.Period = period * 1000
		data = data[3:]
		b.havePeriod = true
	}

	// Parse CPU samples from the profile.
	// Each sample is 3+n uint64s:
	//	data[0] = 3+n
	//	data[1] = time stamp (ignored)
	//	data[2] = count
	//	data[3:3+n] = stack
	// If the count is 0 and the stack has length 1,
	// that's an overflow record inserted by the runtime
	// to indicate that stack[0] samples were lost.
	// Otherwise the count is usually 1,
	// but in a few special cases like lost non-Go samples
	// there can be larger counts.
	// Because many samples with the same stack arrive,
	// we want to deduplicate immediately, which we do
	// using the b.m profMap.
	for len(data) > 0 {
		if len(data) < 3 || data[0] > uint64(len(data)) {
			return fmt.Errorf("truncated profile")
		}
		if data[0] < 3 || tags != nil && len(tags) < 1 {
			return fmt.Errorf("malformed profile")
		}
		count := data[2]
		stk := data[3:data[0]]
		data = data[data[0]:]
		var tag unsafe.Pointer
		if tags != nil {
			tag = tags[0]
			tags = tags[1:]
		}

		if count == 0 && len(stk) == 1 {
			// overflow record
			count = uint64(stk[0])
			stk = []uint64{
				uint64(funcPC(lostProfileEvent)),
			}
		}
		b.m.lookup(stk, tag).count += int64(count)
	}
	return nil
}

// build completes and returns the constructed profile.
func (b *profileBuilder) build() *profile.Profile {
	b.p.DurationNanos = time.Since(b.start).Nanoseconds()

	for e := b.m.all; e != nil; e = e.nextAll {
		s := &profile.Sample{
			Value:    []int64{e.count, e.count * int64(b.p.Period)},
			Location: make([]*profile.Location, len(e.stk)),
		}
		for i, addr := range e.stk {
			addr := uintptr(addr)
			// Addresses from stack traces point to the next instruction after
			// each call.  Adjust by -1 to land somewhere on the actual call
			// (except for the leaf, which is not a call).
			if i > 0 {
				addr--
			}
			loc := b.locs[addr]
			if loc == nil {
				loc = &profile.Location{
					ID:      uint64(len(b.p.Location) + 1),
					Address: uint64(addr),
				}
				b.locs[addr] = loc
				b.p.Location = append(b.p.Location, loc)
			}
			s.Location[i] = loc
		}
		b.p.Sample = append(b.p.Sample, s)
	}

	if runtime.GOOS == "linux" {
		addMappings(b.p)
	}
	symbolize(b.p)
	return b.p
}

// addMappings adds information from /proc/self/maps
// to the profile if possible.
func addMappings(p *profile.Profile) {
	// Parse memory map from /proc/self/maps
	f, err := os.Open("/proc/self/maps")
	if err != nil {
		return
	}
	p.ParseMemoryMap(f)
	f.Close()
}

type function interface {
	Name() string
	FileLine(pc uintptr) (string, int)
}

// funcForPC is a wrapper for runtime.FuncForPC. Defined as var for testing.
var funcForPC = func(pc uintptr) function {
	if f := runtime.FuncForPC(pc); f != nil {
		return f
	}
	return nil
}

func symbolize(p *profile.Profile) {
	fns := profileFunctionMap{}
	for _, l := range p.Location {
		pc := uintptr(l.Address)
		f := funcForPC(pc)
		if f == nil {
			continue
		}
		file, lineno := f.FileLine(pc)
		l.Line = []profile.Line{
			{
				Function: fns.findOrAddFunction(f.Name(), file, p),
				Line:     int64(lineno),
			},
		}
	}
	// Trim runtime functions. Always hide runtime.goexit. Other runtime
	// functions are only hidden for heapz when they appear at the beginning.
	isHeapz := p.PeriodType != nil && p.PeriodType.Type == "space"
	for _, s := range p.Sample {
		show := !isHeapz
		var i int
		for _, l := range s.Location {
			if len(l.Line) > 0 && l.Line[0].Function != nil {
				name := l.Line[0].Function.Name
				if name == "runtime.goexit" || !show && strings.HasPrefix(name, "runtime.") {
					continue
				}
			}
			show = true
			s.Location[i] = l
			i++
		}
		s.Location = s.Location[:i]
	}
}

type profileFunctionMap map[profile.Function]*profile.Function

func (fns profileFunctionMap) findOrAddFunction(name, filename string, p *profile.Profile) *profile.Function {
	f := profile.Function{
		Name:       name,
		SystemName: name,
		Filename:   filename,
	}
	if fp := fns[f]; fp != nil {
		return fp
	}
	fp := new(profile.Function)
	fns[f] = fp

	*fp = f
	fp.ID = uint64(len(p.Function) + 1)
	p.Function = append(p.Function, fp)
	return fp
}
