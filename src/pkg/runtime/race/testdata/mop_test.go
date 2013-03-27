// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package race_test

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type Point struct {
	x, y int
}

type NamedPoint struct {
	name string
	p    Point
}

type DummyWriter struct {
	state int
}
type Writer interface {
	Write(p []byte) (n int)
}

func (d DummyWriter) Write(p []byte) (n int) {
	return 0
}

var GlobalX, GlobalY int = 0, 0
var GlobalCh chan int = make(chan int, 2)

func GlobalFunc1() {
	GlobalY = GlobalX
	GlobalCh <- 1
}

func GlobalFunc2() {
	GlobalX = 1
	GlobalCh <- 1
}

func TestRaceIntRWGlobalFuncs(t *testing.T) {
	go GlobalFunc1()
	go GlobalFunc2()
	<-GlobalCh
	<-GlobalCh
}

func TestRaceIntRWClosures(t *testing.T) {
	var x, y int
	ch := make(chan int, 2)

	go func() {
		y = x
		ch <- 1
	}()
	go func() {
		x = 1
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestNoRaceIntRWClosures(t *testing.T) {
	var x, y int
	ch := make(chan int, 1)

	go func() {
		y = x
		ch <- 1
	}()
	<-ch
	go func() {
		x = 1
		ch <- 1
	}()
	<-ch

}

func TestRaceInt32RWClosures(t *testing.T) {
	var x, y int32
	ch := make(chan bool, 2)

	go func() {
		y = x
		ch <- true
	}()
	go func() {
		x = 1
		ch <- true
	}()
	<-ch
	<-ch
}

func TestNoRaceCase(t *testing.T) {
	var y int
	for x := -1; x <= 1; x++ {
		switch {
		case x < 0:
			y = -1
		case x == 0:
			y = 0
		case x > 0:
			y = 1
		}
	}
	y++
}

func TestRaceCaseCondition(t *testing.T) {
	var x int = 0
	ch := make(chan int, 2)

	go func() {
		x = 2
		ch <- 1
	}()
	go func() {
		switch x < 2 {
		case true:
			x = 1
			//case false:
			//	x = 5
		}
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceCaseCondition2(t *testing.T) {
	// switch body is rearranged by the compiler so the tests
	// passes even if we don't instrument '<'
	var x int = 0
	ch := make(chan int, 2)

	go func() {
		x = 2
		ch <- 1
	}()
	go func() {
		switch x < 2 {
		case true:
			x = 1
		case false:
			x = 5
		}
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceCaseBody(t *testing.T) {
	var x, y int
	ch := make(chan int, 2)

	go func() {
		y = x
		ch <- 1
	}()
	go func() {
		switch {
		default:
			x = 1
		case x == 100:
			x = -x
		}
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestNoRaceCaseFallthrough(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)
	z = 1

	go func() {
		y = x
		ch <- 1
	}()
	go func() {
		switch {
		case z == 1:
		case z == 2:
			x = 2
		}
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceCaseFallthrough(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)
	z = 1

	go func() {
		y = x
		ch <- 1
	}()
	go func() {
		switch {
		case z == 1:
			fallthrough
		case z == 2:
			x = 2
		}
		ch <- 1
	}()

	<-ch
	<-ch
}

func TestRaceCaseType(t *testing.T) {
	var x, y int
	var i interface{} = x
	c := make(chan int, 1)
	go func() {
		switch i.(type) {
		case nil:
		case int:
		}
		c <- 1
	}()
	i = y
	<-c
}

func TestRaceCaseTypeBody(t *testing.T) {
	var x, y int
	var i interface{} = &x
	c := make(chan int, 1)
	go func() {
		switch i := i.(type) {
		case nil:
		case *int:
			*i = y
		}
		c <- 1
	}()
	x = y
	<-c
}

func TestNoRaceRange(t *testing.T) {
	ch := make(chan int, 3)
	a := [...]int{1, 2, 3}
	for _, v := range a {
		ch <- v
	}
	close(ch)
}

func TestRaceRange(t *testing.T) {
	const N = 2
	var a [N]int
	var x, y int
	done := make(chan bool, N)
	for i, v := range a {
		go func(i int) {
			// we don't want a write-vs-write race
			// so there is no array b here
			if i == 0 {
				x = v
			} else {
				y = v
			}
			done <- true
		}(i)
	}
	for i := 0; i < N; i++ {
		<-done
	}
}

func TestRacePlus(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)

	go func() {
		y = x + z
		ch <- 1
	}()
	go func() {
		y = x + z + z
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRacePlus2(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)

	go func() {
		x = 1
		ch <- 1
	}()
	go func() {
		y = +x + z
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestNoRacePlus(t *testing.T) {
	var x, y, z, f int
	ch := make(chan int, 2)

	go func() {
		y = x + z
		ch <- 1
	}()
	go func() {
		f = z + x
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceComplement(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)

	go func() {
		x = ^y
		ch <- 1
	}()
	go func() {
		y = ^z
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceDiv(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)

	go func() {
		x = y / (z + 1)
		ch <- 1
	}()
	go func() {
		y = z
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceDivConst(t *testing.T) {
	var x, y, z uint32
	ch := make(chan int, 2)

	go func() {
		x = y / 3 // involves only a HMUL node
		ch <- 1
	}()
	go func() {
		y = z
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceMod(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)

	go func() {
		x = y % (z + 1)
		ch <- 1
	}()
	go func() {
		y = z
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceModConst(t *testing.T) {
	var x, y, z int
	ch := make(chan int, 2)

	go func() {
		x = y % 3
		ch <- 1
	}()
	go func() {
		y = z
		ch <- 1
	}()
	<-ch
	<-ch
}

func TestRaceRotate(t *testing.T) {
	var x, y, z uint32
	ch := make(chan int, 2)

	go func() {
		x = y<<12 | y>>20
		ch <- 1
	}()
	go func() {
		y = z
		ch <- 1
	}()
	<-ch
	<-ch
}

// May crash if the instrumentation is reckless.
func TestNoRaceEnoughRegisters(t *testing.T) {
	// from erf.go
	const (
		sa1 = 1
		sa2 = 2
		sa3 = 3
		sa4 = 4
		sa5 = 5
		sa6 = 6
		sa7 = 7
		sa8 = 8
	)
	var s, S float64
	s = 3.1415
	S = 1 + s*(sa1+s*(sa2+s*(sa3+s*(sa4+s*(sa5+s*(sa6+s*(sa7+s*sa8)))))))
	s = S
}

// emptyFunc should not be inlined.
func emptyFunc(x int) {
	if false {
		fmt.Println(x)
	}
}

func TestRaceFuncArgument(t *testing.T) {
	var x int
	ch := make(chan bool, 1)
	go func() {
		emptyFunc(x)
		ch <- true
	}()
	x = 1
	<-ch
}

func TestRaceFuncArgument2(t *testing.T) {
	var x int
	ch := make(chan bool, 2)
	go func() {
		x = 42
		ch <- true
	}()
	go func(y int) {
		ch <- true
	}(x)
	<-ch
	<-ch
}

func TestRaceSprint(t *testing.T) {
	var x int
	ch := make(chan bool, 1)
	go func() {
		fmt.Sprint(x)
		ch <- true
	}()
	x = 1
	<-ch
}

// Not implemented.
func TestRaceFailingArrayCopy(t *testing.T) {
	ch := make(chan bool, 1)
	var a [5]int
	go func() {
		a[3] = 1
		ch <- true
	}()
	a = [5]int{1, 2, 3, 4, 5}
	<-ch
}

func TestRaceStructRW(t *testing.T) {
	p := Point{0, 0}
	ch := make(chan bool, 1)
	go func() {
		p = Point{1, 1}
		ch <- true
	}()
	q := p
	<-ch
	p = q
}

func TestRaceStructFieldRW1(t *testing.T) {
	p := Point{0, 0}
	ch := make(chan bool, 1)
	go func() {
		p.x = 1
		ch <- true
	}()
	_ = p.x
	<-ch
}

func TestNoRaceStructFieldRW1(t *testing.T) {
	// Same struct, different variables, no
	// pointers. The layout is known (at compile time?) ->
	// no read on p
	// writes on x and y
	p := Point{0, 0}
	ch := make(chan bool, 1)
	go func() {
		p.x = 1
		ch <- true
	}()
	p.y = 1
	<-ch
	_ = p
}

func TestNoRaceStructFieldRW2(t *testing.T) {
	// Same as NoRaceStructFieldRW1
	// but p is a pointer, so there is a read on p
	p := Point{0, 0}
	ch := make(chan bool, 1)
	go func() {
		p.x = 1
		ch <- true
	}()
	p.y = 1
	<-ch
	_ = p
}

func TestRaceStructFieldRW2(t *testing.T) {
	p := &Point{0, 0}
	ch := make(chan bool, 1)
	go func() {
		p.x = 1
		ch <- true
	}()
	_ = p.x
	<-ch
}

func TestRaceStructFieldRW3(t *testing.T) {
	p := NamedPoint{name: "a", p: Point{0, 0}}
	ch := make(chan bool, 1)
	go func() {
		p.p.x = 1
		ch <- true
	}()
	_ = p.p.x
	<-ch
}

func TestRaceEfaceWW(t *testing.T) {
	var a, b interface{}
	ch := make(chan bool, 1)
	go func() {
		a = 1
		ch <- true
	}()
	a = 2
	<-ch
	_, _ = a, b
}

func TestRaceIfaceWW(t *testing.T) {
	var a, b Writer
	ch := make(chan bool, 1)
	go func() {
		a = DummyWriter{1}
		ch <- true
	}()
	a = DummyWriter{2}
	<-ch
	b = a
	a = b
}

func TestRaceIfaceCmp(t *testing.T) {
	var a, b Writer
	a = DummyWriter{1}
	ch := make(chan bool, 1)
	go func() {
		a = DummyWriter{1}
		ch <- true
	}()
	_ = a == b
	<-ch
}

func TestRaceIfaceCmpNil(t *testing.T) {
	var a Writer
	a = DummyWriter{1}
	ch := make(chan bool, 1)
	go func() {
		a = DummyWriter{1}
		ch <- true
	}()
	_ = a == nil
	<-ch
}

func TestRaceEfaceConv(t *testing.T) {
	c := make(chan bool)
	v := 0
	go func() {
		go func(x interface{}) {
		}(v)
		c <- true
	}()
	v = 42
	<-c
}

type OsFile struct{}

func (*OsFile) Read() {
}

type IoReader interface {
	Read()
}

func TestRaceIfaceConv(t *testing.T) {
	c := make(chan bool)
	f := &OsFile{}
	go func() {
		go func(x IoReader) {
		}(f)
		c <- true
	}()
	f = &OsFile{}
	<-c
}

func TestRaceError(t *testing.T) {
	ch := make(chan bool, 1)
	var err error
	go func() {
		err = nil
		ch <- true
	}()
	_ = err
	<-ch
}

func TestRaceIntptrRW(t *testing.T) {
	var x, y int
	var p *int = &x
	ch := make(chan bool, 1)
	go func() {
		*p = 5
		ch <- true
	}()
	y = *p
	x = y
	<-ch
}

func TestRaceStringRW(t *testing.T) {
	ch := make(chan bool, 1)
	s := ""
	go func() {
		s = "abacaba"
		ch <- true
	}()
	_ = s
	<-ch
}

func TestRaceStringPtrRW(t *testing.T) {
	ch := make(chan bool, 1)
	var x string
	p := &x
	go func() {
		*p = "a"
		ch <- true
	}()
	_ = *p
	<-ch
}

func TestRaceFloat64WW(t *testing.T) {
	var x, y float64
	ch := make(chan bool, 1)
	go func() {
		x = 1.0
		ch <- true
	}()
	x = 2.0
	<-ch

	y = x
	x = y
}

func TestRaceComplex128WW(t *testing.T) {
	var x, y complex128
	ch := make(chan bool, 1)
	go func() {
		x = 2 + 2i
		ch <- true
	}()
	x = 4 + 4i
	<-ch

	y = x
	x = y
}

func TestRaceUnsafePtrRW(t *testing.T) {
	var x, y, z int
	x, y, z = 1, 2, 3
	var p unsafe.Pointer = unsafe.Pointer(&x)
	ch := make(chan bool, 1)
	go func() {
		p = (unsafe.Pointer)(&z)
		ch <- true
	}()
	y = *(*int)(p)
	x = y
	<-ch
}

func TestRaceFuncVariableRW(t *testing.T) {
	var f func(x int) int
	f = func(x int) int {
		return x * x
	}
	ch := make(chan bool, 1)
	go func() {
		f = func(x int) int {
			return x
		}
		ch <- true
	}()
	y := f(1)
	<-ch
	x := y
	y = x
}

func TestRaceFuncVariableWW(t *testing.T) {
	var f func(x int) int
	ch := make(chan bool, 1)
	go func() {
		f = func(x int) int {
			return x
		}
		ch <- true
	}()
	f = func(x int) int {
		return x * x
	}
	<-ch
}

// This one should not belong to mop_test
func TestRacePanic(t *testing.T) {
	var x int
	var zero int = 0
	ch := make(chan bool, 2)
	go func() {
		defer func() {
			err := recover()
			if err == nil {
				panic("should be panicking")
			}
			x = 1
			ch <- true
		}()
		var y int = 1 / zero
		zero = y
	}()
	go func() {
		defer func() {
			err := recover()
			if err == nil {
				panic("should be panicking")
			}
			x = 2
			ch <- true
		}()
		var y int = 1 / zero
		zero = y
	}()

	<-ch
	<-ch
	if zero != 0 {
		panic("zero has changed")
	}
}

func TestNoRaceBlank(t *testing.T) {
	var a [5]int
	ch := make(chan bool, 1)
	go func() {
		_, _ = a[0], a[1]
		ch <- true
	}()
	_, _ = a[2], a[3]
	<-ch
	a[1] = a[0]
}

func TestRaceAppendRW(t *testing.T) {
	a := make([]int, 10)
	ch := make(chan bool)
	go func() {
		_ = append(a, 1)
		ch <- true
	}()
	a[0] = 1
	<-ch
}

func TestRaceAppendLenRW(t *testing.T) {
	a := make([]int, 0)
	ch := make(chan bool)
	go func() {
		a = append(a, 1)
		ch <- true
	}()
	_ = len(a)
	<-ch
}

func TestRaceAppendCapRW(t *testing.T) {
	a := make([]int, 0)
	ch := make(chan string)
	go func() {
		a = append(a, 1)
		ch <- ""
	}()
	_ = cap(a)
	<-ch
}

func TestNoRaceFuncArgsRW(t *testing.T) {
	ch := make(chan byte, 1)
	var x byte
	go func(y byte) {
		_ = y
		ch <- 0
	}(x)
	x = 1
	<-ch
}

func TestRaceFuncArgsRW(t *testing.T) {
	ch := make(chan byte, 1)
	var x byte
	go func(y *byte) {
		_ = *y
		ch <- 0
	}(&x)
	x = 1
	<-ch
}

// from the mailing list, slightly modified
// unprotected concurrent access to seen[]
func TestRaceCrawl(t *testing.T) {
	url := "dummyurl"
	depth := 3
	seen := make(map[string]bool)
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	var crawl func(string, int)
	crawl = func(u string, d int) {
		nurl := 0
		defer func() {
			ch <- nurl
		}()
		seen[u] = true
		if d <= 0 {
			return
		}
		urls := [...]string{"a", "b", "c"}
		for _, uu := range urls {
			if _, ok := seen[uu]; !ok {
				wg.Add(1)
				go crawl(uu, d-1)
				nurl++
			}
		}
		wg.Done()
	}
	wg.Add(1)
	go crawl(url, depth)
	wg.Wait()
}

func TestRaceIndirection(t *testing.T) {
	ch := make(chan struct{}, 1)
	var y int
	var x *int = &y
	go func() {
		*x = 1
		ch <- struct{}{}
	}()
	*x = 2
	<-ch
	_ = *x
}

func TestRaceRune(t *testing.T) {
	c := make(chan bool)
	var x rune
	go func() {
		x = 1
		c <- true
	}()
	_ = x
	<-c
}

func TestRaceEmptyInterface1(t *testing.T) {
	c := make(chan bool)
	var x interface{}
	go func() {
		x = nil
		c <- true
	}()
	_ = x
	<-c
}

func TestRaceEmptyInterface2(t *testing.T) {
	c := make(chan bool)
	var x interface{}
	go func() {
		x = &Point{}
		c <- true
	}()
	_ = x
	<-c
}

func TestRaceTLS(t *testing.T) {
	comm := make(chan *int)
	done := make(chan bool, 2)
	go func() {
		var x int
		comm <- &x
		x = 1
		x = *(<-comm)
		done <- true
	}()
	go func() {
		p := <-comm
		*p = 2
		comm <- p
		done <- true
	}()
	<-done
	<-done
}

func TestNoRaceHeapReallocation(t *testing.T) {
	// It is possible that a future implementation
	// of memory allocation will ruin this test.
	// Increasing n might help in this case, so
	// this test is a bit more generic than most of the
	// others.
	const n = 2
	done := make(chan bool, n)
	empty := func(p *int) {}
	for i := 0; i < n; i++ {
		ms := i
		go func() {
			<-time.After(time.Duration(ms) * time.Millisecond)
			runtime.GC()
			var x int
			empty(&x) // x goes to the heap
			done <- true
		}()
	}
	for i := 0; i < n; i++ {
		<-done
	}
}

func TestRaceAnd(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		x = 1
		c <- true
	}()
	if x == 1 && y == 1 {
	}
	<-c
}

func TestRaceAnd2(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		x = 1
		c <- true
	}()
	if y == 0 && x == 1 {
	}
	<-c
}

func TestNoRaceAnd(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		x = 1
		c <- true
	}()
	if y == 1 && x == 1 {
	}
	<-c
}

func TestRaceOr(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		x = 1
		c <- true
	}()
	if x == 1 || y == 1 {
	}
	<-c
}

func TestRaceOr2(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		x = 1
		c <- true
	}()
	if y == 1 || x == 1 {
	}
	<-c
}

func TestNoRaceOr(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		x = 1
		c <- true
	}()
	if y == 0 || x == 1 {
	}
	<-c
}

func TestNoRaceShortCalc(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		y = 1
		c <- true
	}()
	if x == 0 || y == 0 {
	}
	<-c
}

func TestNoRaceShortCalc2(t *testing.T) {
	c := make(chan bool)
	x, y := 0, 0
	go func() {
		y = 1
		c <- true
	}()
	if x == 1 && y == 0 {
	}
	<-c
}

func TestRaceFuncItself(t *testing.T) {
	c := make(chan bool)
	f := func() {}
	go func() {
		f()
		c <- true
	}()
	f = func() {}
	<-c
}

func TestNoRaceFuncUnlock(t *testing.T) {
	ch := make(chan bool, 1)
	var mu sync.Mutex
	x := 0
	go func() {
		mu.Lock()
		x = 42
		mu.Unlock()
		ch <- true
	}()
	x = func(mu *sync.Mutex) int {
		mu.Lock()
		return 43
	}(&mu)
	mu.Unlock()
	<-ch
}

func TestRaceStructInit(t *testing.T) {
	type X struct {
		x, y int
	}
	c := make(chan bool, 1)
	y := 0
	go func() {
		y = 42
		c <- true
	}()
	x := X{x: y}
	_ = x
	<-c
}

func TestRaceArrayInit(t *testing.T) {
	c := make(chan bool, 1)
	y := 0
	go func() {
		y = 42
		c <- true
	}()
	x := []int{0, y, 42}
	_ = x
	<-c
}

func TestRaceMapInit(t *testing.T) {
	c := make(chan bool, 1)
	y := 0
	go func() {
		y = 42
		c <- true
	}()
	x := map[int]int{0: 42, y: 42}
	_ = x
	<-c
}

func TestRaceMapInit2(t *testing.T) {
	c := make(chan bool, 1)
	y := 0
	go func() {
		y = 42
		c <- true
	}()
	x := map[int]int{0: 42, 42: y}
	_ = x
	<-c
}

type Inter interface {
	Foo(x int)
}
type InterImpl struct {
	x, y int
}

func (p InterImpl) Foo(x int) {
	// prevent inlining
	z := 42
	x = 85
	y := x / z
	z = y * z
	x = z * y
	_, _, _ = x, y, z
}

type InterImpl2 InterImpl

func (p *InterImpl2) Foo(x int) {
	if p == nil {
		InterImpl{}.Foo(x)
	}
	InterImpl(*p).Foo(x)
}

func TestRaceInterCall(t *testing.T) {
	c := make(chan bool, 1)
	p := InterImpl{}
	var x Inter = p
	go func() {
		p2 := InterImpl{}
		x = p2
		c <- true
	}()
	x.Foo(0)
	<-c
}

func TestRaceInterCall2(t *testing.T) {
	c := make(chan bool, 1)
	p := InterImpl{}
	var x Inter = p
	z := 0
	go func() {
		z = 42
		c <- true
	}()
	x.Foo(z)
	<-c
}

func TestRaceFuncCall(t *testing.T) {
	c := make(chan bool, 1)
	f := func(x, y int) {}
	x, y := 0, 0
	go func() {
		y = 42
		c <- true
	}()
	f(x, y)
	<-c
}

func TestRaceMethodCall(t *testing.T) {
	c := make(chan bool, 1)
	i := InterImpl{}
	x := 0
	go func() {
		x = 42
		c <- true
	}()
	i.Foo(x)
	<-c
}

func TestRaceMethodCall2(t *testing.T) {
	c := make(chan bool, 1)
	i := &InterImpl{}
	go func() {
		i = &InterImpl{}
		c <- true
	}()
	i.Foo(0)
	<-c
}

// Method value with concrete value receiver.
func TestRaceMethodValue(t *testing.T) {
	c := make(chan bool, 1)
	i := InterImpl{}
	go func() {
		i = InterImpl{}
		c <- true
	}()
	_ = i.Foo
	<-c
}

// Method value with interface receiver.
func TestRaceMethodValue2(t *testing.T) {
	c := make(chan bool, 1)
	var i Inter = InterImpl{}
	go func() {
		i = InterImpl{}
		c <- true
	}()
	_ = i.Foo
	<-c
}

// Method value with implicit dereference.
func TestRaceMethodValue3(t *testing.T) {
	c := make(chan bool, 1)
	i := &InterImpl{}
	go func() {
		*i = InterImpl{}
		c <- true
	}()
	_ = i.Foo // dereferences i.
	<-c
}

// Method value implicitly taking receiver address.
func TestNoRaceMethodValue(t *testing.T) {
	c := make(chan bool, 1)
	i := InterImpl2{}
	go func() {
		i = InterImpl2{}
		c <- true
	}()
	_ = i.Foo // takes the address of i only.
	<-c
}

func TestRacePanicArg(t *testing.T) {
	c := make(chan bool, 1)
	err := errors.New("err")
	go func() {
		err = errors.New("err2")
		c <- true
	}()
	defer func() {
		recover()
		<-c
	}()
	panic(err)
}

func TestRaceDeferArg(t *testing.T) {
	c := make(chan bool, 1)
	x := 0
	go func() {
		x = 42
		c <- true
	}()
	func() {
		defer func(x int) {
		}(x)
	}()
	<-c
}

type DeferT int

func (d DeferT) Foo() {
}

func TestRaceDeferArg2(t *testing.T) {
	c := make(chan bool, 1)
	var x DeferT
	go func() {
		var y DeferT
		x = y
		c <- true
	}()
	func() {
		defer x.Foo()
	}()
	<-c
}

func TestNoRaceAddrExpr(t *testing.T) {
	c := make(chan bool, 1)
	x := 0
	go func() {
		x = 42
		c <- true
	}()
	_ = &x
	<-c
}

type AddrT struct {
	_ [256]byte
	x int
}

type AddrT2 struct {
	_ [512]byte
	p *AddrT
}

func TestRaceAddrExpr(t *testing.T) {
	c := make(chan bool, 1)
	a := AddrT2{p: &AddrT{x: 42}}
	go func() {
		a.p = &AddrT{x: 43}
		c <- true
	}()
	_ = &a.p.x
	<-c
}

func TestRaceTypeAssert(t *testing.T) {
	c := make(chan bool, 1)
	x := 0
	var i interface{} = x
	go func() {
		y := 0
		i = y
		c <- true
	}()
	_ = i.(int)
	<-c
}

func TestRaceBlockAs(t *testing.T) {
	c := make(chan bool, 1)
	var x, y int
	go func() {
		x = 42
		c <- true
	}()
	x, y = y, x
	<-c
}

func TestRaceSliceSlice(t *testing.T) {
	c := make(chan bool, 1)
	x := make([]int, 10)
	go func() {
		x = make([]int, 20)
		c <- true
	}()
	_ = x[2:3]
	<-c
}

func TestRaceSliceSlice2(t *testing.T) {
	c := make(chan bool, 1)
	x := make([]int, 10)
	i := 2
	go func() {
		i = 3
		c <- true
	}()
	_ = x[i:4]
	<-c
}

func TestRaceSliceString(t *testing.T) {
	c := make(chan bool, 1)
	x := "hello"
	go func() {
		x = "world"
		c <- true
	}()
	_ = x[2:3]
	<-c
}

// http://golang.org/issue/4453
func TestRaceFailingSliceStruct(t *testing.T) {
	type X struct {
		x, y int
	}
	c := make(chan bool, 1)
	x := make([]X, 10)
	go func() {
		y := make([]X, 10)
		copy(y, x)
		c <- true
	}()
	x[1].y = 42
	<-c
}

func TestRaceFailingAppendSliceStruct(t *testing.T) {
	type X struct {
		x, y int
	}
	c := make(chan bool, 1)
	x := make([]X, 10)
	go func() {
		y := make([]X, 0, 10)
		y = append(y, x...)
		c <- true
	}()
	x[1].y = 42
	<-c
}

func TestRaceStructInd(t *testing.T) {
	c := make(chan bool, 1)
	type Item struct {
		x, y int
	}
	i := Item{}
	go func(p *Item) {
		*p = Item{}
		c <- true
	}(&i)
	i.y = 42
	<-c
}

func TestRaceAsFunc1(t *testing.T) {
	var s []byte
	c := make(chan bool, 1)
	go func() {
		var err error
		s, err = func() ([]byte, error) {
			t := []byte("hello world")
			return t, nil
		}()
		c <- true
		_ = err
	}()
	_ = string(s)
	<-c
}

func TestRaceAsFunc2(t *testing.T) {
	c := make(chan bool, 1)
	x := 0
	go func() {
		func(x int) {
		}(x)
		c <- true
	}()
	x = 42
	<-c
}

func TestRaceAsFunc3(t *testing.T) {
	c := make(chan bool, 1)
	var mu sync.Mutex
	x := 0
	go func() {
		func(x int) {
			mu.Lock()
		}(x) // Read of x must be outside of the mutex.
		mu.Unlock()
		c <- true
	}()
	mu.Lock()
	x = 42
	mu.Unlock()
	<-c
}

func TestNoRaceAsFunc4(t *testing.T) {
	c := make(chan bool, 1)
	var mu sync.Mutex
	x := 0
	go func() {
		x = func() int { // Write of x must be under the mutex.
			mu.Lock()
			return 42
		}()
		mu.Unlock()
		c <- true
	}()
	mu.Lock()
	x = 42
	mu.Unlock()
	<-c
}

func TestRaceHeapParam(t *testing.T) {
	x := func() (x int) {
		go func() {
			x = 42
		}()
		return
	}()
	_ = x
}

func TestNoRaceEmptyStruct(t *testing.T) {
	type Empty struct{}
	type X struct {
		y int64
		Empty
	}
	type Y struct {
		x X
		y int64
	}
	c := make(chan X)
	var y Y
	go func() {
		x := y.x
		c <- x
	}()
	y.y = 42
	<-c
}

func TestRaceNestedStruct(t *testing.T) {
	type X struct {
		x, y int
	}
	type Y struct {
		x X
	}
	c := make(chan Y)
	var y Y
	go func() {
		c <- y
	}()
	y.x.y = 42
	<-c
}
