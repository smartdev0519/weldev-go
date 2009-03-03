// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Unicode table access.  A quick start for now.

// TODO: Generated by hand.
// Should generate automatically from Unicode
// tables, expand to other properties, split into files,
// link in only the tables that are used by the program,
// etc.

package unicode

type Range struct {
	lo int;
	hi int;
	stride int;
}

var Upper = []Range{
	Range{0x0041, 0x005a, 1},
	Range{0x00c0, 0x00d6, 1},
	Range{0x00d8, 0x00de, 1},
	Range{0x0100, 0x0136, 2},
	Range{0x0139, 0x0147, 2},
	Range{0x014a, 0x0176, 2},
	Range{0x0178, 0x0179, 1},
	Range{0x017b, 0x017d, 2},
	Range{0x0181, 0x0182, 1},
	Range{0x0184, 0x0184, 1},
	Range{0x0186, 0x0187, 1},
	Range{0x0189, 0x018b, 1},
	Range{0x018e, 0x0191, 1},
	Range{0x0193, 0x0194, 1},
	Range{0x0196, 0x0198, 1},
	Range{0x019c, 0x019d, 1},
	Range{0x019f, 0x01a0, 1},
	Range{0x01a2, 0x01a4, 2},
	Range{0x01a6, 0x01a7, 1},
	Range{0x01a9, 0x01ac, 3},
	Range{0x01ae, 0x01af, 1},
	Range{0x01b1, 0x01b3, 1},
	Range{0x01b5, 0x01b5, 1},
	Range{0x01b7, 0x01b8, 1},
	Range{0x01bc, 0x01c4, 8},
	Range{0x01c7, 0x01cd, 3},
	Range{0x01cf, 0x01db, 2},
	Range{0x01de, 0x01ee, 2},
	Range{0x01f1, 0x01f4, 3},
	Range{0x01f6, 0x01f8, 1},
	Range{0x01fa, 0x0232, 2},
	Range{0x023a, 0x023b, 1},
	Range{0x023d, 0x023e, 1},
	Range{0x0241, 0x0241, 1},
	Range{0x0243, 0x0246, 1},
	Range{0x0248, 0x024e, 2},
	Range{0x0370, 0x0372, 2},
	Range{0x0376, 0x0386, 16},
	Range{0x0388, 0x038a, 1},
	Range{0x038c, 0x038c, 1},
	Range{0x038e, 0x038f, 1},
	Range{0x0391, 0x03a1, 1},
	Range{0x03a3, 0x03ab, 1},
	Range{0x03cf, 0x03cf, 1},
	Range{0x03d2, 0x03d4, 1},
	Range{0x03d8, 0x03ee, 2},
	Range{0x03f4, 0x03f7, 3},
	Range{0x03f9, 0x03fa, 1},
	Range{0x03fd, 0x042f, 1},
	Range{0x0460, 0x0480, 2},
	Range{0x048a, 0x04be, 2},
	Range{0x04c0, 0x04c1, 1},
	Range{0x04c3, 0x04cd, 2},
	Range{0x04d0, 0x0522, 2},
	Range{0x0531, 0x0556, 1},
	Range{0x10a0, 0x10c5, 1},
	Range{0x1e00, 0x1e94, 2},
	Range{0x1e9e, 0x1efe, 2},
	Range{0x1f08, 0x1f0f, 1},
	Range{0x1f18, 0x1f1d, 1},
	Range{0x1f28, 0x1f2f, 1},
	Range{0x1f38, 0x1f3f, 1},
	Range{0x1f48, 0x1f4d, 1},
	Range{0x1f59, 0x1f5f, 2},
	Range{0x1f68, 0x1f6f, 1},
	Range{0x1fb8, 0x1fbb, 1},
	Range{0x1fc8, 0x1fcb, 1},
	Range{0x1fd8, 0x1fdb, 1},
	Range{0x1fe8, 0x1fec, 1},
	Range{0x1ff8, 0x1ffb, 1},
	Range{0x2102, 0x2107, 5},
	Range{0x210b, 0x210d, 1},
	Range{0x2110, 0x2112, 1},
	Range{0x2115, 0x2115, 1},
	Range{0x2119, 0x211d, 1},
	Range{0x2124, 0x2128, 2},
	Range{0x212a, 0x212d, 1},
	Range{0x2130, 0x2133, 1},
	Range{0x213e, 0x213f, 1},
	Range{0x2145, 0x2183, 62},
	Range{0x2c00, 0x2c2e, 1},
	Range{0x2c60, 0x2c60, 1},
	Range{0x2c62, 0x2c64, 1},
	Range{0x2c67, 0x2c6b, 2},
	Range{0x2c6d, 0x2c6f, 1},
	Range{0x2c72, 0x2c75, 3},
	Range{0x2c80, 0x2ce2, 2},
	Range{0xa640, 0xa65e, 2},
	Range{0xa662, 0xa66c, 2},
	Range{0xa680, 0xa696, 2},
	Range{0xa722, 0xa72e, 2},
	Range{0xa732, 0xa76e, 2},
	Range{0xa779, 0xa77b, 2},
	Range{0xa77d, 0xa77e, 1},
	Range{0xa780, 0xa786, 2},
	Range{0xa78b, 0xa78b, 1},
	Range{0xff21, 0xff3a, 1},
	Range{0x10400, 0x10427, 1},
	Range{0x1d400, 0x1d419, 1},
	Range{0x1d434, 0x1d44d, 1},
	Range{0x1d468, 0x1d481, 1},
	Range{0x1d49c, 0x1d49c, 1},
	Range{0x1d49e, 0x1d49f, 1},
	Range{0x1d4a2, 0x1d4a2, 1},
	Range{0x1d4a5, 0x1d4a6, 1},
	Range{0x1d4a9, 0x1d4ac, 1},
	Range{0x1d4ae, 0x1d4b5, 1},
	Range{0x1d4d0, 0x1d4e9, 1},
	Range{0x1d504, 0x1d505, 1},
	Range{0x1d507, 0x1d50a, 1},
	Range{0x1d50d, 0x1d514, 1},
	Range{0x1d516, 0x1d51c, 1},
	Range{0x1d538, 0x1d539, 1},
	Range{0x1d53b, 0x1d53e, 1},
	Range{0x1d540, 0x1d544, 1},
	Range{0x1d546, 0x1d546, 1},
	Range{0x1d54a, 0x1d550, 1},
	Range{0x1d56c, 0x1d585, 1},
	Range{0x1d5a0, 0x1d5b9, 1},
	Range{0x1d5d4, 0x1d5ed, 1},
	Range{0x1d608, 0x1d621, 1},
	Range{0x1d63c, 0x1d655, 1},
	Range{0x1d670, 0x1d689, 1},
	Range{0x1d6a8, 0x1d6c0, 1},
	Range{0x1d6e2, 0x1d6fa, 1},
	Range{0x1d71c, 0x1d734, 1},
	Range{0x1d756, 0x1d76e, 1},
	Range{0x1d790, 0x1d7a8, 1},
	Range{0x1d7ca, 0x1d7ca, 1},
}

var Letter = []Range {
	Range{0x0041, 0x005a, 1},
	Range{0x0061, 0x007a, 1},
	Range{0x00aa, 0x00b5, 11},
	Range{0x00ba, 0x00ba, 1},
	Range{0x00c0, 0x00d6, 1},
	Range{0x00d8, 0x00f6, 1},
	Range{0x00f8, 0x02c1, 1},
	Range{0x02c6, 0x02d1, 1},
	Range{0x02e0, 0x02e4, 1},
	Range{0x02ec, 0x02ee, 2},
	Range{0x0370, 0x0374, 1},
	Range{0x0376, 0x0377, 1},
	Range{0x037a, 0x037d, 1},
	Range{0x0386, 0x0386, 1},
	Range{0x0388, 0x038a, 1},
	Range{0x038c, 0x038c, 1},
	Range{0x038e, 0x03a1, 1},
	Range{0x03a3, 0x03f5, 1},
	Range{0x03f7, 0x0481, 1},
	Range{0x048a, 0x0523, 1},
	Range{0x0531, 0x0556, 1},
	Range{0x0559, 0x0559, 1},
	Range{0x0561, 0x0587, 1},
	Range{0x05d0, 0x05ea, 1},
	Range{0x05f0, 0x05f2, 1},
	Range{0x0621, 0x064a, 1},
	Range{0x066e, 0x066f, 1},
	Range{0x0671, 0x06d3, 1},
	Range{0x06d5, 0x06d5, 1},
	Range{0x06e5, 0x06e6, 1},
	Range{0x06ee, 0x06ef, 1},
	Range{0x06fa, 0x06fc, 1},
	Range{0x06ff, 0x0710, 17},
	Range{0x0712, 0x072f, 1},
	Range{0x074d, 0x07a5, 1},
	Range{0x07b1, 0x07b1, 1},
	Range{0x07ca, 0x07ea, 1},
	Range{0x07f4, 0x07f5, 1},
	Range{0x07fa, 0x07fa, 1},
	Range{0x0904, 0x0939, 1},
	Range{0x093d, 0x0950, 19},
	Range{0x0958, 0x0961, 1},
	Range{0x0971, 0x0972, 1},
	Range{0x097b, 0x097f, 1},
	Range{0x0985, 0x098c, 1},
	Range{0x098f, 0x0990, 1},
	Range{0x0993, 0x09a8, 1},
	Range{0x09aa, 0x09b0, 1},
	Range{0x09b2, 0x09b2, 1},
	Range{0x09b6, 0x09b9, 1},
	Range{0x09bd, 0x09ce, 17},
	Range{0x09dc, 0x09dd, 1},
	Range{0x09df, 0x09e1, 1},
	Range{0x09f0, 0x09f1, 1},
	Range{0x0a05, 0x0a0a, 1},
	Range{0x0a0f, 0x0a10, 1},
	Range{0x0a13, 0x0a28, 1},
	Range{0x0a2a, 0x0a30, 1},
	Range{0x0a32, 0x0a33, 1},
	Range{0x0a35, 0x0a36, 1},
	Range{0x0a38, 0x0a39, 1},
	Range{0x0a59, 0x0a5c, 1},
	Range{0x0a5e, 0x0a5e, 1},
	Range{0x0a72, 0x0a74, 1},
	Range{0x0a85, 0x0a8d, 1},
	Range{0x0a8f, 0x0a91, 1},
	Range{0x0a93, 0x0aa8, 1},
	Range{0x0aaa, 0x0ab0, 1},
	Range{0x0ab2, 0x0ab3, 1},
	Range{0x0ab5, 0x0ab9, 1},
	Range{0x0abd, 0x0ad0, 19},
	Range{0x0ae0, 0x0ae1, 1},
	Range{0x0b05, 0x0b0c, 1},
	Range{0x0b0f, 0x0b10, 1},
	Range{0x0b13, 0x0b28, 1},
	Range{0x0b2a, 0x0b30, 1},
	Range{0x0b32, 0x0b33, 1},
	Range{0x0b35, 0x0b39, 1},
	Range{0x0b3d, 0x0b3d, 1},
	Range{0x0b5c, 0x0b5d, 1},
	Range{0x0b5f, 0x0b61, 1},
	Range{0x0b71, 0x0b83, 18},
	Range{0x0b85, 0x0b8a, 1},
	Range{0x0b8e, 0x0b90, 1},
	Range{0x0b92, 0x0b95, 1},
	Range{0x0b99, 0x0b9a, 1},
	Range{0x0b9c, 0x0b9c, 1},
	Range{0x0b9e, 0x0b9f, 1},
	Range{0x0ba3, 0x0ba4, 1},
	Range{0x0ba8, 0x0baa, 1},
	Range{0x0bae, 0x0bb9, 1},
	Range{0x0bd0, 0x0bd0, 1},
	Range{0x0c05, 0x0c0c, 1},
	Range{0x0c0e, 0x0c10, 1},
	Range{0x0c12, 0x0c28, 1},
	Range{0x0c2a, 0x0c33, 1},
	Range{0x0c35, 0x0c39, 1},
	Range{0x0c3d, 0x0c3d, 1},
	Range{0x0c58, 0x0c59, 1},
	Range{0x0c60, 0x0c61, 1},
	Range{0x0c85, 0x0c8c, 1},
	Range{0x0c8e, 0x0c90, 1},
	Range{0x0c92, 0x0ca8, 1},
	Range{0x0caa, 0x0cb3, 1},
	Range{0x0cb5, 0x0cb9, 1},
	Range{0x0cbd, 0x0cde, 33},
	Range{0x0ce0, 0x0ce1, 1},
	Range{0x0d05, 0x0d0c, 1},
	Range{0x0d0e, 0x0d10, 1},
	Range{0x0d12, 0x0d28, 1},
	Range{0x0d2a, 0x0d39, 1},
	Range{0x0d3d, 0x0d3d, 1},
	Range{0x0d60, 0x0d61, 1},
	Range{0x0d7a, 0x0d7f, 1},
	Range{0x0d85, 0x0d96, 1},
	Range{0x0d9a, 0x0db1, 1},
	Range{0x0db3, 0x0dbb, 1},
	Range{0x0dbd, 0x0dbd, 1},
	Range{0x0dc0, 0x0dc6, 1},
	Range{0x0e01, 0x0e30, 1},
	Range{0x0e32, 0x0e33, 1},
	Range{0x0e40, 0x0e46, 1},
	Range{0x0e81, 0x0e82, 1},
	Range{0x0e84, 0x0e84, 1},
	Range{0x0e87, 0x0e88, 1},
	Range{0x0e8a, 0x0e8d, 3},
	Range{0x0e94, 0x0e97, 1},
	Range{0x0e99, 0x0e9f, 1},
	Range{0x0ea1, 0x0ea3, 1},
	Range{0x0ea5, 0x0ea7, 2},
	Range{0x0eaa, 0x0eab, 1},
	Range{0x0ead, 0x0eb0, 1},
	Range{0x0eb2, 0x0eb3, 1},
	Range{0x0ebd, 0x0ebd, 1},
	Range{0x0ec0, 0x0ec4, 1},
	Range{0x0ec6, 0x0ec6, 1},
	Range{0x0edc, 0x0edd, 1},
	Range{0x0f00, 0x0f00, 1},
	Range{0x0f40, 0x0f47, 1},
	Range{0x0f49, 0x0f6c, 1},
	Range{0x0f88, 0x0f8b, 1},
	Range{0x1000, 0x102a, 1},
	Range{0x103f, 0x103f, 1},
	Range{0x1050, 0x1055, 1},
	Range{0x105a, 0x105d, 1},
	Range{0x1061, 0x1061, 1},
	Range{0x1065, 0x1066, 1},
	Range{0x106e, 0x1070, 1},
	Range{0x1075, 0x1081, 1},
	Range{0x108e, 0x108e, 1},
	Range{0x10a0, 0x10c5, 1},
	Range{0x10d0, 0x10fa, 1},
	Range{0x10fc, 0x10fc, 1},
	Range{0x1100, 0x1159, 1},
	Range{0x115f, 0x11a2, 1},
	Range{0x11a8, 0x11f9, 1},
	Range{0x1200, 0x1248, 1},
	Range{0x124a, 0x124d, 1},
	Range{0x1250, 0x1256, 1},
	Range{0x1258, 0x1258, 1},
	Range{0x125a, 0x125d, 1},
	Range{0x1260, 0x1288, 1},
	Range{0x128a, 0x128d, 1},
	Range{0x1290, 0x12b0, 1},
	Range{0x12b2, 0x12b5, 1},
	Range{0x12b8, 0x12be, 1},
	Range{0x12c0, 0x12c0, 1},
	Range{0x12c2, 0x12c5, 1},
	Range{0x12c8, 0x12d6, 1},
	Range{0x12d8, 0x1310, 1},
	Range{0x1312, 0x1315, 1},
	Range{0x1318, 0x135a, 1},
	Range{0x1380, 0x138f, 1},
	Range{0x13a0, 0x13f4, 1},
	Range{0x1401, 0x166c, 1},
	Range{0x166f, 0x1676, 1},
	Range{0x1681, 0x169a, 1},
	Range{0x16a0, 0x16ea, 1},
	Range{0x1700, 0x170c, 1},
	Range{0x170e, 0x1711, 1},
	Range{0x1720, 0x1731, 1},
	Range{0x1740, 0x1751, 1},
	Range{0x1760, 0x176c, 1},
	Range{0x176e, 0x1770, 1},
	Range{0x1780, 0x17b3, 1},
	Range{0x17d7, 0x17dc, 5},
	Range{0x1820, 0x1877, 1},
	Range{0x1880, 0x18a8, 1},
	Range{0x18aa, 0x18aa, 1},
	Range{0x1900, 0x191c, 1},
	Range{0x1950, 0x196d, 1},
	Range{0x1970, 0x1974, 1},
	Range{0x1980, 0x19a9, 1},
	Range{0x19c1, 0x19c7, 1},
	Range{0x1a00, 0x1a16, 1},
	Range{0x1b05, 0x1b33, 1},
	Range{0x1b45, 0x1b4b, 1},
	Range{0x1b83, 0x1ba0, 1},
	Range{0x1bae, 0x1baf, 1},
	Range{0x1c00, 0x1c23, 1},
	Range{0x1c4d, 0x1c4f, 1},
	Range{0x1c5a, 0x1c7d, 1},
	Range{0x1d00, 0x1dbf, 1},
	Range{0x1e00, 0x1f15, 1},
	Range{0x1f18, 0x1f1d, 1},
	Range{0x1f20, 0x1f45, 1},
	Range{0x1f48, 0x1f4d, 1},
	Range{0x1f50, 0x1f57, 1},
	Range{0x1f59, 0x1f5d, 2},
	Range{0x1f5f, 0x1f7d, 1},
	Range{0x1f80, 0x1fb4, 1},
	Range{0x1fb6, 0x1fbc, 1},
	Range{0x1fbe, 0x1fbe, 1},
	Range{0x1fc2, 0x1fc4, 1},
	Range{0x1fc6, 0x1fcc, 1},
	Range{0x1fd0, 0x1fd3, 1},
	Range{0x1fd6, 0x1fdb, 1},
	Range{0x1fe0, 0x1fec, 1},
	Range{0x1ff2, 0x1ff4, 1},
	Range{0x1ff6, 0x1ffc, 1},
	Range{0x2071, 0x207f, 14},
	Range{0x2090, 0x2094, 1},
	Range{0x2102, 0x2107, 5},
	Range{0x210a, 0x2113, 1},
	Range{0x2115, 0x2115, 1},
	Range{0x2119, 0x211d, 1},
	Range{0x2124, 0x2128, 2},
	Range{0x212a, 0x212d, 1},
	Range{0x212f, 0x2139, 1},
	Range{0x213c, 0x213f, 1},
	Range{0x2145, 0x2149, 1},
	Range{0x214e, 0x214e, 1},
	Range{0x2183, 0x2184, 1},
	Range{0x2c00, 0x2c2e, 1},
	Range{0x2c30, 0x2c5e, 1},
	Range{0x2c60, 0x2c6f, 1},
	Range{0x2c71, 0x2c7d, 1},
	Range{0x2c80, 0x2ce4, 1},
	Range{0x2d00, 0x2d25, 1},
	Range{0x2d30, 0x2d65, 1},
	Range{0x2d6f, 0x2d6f, 1},
	Range{0x2d80, 0x2d96, 1},
	Range{0x2da0, 0x2da6, 1},
	Range{0x2da8, 0x2dae, 1},
	Range{0x2db0, 0x2db6, 1},
	Range{0x2db8, 0x2dbe, 1},
	Range{0x2dc0, 0x2dc6, 1},
	Range{0x2dc8, 0x2dce, 1},
	Range{0x2dd0, 0x2dd6, 1},
	Range{0x2dd8, 0x2dde, 1},
	Range{0x2e2f, 0x2e2f, 1},
	Range{0x3005, 0x3006, 1},
	Range{0x3031, 0x3035, 1},
	Range{0x303b, 0x303c, 1},
	Range{0x3041, 0x3096, 1},
	Range{0x309d, 0x309f, 1},
	Range{0x30a1, 0x30fa, 1},
	Range{0x30fc, 0x30ff, 1},
	Range{0x3105, 0x312d, 1},
	Range{0x3131, 0x318e, 1},
	Range{0x31a0, 0x31b7, 1},
	Range{0x31f0, 0x31ff, 1},
	Range{0x3400, 0x4db5, 1},
	Range{0x4e00, 0x9fc3, 1},
	Range{0xa000, 0xa48c, 1},
	Range{0xa500, 0xa60c, 1},
	Range{0xa610, 0xa61f, 1},
	Range{0xa62a, 0xa62b, 1},
	Range{0xa640, 0xa65f, 1},
	Range{0xa662, 0xa66e, 1},
	Range{0xa67f, 0xa697, 1},
	Range{0xa717, 0xa71f, 1},
	Range{0xa722, 0xa788, 1},
	Range{0xa78b, 0xa78c, 1},
	Range{0xa7fb, 0xa801, 1},
	Range{0xa803, 0xa805, 1},
	Range{0xa807, 0xa80a, 1},
	Range{0xa80c, 0xa822, 1},
	Range{0xa840, 0xa873, 1},
	Range{0xa882, 0xa8b3, 1},
	Range{0xa90a, 0xa925, 1},
	Range{0xa930, 0xa946, 1},
	Range{0xaa00, 0xaa28, 1},
	Range{0xaa40, 0xaa42, 1},
	Range{0xaa44, 0xaa4b, 1},
	Range{0xac00, 0xd7a3, 1},
	Range{0xf900, 0xfa2d, 1},
	Range{0xfa30, 0xfa6a, 1},
	Range{0xfa70, 0xfad9, 1},
	Range{0xfb00, 0xfb06, 1},
	Range{0xfb13, 0xfb17, 1},
	Range{0xfb1d, 0xfb1d, 1},
	Range{0xfb1f, 0xfb28, 1},
	Range{0xfb2a, 0xfb36, 1},
	Range{0xfb38, 0xfb3c, 1},
	Range{0xfb3e, 0xfb3e, 1},
	Range{0xfb40, 0xfb41, 1},
	Range{0xfb43, 0xfb44, 1},
	Range{0xfb46, 0xfbb1, 1},
	Range{0xfbd3, 0xfd3d, 1},
	Range{0xfd50, 0xfd8f, 1},
	Range{0xfd92, 0xfdc7, 1},
	Range{0xfdf0, 0xfdfb, 1},
	Range{0xfe70, 0xfe74, 1},
	Range{0xfe76, 0xfefc, 1},
	Range{0xff21, 0xff3a, 1},
	Range{0xff41, 0xff5a, 1},
	Range{0xff66, 0xffbe, 1},
	Range{0xffc2, 0xffc7, 1},
	Range{0xffca, 0xffcf, 1},
	Range{0xffd2, 0xffd7, 1},
	Range{0xffda, 0xffdc, 1},
	Range{0x10000, 0x1000b, 1},
	Range{0x1000d, 0x10026, 1},
	Range{0x10028, 0x1003a, 1},
	Range{0x1003c, 0x1003d, 1},
	Range{0x1003f, 0x1004d, 1},
	Range{0x10050, 0x1005d, 1},
	Range{0x10080, 0x100fa, 1},
	Range{0x10280, 0x1029c, 1},
	Range{0x102a0, 0x102d0, 1},
	Range{0x10300, 0x1031e, 1},
	Range{0x10330, 0x10340, 1},
	Range{0x10342, 0x10349, 1},
	Range{0x10380, 0x1039d, 1},
	Range{0x103a0, 0x103c3, 1},
	Range{0x103c8, 0x103cf, 1},
	Range{0x10400, 0x1049d, 1},
	Range{0x10800, 0x10805, 1},
	Range{0x10808, 0x10808, 1},
	Range{0x1080a, 0x10835, 1},
	Range{0x10837, 0x10838, 1},
	Range{0x1083c, 0x1083f, 3},
	Range{0x10900, 0x10915, 1},
	Range{0x10920, 0x10939, 1},
	Range{0x10a00, 0x10a00, 1},
	Range{0x10a10, 0x10a13, 1},
	Range{0x10a15, 0x10a17, 1},
	Range{0x10a19, 0x10a33, 1},
	Range{0x12000, 0x1236e, 1},
	Range{0x1d400, 0x1d454, 1},
	Range{0x1d456, 0x1d49c, 1},
	Range{0x1d49e, 0x1d49f, 1},
	Range{0x1d4a2, 0x1d4a2, 1},
	Range{0x1d4a5, 0x1d4a6, 1},
	Range{0x1d4a9, 0x1d4ac, 1},
	Range{0x1d4ae, 0x1d4b9, 1},
	Range{0x1d4bb, 0x1d4bb, 1},
	Range{0x1d4bd, 0x1d4c3, 1},
	Range{0x1d4c5, 0x1d505, 1},
	Range{0x1d507, 0x1d50a, 1},
	Range{0x1d50d, 0x1d514, 1},
	Range{0x1d516, 0x1d51c, 1},
	Range{0x1d51e, 0x1d539, 1},
	Range{0x1d53b, 0x1d53e, 1},
	Range{0x1d540, 0x1d544, 1},
	Range{0x1d546, 0x1d546, 1},
	Range{0x1d54a, 0x1d550, 1},
	Range{0x1d552, 0x1d6a5, 1},
	Range{0x1d6a8, 0x1d6c0, 1},
	Range{0x1d6c2, 0x1d6da, 1},
	Range{0x1d6dc, 0x1d6fa, 1},
	Range{0x1d6fc, 0x1d714, 1},
	Range{0x1d716, 0x1d734, 1},
	Range{0x1d736, 0x1d74e, 1},
	Range{0x1d750, 0x1d76e, 1},
	Range{0x1d770, 0x1d788, 1},
	Range{0x1d78a, 0x1d7a8, 1},
	Range{0x1d7aa, 0x1d7c2, 1},
	Range{0x1d7c4, 0x1d7cb, 1},
	Range{0x20000, 0x2a6d6, 1},
	Range{0x2f800, 0x2fa1d, 1},
}

func Is(ranges []Range, rune int) bool {
	// common case: rune is ASCII or Latin-1
	if rune < 0x100 {
		for i := 0; i < len(ranges); i++ {
			r := ranges[i];
			if rune > r.hi {
				continue;
			}
			if rune < r.lo {
				return false;
			}
			return (rune - r.lo) % r.stride == 0;
		}
		return false;
	}

	// binary search over ranges
	lo := 0;
	hi := len(ranges);
	for lo < hi {
		m := lo + (hi - lo)/2;
		r := ranges[m];
		if r.lo <= rune && rune <= r.hi {
			return (rune - r.lo) % r.stride == 0;
		}
		if rune < r.lo {
			hi = m;
		} else {
			lo = m+1;
		}
	}
	return false;
}

func IsUpper(rune int) bool {
	return Is(Upper, rune);
}

func IsLetter(rune int) bool {
	return Is(Letter, rune);
}

