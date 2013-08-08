// +build amd64
// errorcheck

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Issue 2444
// Issue 4666: issue with arrays of exactly 4GB.

package main

var z [10<<20]byte

func main() { // ERROR "stack frame too large"
	// seq 1 206 | sed 's/.*/	var x& [10<<20]byte; z = x&/'
	var x1 [10<<20]byte; z = x1
	var x2 [10<<20]byte; z = x2
	var x3 [10<<20]byte; z = x3
	var x4 [10<<20]byte; z = x4
	var x5 [10<<20]byte; z = x5
	var x6 [10<<20]byte; z = x6
	var x7 [10<<20]byte; z = x7
	var x8 [10<<20]byte; z = x8
	var x9 [10<<20]byte; z = x9
	var x10 [10<<20]byte; z = x10
	var x11 [10<<20]byte; z = x11
	var x12 [10<<20]byte; z = x12
	var x13 [10<<20]byte; z = x13
	var x14 [10<<20]byte; z = x14
	var x15 [10<<20]byte; z = x15
	var x16 [10<<20]byte; z = x16
	var x17 [10<<20]byte; z = x17
	var x18 [10<<20]byte; z = x18
	var x19 [10<<20]byte; z = x19
	var x20 [10<<20]byte; z = x20
	var x21 [10<<20]byte; z = x21
	var x22 [10<<20]byte; z = x22
	var x23 [10<<20]byte; z = x23
	var x24 [10<<20]byte; z = x24
	var x25 [10<<20]byte; z = x25
	var x26 [10<<20]byte; z = x26
	var x27 [10<<20]byte; z = x27
	var x28 [10<<20]byte; z = x28
	var x29 [10<<20]byte; z = x29
	var x30 [10<<20]byte; z = x30
	var x31 [10<<20]byte; z = x31
	var x32 [10<<20]byte; z = x32
	var x33 [10<<20]byte; z = x33
	var x34 [10<<20]byte; z = x34
	var x35 [10<<20]byte; z = x35
	var x36 [10<<20]byte; z = x36
	var x37 [10<<20]byte; z = x37
	var x38 [10<<20]byte; z = x38
	var x39 [10<<20]byte; z = x39
	var x40 [10<<20]byte; z = x40
	var x41 [10<<20]byte; z = x41
	var x42 [10<<20]byte; z = x42
	var x43 [10<<20]byte; z = x43
	var x44 [10<<20]byte; z = x44
	var x45 [10<<20]byte; z = x45
	var x46 [10<<20]byte; z = x46
	var x47 [10<<20]byte; z = x47
	var x48 [10<<20]byte; z = x48
	var x49 [10<<20]byte; z = x49
	var x50 [10<<20]byte; z = x50
	var x51 [10<<20]byte; z = x51
	var x52 [10<<20]byte; z = x52
	var x53 [10<<20]byte; z = x53
	var x54 [10<<20]byte; z = x54
	var x55 [10<<20]byte; z = x55
	var x56 [10<<20]byte; z = x56
	var x57 [10<<20]byte; z = x57
	var x58 [10<<20]byte; z = x58
	var x59 [10<<20]byte; z = x59
	var x60 [10<<20]byte; z = x60
	var x61 [10<<20]byte; z = x61
	var x62 [10<<20]byte; z = x62
	var x63 [10<<20]byte; z = x63
	var x64 [10<<20]byte; z = x64
	var x65 [10<<20]byte; z = x65
	var x66 [10<<20]byte; z = x66
	var x67 [10<<20]byte; z = x67
	var x68 [10<<20]byte; z = x68
	var x69 [10<<20]byte; z = x69
	var x70 [10<<20]byte; z = x70
	var x71 [10<<20]byte; z = x71
	var x72 [10<<20]byte; z = x72
	var x73 [10<<20]byte; z = x73
	var x74 [10<<20]byte; z = x74
	var x75 [10<<20]byte; z = x75
	var x76 [10<<20]byte; z = x76
	var x77 [10<<20]byte; z = x77
	var x78 [10<<20]byte; z = x78
	var x79 [10<<20]byte; z = x79
	var x80 [10<<20]byte; z = x80
	var x81 [10<<20]byte; z = x81
	var x82 [10<<20]byte; z = x82
	var x83 [10<<20]byte; z = x83
	var x84 [10<<20]byte; z = x84
	var x85 [10<<20]byte; z = x85
	var x86 [10<<20]byte; z = x86
	var x87 [10<<20]byte; z = x87
	var x88 [10<<20]byte; z = x88
	var x89 [10<<20]byte; z = x89
	var x90 [10<<20]byte; z = x90
	var x91 [10<<20]byte; z = x91
	var x92 [10<<20]byte; z = x92
	var x93 [10<<20]byte; z = x93
	var x94 [10<<20]byte; z = x94
	var x95 [10<<20]byte; z = x95
	var x96 [10<<20]byte; z = x96
	var x97 [10<<20]byte; z = x97
	var x98 [10<<20]byte; z = x98
	var x99 [10<<20]byte; z = x99
	var x100 [10<<20]byte; z = x100
	var x101 [10<<20]byte; z = x101
	var x102 [10<<20]byte; z = x102
	var x103 [10<<20]byte; z = x103
	var x104 [10<<20]byte; z = x104
	var x105 [10<<20]byte; z = x105
	var x106 [10<<20]byte; z = x106
	var x107 [10<<20]byte; z = x107
	var x108 [10<<20]byte; z = x108
	var x109 [10<<20]byte; z = x109
	var x110 [10<<20]byte; z = x110
	var x111 [10<<20]byte; z = x111
	var x112 [10<<20]byte; z = x112
	var x113 [10<<20]byte; z = x113
	var x114 [10<<20]byte; z = x114
	var x115 [10<<20]byte; z = x115
	var x116 [10<<20]byte; z = x116
	var x117 [10<<20]byte; z = x117
	var x118 [10<<20]byte; z = x118
	var x119 [10<<20]byte; z = x119
	var x120 [10<<20]byte; z = x120
	var x121 [10<<20]byte; z = x121
	var x122 [10<<20]byte; z = x122
	var x123 [10<<20]byte; z = x123
	var x124 [10<<20]byte; z = x124
	var x125 [10<<20]byte; z = x125
	var x126 [10<<20]byte; z = x126
	var x127 [10<<20]byte; z = x127
	var x128 [10<<20]byte; z = x128
	var x129 [10<<20]byte; z = x129
	var x130 [10<<20]byte; z = x130
	var x131 [10<<20]byte; z = x131
	var x132 [10<<20]byte; z = x132
	var x133 [10<<20]byte; z = x133
	var x134 [10<<20]byte; z = x134
	var x135 [10<<20]byte; z = x135
	var x136 [10<<20]byte; z = x136
	var x137 [10<<20]byte; z = x137
	var x138 [10<<20]byte; z = x138
	var x139 [10<<20]byte; z = x139
	var x140 [10<<20]byte; z = x140
	var x141 [10<<20]byte; z = x141
	var x142 [10<<20]byte; z = x142
	var x143 [10<<20]byte; z = x143
	var x144 [10<<20]byte; z = x144
	var x145 [10<<20]byte; z = x145
	var x146 [10<<20]byte; z = x146
	var x147 [10<<20]byte; z = x147
	var x148 [10<<20]byte; z = x148
	var x149 [10<<20]byte; z = x149
	var x150 [10<<20]byte; z = x150
	var x151 [10<<20]byte; z = x151
	var x152 [10<<20]byte; z = x152
	var x153 [10<<20]byte; z = x153
	var x154 [10<<20]byte; z = x154
	var x155 [10<<20]byte; z = x155
	var x156 [10<<20]byte; z = x156
	var x157 [10<<20]byte; z = x157
	var x158 [10<<20]byte; z = x158
	var x159 [10<<20]byte; z = x159
	var x160 [10<<20]byte; z = x160
	var x161 [10<<20]byte; z = x161
	var x162 [10<<20]byte; z = x162
	var x163 [10<<20]byte; z = x163
	var x164 [10<<20]byte; z = x164
	var x165 [10<<20]byte; z = x165
	var x166 [10<<20]byte; z = x166
	var x167 [10<<20]byte; z = x167
	var x168 [10<<20]byte; z = x168
	var x169 [10<<20]byte; z = x169
	var x170 [10<<20]byte; z = x170
	var x171 [10<<20]byte; z = x171
	var x172 [10<<20]byte; z = x172
	var x173 [10<<20]byte; z = x173
	var x174 [10<<20]byte; z = x174
	var x175 [10<<20]byte; z = x175
	var x176 [10<<20]byte; z = x176
	var x177 [10<<20]byte; z = x177
	var x178 [10<<20]byte; z = x178
	var x179 [10<<20]byte; z = x179
	var x180 [10<<20]byte; z = x180
	var x181 [10<<20]byte; z = x181
	var x182 [10<<20]byte; z = x182
	var x183 [10<<20]byte; z = x183
	var x184 [10<<20]byte; z = x184
	var x185 [10<<20]byte; z = x185
	var x186 [10<<20]byte; z = x186
	var x187 [10<<20]byte; z = x187
	var x188 [10<<20]byte; z = x188
	var x189 [10<<20]byte; z = x189
	var x190 [10<<20]byte; z = x190
	var x191 [10<<20]byte; z = x191
	var x192 [10<<20]byte; z = x192
	var x193 [10<<20]byte; z = x193
	var x194 [10<<20]byte; z = x194
	var x195 [10<<20]byte; z = x195
	var x196 [10<<20]byte; z = x196
	var x197 [10<<20]byte; z = x197
	var x198 [10<<20]byte; z = x198
	var x199 [10<<20]byte; z = x199
	var x200 [10<<20]byte; z = x200
	var x201 [10<<20]byte; z = x201
	var x202 [10<<20]byte; z = x202
	var x203 [10<<20]byte; z = x203
	var x204 [10<<20]byte; z = x204
	var x205 [10<<20]byte; z = x205
	var x206 [10<<20]byte; z = x206
}
