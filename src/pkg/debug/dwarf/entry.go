// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DWARF debug information entry parser.
// An entry is a sequence of data items of a given format.
// The first word in the entry is an index into what DWARF
// calls the ``abbreviation table.''  An abbreviation is really
// just a type descriptor: it's an array of attribute tag/value format pairs.

package dwarf

import (
	"os";
	"strconv";
)

// a single entry's description: a sequence of attributes
type abbrev struct {
	tag Tag;
	children bool;
	field []afield;
}

type afield struct {
	attr Attr;
	fmt format;
}

// a map from entry format ids to their descriptions
type abbrevTable map[uint32]abbrev

// ParseAbbrev returns the abbreviation table that starts at byte off
// in the .debug_abbrev section.
func (d *Data) parseAbbrev(off uint32) (abbrevTable, os.Error) {
	if m, ok := d.abbrevCache[off]; ok {
		return m, nil;
	}

	data := d.abbrev;
	if off > uint32(len(data)) {
		data = nil;
	} else {
		data = data[off:len(data)];
	}
	b := makeBuf(d, "abbrev", 0, data, 0);

	// Error handling is simplified by the buf getters
	// returning an endless stream of 0s after an error.
	m := make(abbrevTable);
	for {
		// Table ends with id == 0.
		id := uint32(b.uint());
		if id == 0 {
			break;
		}

		// Walk over attributes, counting.
		n := 0;
		b1 := b;	// Read from copy of b.
		b1.uint();
		b1.uint8();
		for {
			tag := b1.uint();
			fmt := b1.uint();
			if tag == 0 && fmt == 0 {
				break;
			}
			n++;
		}
		if b1.err != nil {
			return nil, b1.err;
		}

		// Walk over attributes again, this time writing them down.
		var a abbrev;
		a.tag = Tag(b.uint());
		a.children = b.uint8() != 0;
		a.field = make([]afield, n);
		for i := range a.field {
			a.field[i].attr = Attr(b.uint());
			a.field[i].fmt = format(b.uint());
		}
		b.uint();
		b.uint();

		m[id] = a;
	}
	if b.err != nil {
		return nil, b.err;
	}
	d.abbrevCache[off] = m;
	return m, nil;
}

// An entry is a sequence of attribute/value pairs.
type Entry struct {
	Offset Offset;	// offset of Entry in DWARF info
	Tag Tag;	// tag (kind of Entry)
	Children bool;	// whether Entry is followed by children
	Field []Field;
}

// A Field is a single attribute/value pair in an Entry.
type Field struct {
	Attr Attr;
	Val interface{};
}

// An Offset represents the location of an Entry within the DWARF info.
// (See Reader.Seek.)
type Offset uint32

// Entry reads a single entry from buf, decoding
// according to the given abbreviation table.
func (b *buf) entry(atab abbrevTable, ubase Offset) *Entry {
	off := b.off;
	id := uint32(b.uint());
	if id == 0 {
		return &Entry{};
	}
	a, ok := atab[id];
	if !ok {
		b.error("unknown abbreviation table index");
		return nil;
	}
	e := &Entry{
		Offset: off,
		Tag: a.tag,
		Children: a.children,
		Field: make([]Field, len(a.field))
	};
	for i := range e.Field {
		e.Field[i].Attr = a.field[i].attr;
		fmt := a.field[i].fmt;
		if fmt == formIndirect {
			fmt = format(b.uint());
		}
		var val interface{};
		switch fmt {
		default:
			b.error("unknown entry attr format");

		// address
		case formAddr:
			val = b.addr();

		// block
		case formDwarfBlock1:
			val = b.bytes(int(b.uint8()));
		case formDwarfBlock2:
			val = b.bytes(int(b.uint16()));
		case formDwarfBlock4:
			val = b.bytes(int(b.uint32()));
		case formDwarfBlock:
			val = b.bytes(int(b.uint()));

		// constant
		case formData1:
			val = uint64(b.uint8());
		case formData2:
			val = uint64(b.uint16());
		case formData4:
			val = uint64(b.uint32());
		case formData8:
			val = uint64(b.uint64());
		case formSdata:
			val = int64(b.int());
		case formUdata:
			val = uint64(b.uint());

		// flag
		case formFlag:
			val = b.uint8() == 1;

		// reference to other entry
		case formRefAddr:
			val = Offset(b.addr());
		case formRef1:
			val = Offset(b.uint8()) + ubase;
		case formRef2:
			val = Offset(b.uint16()) + ubase;
		case formRef4:
			val = Offset(b.uint32()) + ubase;
		case formRef8:
			val = Offset(b.uint64()) + ubase;
		case formRefUdata:
			val = Offset(b.uint()) + ubase;

		// string
		case formString:
			val = b.string();
		case formStrp:
			off := b.uint32();	// offset into .debug_str
			if b.err != nil {
				return nil;
			}
			b1 := makeBuf(b.dwarf, "str", 0, b.dwarf.str, 0);
			b1.skip(int(off));
			val = b1.string();
			if b1.err != nil {
				b.err = b1.err;
				return nil;
			}
		}
		e.Field[i].Val = val;
	}
	if b.err != nil {
		return nil;
	}
	return e;
}

// A Reader allows reading Entry structures from a DWARF ``info'' section.
type Reader struct {
	b buf;
	d *Data;
	err os.Error;
	unit int;
}

// Reader returns a new Reader for Data.
// The reader is positioned at byte offset 0 in the DWARF ``info'' section.
func (d *Data) Reader() *Reader {
	r := &Reader{d: d};
	r.Seek(0);
	return r;
}

// Seek positions the Reader at offset off in the encoded entry stream.
// Offset 0 can be used to denote the first entry.
func (r *Reader) Seek(off Offset) {
	d := r.d;
	r.err = nil;
	if off == 0 {
		if len(d.unit) == 0 {
			return;
		}
		u := &d.unit[0];
		r.unit = 0;
		r.b = makeBuf(r.d, "info", u.off, u.data, u.addrsize);
		return;
	}

	// TODO(rsc): binary search (maybe a new package)
	var i int;
	var u *unit;
	for i = range d.unit {
		u = &d.unit[i];
		if u.off <= off && off < u.off+Offset(len(u.data)) {
			r.unit = i;
			r.b = makeBuf(r.d, "info", off, u.data[off-u.off:len(u.data)], u.addrsize);
			return;
		}
	}
	r.err = os.NewError("offset out of range");
}

// maybeNextUnit advances to the next unit if this one is finished.
func (r *Reader) maybeNextUnit() {
	for len(r.b.data) == 0 && r.unit < len(r.d.unit) {
		r.unit++;
		u := &r.d.unit[r.unit];
		r.b = makeBuf(r.d, "info", u.off, u.data, u.addrsize);
	}
}

// Next reads the next entry from the encoded entry stream.
// It returns nil, nil when it reaches the end of the section.
// It returns an error if the current offset is invalid or the data at the
// offset cannot be decoded as a valid Entry.
func (r *Reader) Next() (*Entry, os.Error) {
	if r.err != nil {
		return nil, r.err;
	}
	r.maybeNextUnit();
	if len(r.b.data) == 0 {
		return nil, nil;
	}
	u := &r.d.unit[r.unit];
	e := r.b.entry(u.atable, u.base);
	r.err = r.b.err;
	return e, r.err;
}
