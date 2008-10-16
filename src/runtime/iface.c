// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "runtime.h"

static	int32	debug	= 0;

enum
{
	ASIMP		= 0,
	ASTRING,
	APTR,
	AINTER,
};

typedef	struct	Sigt	Sigt;
typedef	struct	Sigi	Sigi;
typedef	struct	Map	Map;

struct	Sigt
{
	byte*	name;
	uint32	hash;		// hash of type		// first is alg
	uint32	offset;		// offset of substruct	// first is width
	void	(*fun)(void);
};

struct	Sigi
{
	byte*	name;
	uint32	hash;
	uint32	perm;		// location of fun in Sigt // first is size
};

struct	Map
{
	Sigi*	sigi;
	Sigt*	sigt;
	Map*	link;
	int32	bad;
	int32	unused;
	void	(*fun[])(void);
};

static	Map*	hash[1009];

#define	END	nil,0,0,nil

Sigi	sys·sigi_inter[2] =	{ (byte*)"sys·nilinter", 0, 0, nil, 0, 0 };

Sigt	sys·sigt_int8[2] =	{ (byte*)"sys·int8", ASIMP, 1, nil, END };
Sigt	sys·sigt_int16[2] =	{ (byte*)"sys·int16", ASIMP, 2, nil, END };
Sigt	sys·sigt_int32[2] =	{ (byte*)"sys·int32", ASIMP, 4, nil, END };
Sigt	sys·sigt_int64[2] =	{ (byte*)"sys·int64", ASIMP, 8, nil, END };

Sigt	sys·sigt_uint8[2] =	{ (byte*)"sys·uint8", ASIMP, 1, nil, END };
Sigt	sys·sigt_uint16[2] =	{ (byte*)"sys·uint16", ASIMP, 2, nil, END };
Sigt	sys·sigt_uint32[2] =	{ (byte*)"sys·uint32", ASIMP, 4, nil, END };
Sigt	sys·sigt_uint64[2] =	{ (byte*)"sys·uint64", ASIMP, 8, nil, END };

Sigt	sys·sigt_float32[2] =	{ (byte*)"sys·float32", ASIMP, 4, nil, END };
Sigt	sys·sigt_float64[2] =	{ (byte*)"sys·float64", ASIMP, 8, nil, END };
//Sigt	sys·sigt_float80[2] =	{ (byte*)"sys·float80", ASIMP, 0, nil, END };

Sigt	sys·sigt_bool[2] =	{ (byte*)"sys·bool", ASIMP, 1, nil, END };
Sigt	sys·sigt_string[2] =	{ (byte*)"sys·string", ASTRING, 8, nil, END };

static void
printsigi(Sigi *si)
{
	int32 i;
	byte *name;

	sys·printpointer(si);
	prints("{");
	for(i=1;; i++) {
		name = si[i].name;
		if(name == nil)
			break;
		prints("[");
		sys·printint(i);
		prints("]\"");
		prints((int8*)name);
		prints("\"");
		sys·printint(si[i].hash%999);
		prints("/");
		sys·printint(si[i].perm);
	}
	prints("}");
}

static void
printsigt(Sigt *st)
{
	int32 i;
	byte *name;

	sys·printpointer(st);
	prints("{");
	sys·printint(st[0].hash);	// first element has alg
	prints(",");
	sys·printint(st[0].offset);	// first element has width
	for(i=1;; i++) {
		name = st[i].name;
		if(name == nil)
			break;
		prints("[");
		sys·printint(i);
		prints("]\"");
		prints((int8*)name);
		prints("\"");
		sys·printint(st[i].hash%999);
		prints("/");
		sys·printint(st[i].offset);
		prints("/");
		sys·printpointer(st[i].fun);
	}
	prints("}");
}

static void
printiface(Map *im, void *it)
{
	prints("(");
	sys·printpointer(im);
	prints(",");
	sys·printpointer(it);
	prints(")");
}

static Map*
hashmap(Sigi *si, Sigt *st)
{
	int32 nt, ni;
	uint32 ihash, h;
	byte *sname, *iname;
	Map *m;

	h = ((uint32)(uint64)si + (uint32)(uint64)st) % nelem(hash);
	for(m=hash[h]; m!=nil; m=m->link) {
		if(m->sigi == si && m->sigt == st) {
			if(m->bad) {
				throw("bad hashmap");
				m = nil;
			}
			// prints("old hashmap\n");
			return m;
		}
	}

	ni = si[0].perm;	// first entry has size
	m = mal(sizeof(*m) + ni*sizeof(m->fun[0]));
	m->sigi = si;
	m->sigt = st;

	nt = 1;
	for(ni=1;; ni++) {	// ni=1: skip first word
		iname = si[ni].name;
		if(iname == nil)
			break;

		// pick up next name from
		// interface signature
		ihash = si[ni].hash;

		for(;; nt++) {
			// pick up and compare next name
			// from structure signature
			sname = st[nt].name;
			if(sname == nil) {
				prints("cannot convert type ");
				prints((int8*)st[0].name);
				prints(" to interface ");
				prints((int8*)si[0].name);
				prints(": missing method ");
				prints((int8*)iname);
				prints("\n");
				throw("interface conversion");
				m->bad = 1;
				m->link = hash[h];
				hash[h] = m;
				return nil;
			}
			if(ihash == st[nt].hash && strcmp(sname, iname) == 0)
				break;
		}
		m->fun[si[ni].perm] = st[nt].fun;
	}
	m->link = hash[h];
	hash[h] = m;
	// prints("new hashmap\n");
	return m;
}

// ifaceT2I(sigi *byte, sigt *byte, elem any) (ret any);
void
sys·ifaceT2I(Sigi *si, Sigt *st, void *elem, Map *retim, void *retit)
{
//	int32 alg, wid;

	if(debug) {
		prints("T2I sigi=");
		printsigi(si);
		prints(" sigt=");
		printsigt(st);
		prints(" elem=");
		sys·printpointer(elem);
		prints("\n");
	}

	retim = hashmap(si, st);

//	alg = st->hash;
//	wid = st->offset;
//	algarray[alg].copy(wid, &retit, &elem);
	retit = elem;		// for speed could do this

	if(debug) {
		prints("T2I ret=");
		printiface(retim, retit);
		prints("\n");
	}

	FLUSH(&retim);
	FLUSH(&retit);
}

// ifaceI2T(sigt *byte, iface any) (ret any);
void
sys·ifaceI2T(Sigt *st, Map *im, void *it, void *ret)
{
//	int32 alg, wid;

	if(debug) {
		prints("I2T sigt=");
		printsigt(st);
		prints(" iface=");
		printiface(im, it);
		prints("\n");
	}

	if(im == nil)
		throw("ifaceI2T: nil map");

	if(im->sigt != st)
		throw("ifaceI2T: wrong type");

//	alg = st->hash;
//	wid = st->offset;
//	algarray[alg].copy(wid, &ret, &it);
	ret = it;

	if(debug) {
		prints("I2T ret=");
		sys·printpointer(ret);
		prints("\n");
	}

	FLUSH(&ret);
}

// ifaceI2I(sigi *byte, iface any) (ret any);
void
sys·ifaceI2I(Sigi *si, Map *im, void *it, Map *retim, void *retit)
{
	if(debug) {
		prints("I2I sigi=");
		printsigi(si);
		prints(" iface=");
		printiface(im, it);
		prints("\n");
	}

	if(im == nil) {
		// If incoming interface is uninitialized (zeroed)
		// make the outgoing interface zeroed as well.
		retim = nil;
		retit = nil;
	} else {
		retit = it;
		retim = im;
		if(im->sigi != si)
			retim = hashmap(si, im->sigt);
	}

	if(debug) {
		prints("I2I ret=");
		printiface(retim, retit);
		prints("\n");
	}

	FLUSH(&retim);
	FLUSH(&retit);
}

// ifaceeq(i1 any, i2 any) (ret bool);
void
sys·ifaceeq(Map *im1, void *it1, Map *im2, void *it2, byte ret)
{
	int32 alg, wid;

	if(debug) {
		prints("Ieq i1=");
		printiface(im1, it1);
		prints(" i2=");
		printiface(im2, it2);
		prints("\n");
	}

	ret = false;

	// are they both nil
	if(im1 == nil) {
		if(im2 == nil)
			goto yes;
		goto no;
	}
	if(im2 == nil)
		goto no;

	// value
	alg = im1->sigt->hash;
	if(alg != im2->sigt->hash)
		goto no;

	wid = im1->sigt->offset;
	if(wid != im2->sigt->offset)
		goto no;

	if(!algarray[alg].equal(wid, &it1, &it2))
		goto no;

yes:
	ret = true;
no:
	if(debug) {
		prints("Ieq ret=");
		sys·printbool(ret);
		prints("\n");
	}
	FLUSH(&ret);
}

void
sys·printinter(Map *im, void *it)
{
	printiface(im, it);
}
