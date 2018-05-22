// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This input extends auto-generated amd64enc.s test suite
// with manually added tests.

#include "../../../../../runtime/textflag.h"

TEXT asmtest(SB),DUPOK|NOSPLIT,$0
	// AVX2GATHER: basic combinations.
	VPGATHERDQ Y2, (BP)(X7*2), Y1           // c4e2ed904c7d00
	VPGATHERDQ X12, (R13)(X14*2), X11       // c40299905c7500
	VPGATHERDQ Y12, (R13)(X14*2), Y11       // c4029d905c7500
	VPGATHERDQ Y0, 8(X4*1), Y6              // c4e2fd90342508000000
	VPGATHERDQ Y0, -8(X4*1), Y6             // c4e2fd903425f8ffffff
	VPGATHERDQ Y0, 0(X4*1), Y6              // c4e2fd90342500000000
	VPGATHERDQ Y0, 664(X4*1), Y6            // c4e2fd90342598020000
	VPGATHERDQ Y0, 8(X4*8), Y6              // c4e2fd9034e508000000
	VPGATHERDQ Y0, -8(X4*8), Y6             // c4e2fd9034e5f8ffffff
	VPGATHERDQ Y0, 0(X4*8), Y6              // c4e2fd9034e500000000
	VPGATHERDQ Y0, 664(X4*8), Y6            // c4e2fd9034e598020000
	VPGATHERDQ Y0, 8(X14*1), Y6             // c4a2fd90343508000000
	VPGATHERDQ Y0, -8(X14*1), Y6            // c4a2fd903435f8ffffff
	VPGATHERDQ Y0, 0(X14*1), Y6             // c4a2fd90343500000000
	VPGATHERDQ Y0, 664(X14*1), Y6           // c4a2fd90343598020000
	VPGATHERDQ Y0, 8(X14*8), Y6             // c4a2fd9034f508000000
	VPGATHERDQ Y0, -8(X14*8), Y6            // c4a2fd9034f5f8ffffff
	VPGATHERDQ Y0, 0(X14*8), Y6             // c4a2fd9034f500000000
	VPGATHERDQ Y0, 664(X14*8), Y6           // c4a2fd9034f598020000
	VPGATHERDQ X2, (BP)(X7*2), X1           // c4e2e9904c7d00
	VPGATHERDQ Y2, (BP)(X7*2), Y1           // c4e2ed904c7d00
	VPGATHERDQ X12, (R13)(X14*2), X11       // c40299905c7500
	VPGATHERDQ Y12, (R13)(X14*2), Y11       // c4029d905c7500
	VPGATHERDQ Y0, 8(X4*1), Y6              // c4e2fd90342508000000
	VPGATHERDQ Y0, -8(X4*1), Y6             // c4e2fd903425f8ffffff
	VPGATHERDQ Y0, 0(X4*1), Y6              // c4e2fd90342500000000
	VPGATHERDQ Y0, 664(X4*1), Y6            // c4e2fd90342598020000
	VPGATHERDQ Y0, 8(X4*8), Y6              // c4e2fd9034e508000000
	VPGATHERDQ Y0, -8(X4*8), Y6             // c4e2fd9034e5f8ffffff
	VPGATHERDQ Y0, 0(X4*8), Y6              // c4e2fd9034e500000000
	VPGATHERDQ Y0, 664(X4*8), Y6            // c4e2fd9034e598020000
	VPGATHERDQ Y0, 8(X14*1), Y6             // c4a2fd90343508000000
	VPGATHERDQ Y0, -8(X14*1), Y6            // c4a2fd903435f8ffffff
	VPGATHERDQ Y0, 0(X14*1), Y6             // c4a2fd90343500000000
	VPGATHERDQ Y0, 664(X14*1), Y6           // c4a2fd90343598020000
	VPGATHERDQ Y0, 8(X14*8), Y6             // c4a2fd9034f508000000
	VPGATHERDQ Y0, -8(X14*8), Y6            // c4a2fd9034f5f8ffffff
	VPGATHERDQ Y0, 0(X14*8), Y6             // c4a2fd9034f500000000
	VPGATHERDQ Y0, 664(X14*8), Y6           // c4a2fd9034f598020000
	VPGATHERQQ X2, (BP)(X7*2), X1           // c4e2e9914c7d00
	VPGATHERQQ Y2, (BP)(Y7*2), Y1           // c4e2ed914c7d00
	VPGATHERQQ X12, (R13)(X14*2), X11       // c40299915c7500
	VPGATHERQQ Y12, (R13)(Y14*2), Y11       // c4029d915c7500
	VPGATHERQQ X2, (BP)(X7*2), X1           // c4e2e9914c7d00
	VPGATHERQQ Y2, (BP)(Y7*2), Y1           // c4e2ed914c7d00
	VPGATHERQQ X12, (R13)(X14*2), X11       // c40299915c7500
	VPGATHERQQ Y12, (R13)(Y14*2), Y11       // c4029d915c7500
	VGATHERDPD X2, (BP)(X7*2), X1           // c4e2e9924c7d00
	VGATHERDPD Y2, (BP)(X7*2), Y1           // c4e2ed924c7d00
	VGATHERDPD X12, (R13)(X14*2), X11       // c40299925c7500
	VGATHERDPD Y12, (R13)(X14*2), Y11       // c4029d925c7500
	VGATHERDPD Y0, 8(X4*1), Y6              // c4e2fd92342508000000
	VGATHERDPD Y0, -8(X4*1), Y6             // c4e2fd923425f8ffffff
	VGATHERDPD Y0, 0(X4*1), Y6              // c4e2fd92342500000000
	VGATHERDPD Y0, 664(X4*1), Y6            // c4e2fd92342598020000
	VGATHERDPD Y0, 8(X4*8), Y6              // c4e2fd9234e508000000
	VGATHERDPD Y0, -8(X4*8), Y6             // c4e2fd9234e5f8ffffff
	VGATHERDPD Y0, 0(X4*8), Y6              // c4e2fd9234e500000000
	VGATHERDPD Y0, 664(X4*8), Y6            // c4e2fd9234e598020000
	VGATHERDPD Y0, 8(X14*1), Y6             // c4a2fd92343508000000
	VGATHERDPD Y0, -8(X14*1), Y6            // c4a2fd923435f8ffffff
	VGATHERDPD Y0, 0(X14*1), Y6             // c4a2fd92343500000000
	VGATHERDPD Y0, 664(X14*1), Y6           // c4a2fd92343598020000
	VGATHERDPD Y0, 8(X14*8), Y6             // c4a2fd9234f508000000
	VGATHERDPD Y0, -8(X14*8), Y6            // c4a2fd9234f5f8ffffff
	VGATHERDPD Y0, 0(X14*8), Y6             // c4a2fd9234f500000000
	VGATHERDPD Y0, 664(X14*8), Y6           // c4a2fd9234f598020000
	VGATHERDPD X2, (BP)(X7*2), X1           // c4e2e9924c7d00
	VGATHERDPD Y2, (BP)(X7*2), Y1           // c4e2ed924c7d00
	VGATHERDPD X12, (R13)(X14*2), X11       // c40299925c7500
	VGATHERDPD Y12, (R13)(X14*2), Y11       // c4029d925c7500
	VGATHERDPD Y0, 8(X4*1), Y6              // c4e2fd92342508000000
	VGATHERDPD Y0, -8(X4*1), Y6             // c4e2fd923425f8ffffff
	VGATHERDPD Y0, 0(X4*1), Y6              // c4e2fd92342500000000
	VGATHERDPD Y0, 664(X4*1), Y6            // c4e2fd92342598020000
	VGATHERDPD Y0, 8(X4*8), Y6              // c4e2fd9234e508000000
	VGATHERDPD Y0, -8(X4*8), Y6             // c4e2fd9234e5f8ffffff
	VGATHERDPD Y0, 0(X4*8), Y6              // c4e2fd9234e500000000
	VGATHERDPD Y0, 664(X4*8), Y6            // c4e2fd9234e598020000
	VGATHERDPD Y0, 8(X14*1), Y6             // c4a2fd92343508000000
	VGATHERDPD Y0, -8(X14*1), Y6            // c4a2fd923435f8ffffff
	VGATHERDPD Y0, 0(X14*1), Y6             // c4a2fd92343500000000
	VGATHERDPD Y0, 664(X14*1), Y6           // c4a2fd92343598020000
	VGATHERDPD Y0, 8(X14*8), Y6             // c4a2fd9234f508000000
	VGATHERDPD Y0, -8(X14*8), Y6            // c4a2fd9234f5f8ffffff
	VGATHERDPD Y0, 0(X14*8), Y6             // c4a2fd9234f500000000
	VGATHERDPD Y0, 664(X14*8), Y6           // c4a2fd9234f598020000
	VGATHERQPD X2, (BP)(X7*2), X1           // c4e2e9934c7d00
	VGATHERQPD Y2, (BP)(Y7*2), Y1           // c4e2ed934c7d00
	VGATHERQPD X12, (R13)(X14*2), X11       // c40299935c7500
	VGATHERQPD Y12, (R13)(Y14*2), Y11       // c4029d935c7500
	VGATHERQPD X2, (BP)(X7*2), X1           // c4e2e9934c7d00
	VGATHERQPD Y2, (BP)(Y7*2), Y1           // c4e2ed934c7d00
	VGATHERQPD X12, (R13)(X14*2), X11       // c40299935c7500
	VGATHERQPD Y12, (R13)(Y14*2), Y11       // c4029d935c7500
	VGATHERDPS X2, (BP)(X7*2), X1           // c4e269924c7d00
	VGATHERDPS Y2, (BP)(Y7*2), Y1           // c4e26d924c7d00
	VGATHERDPS X12, (R13)(X14*2), X11       // c40219925c7500
	VGATHERDPS Y12, (R13)(Y14*2), Y11       // c4021d925c7500
	VGATHERDPS X3, 8(X4*1), X6              // c4e26192342508000000
	VGATHERDPS X3, -8(X4*1), X6             // c4e261923425f8ffffff
	VGATHERDPS X3, 0(X4*1), X6              // c4e26192342500000000
	VGATHERDPS X3, 664(X4*1), X6            // c4e26192342598020000
	VGATHERDPS X3, 8(X4*8), X6              // c4e2619234e508000000
	VGATHERDPS X3, -8(X4*8), X6             // c4e2619234e5f8ffffff
	VGATHERDPS X3, 0(X4*8), X6              // c4e2619234e500000000
	VGATHERDPS X3, 664(X4*8), X6            // c4e2619234e598020000
	VGATHERDPS X3, 8(X14*1), X6             // c4a26192343508000000
	VGATHERDPS X3, -8(X14*1), X6            // c4a261923435f8ffffff
	VGATHERDPS X3, 0(X14*1), X6             // c4a26192343500000000
	VGATHERDPS X3, 664(X14*1), X6           // c4a26192343598020000
	VGATHERDPS X3, 8(X14*8), X6             // c4a2619234f508000000
	VGATHERDPS X3, -8(X14*8), X6            // c4a2619234f5f8ffffff
	VGATHERDPS X3, 0(X14*8), X6             // c4a2619234f500000000
	VGATHERDPS X3, 664(X14*8), X6           // c4a2619234f598020000
	VGATHERDPS X2, (BP)(X7*2), X1           // c4e269924c7d00
	VGATHERDPS Y2, (BP)(Y7*2), Y1           // c4e26d924c7d00
	VGATHERDPS X12, (R13)(X14*2), X11       // c40219925c7500
	VGATHERDPS Y12, (R13)(Y14*2), Y11       // c4021d925c7500
	VGATHERDPS X5, 8(X4*1), X6              // c4e25192342508000000
	VGATHERDPS X3, -8(X4*1), X6             // c4e261923425f8ffffff
	VGATHERDPS X3, 0(X4*1), X6              // c4e26192342500000000
	VGATHERDPS X3, 664(X4*1), X6            // c4e26192342598020000
	VGATHERDPS X3, 8(X4*8), X6              // c4e2619234e508000000
	VGATHERDPS X3, -8(X4*8), X6             // c4e2619234e5f8ffffff
	VGATHERDPS X3, 0(X4*8), X6              // c4e2619234e500000000
	VGATHERDPS X3, 664(X4*8), X6            // c4e2619234e598020000
	VGATHERDPS X3, 8(X14*1), X6             // c4a26192343508000000
	VGATHERDPS X3, -8(X14*1), X6            // c4a261923435f8ffffff
	VGATHERDPS X3, 0(X14*1), X6             // c4a26192343500000000
	VGATHERDPS X3, 664(X14*1), X6           // c4a26192343598020000
	VGATHERDPS X3, 8(X14*8), X6             // c4a2619234f508000000
	VGATHERDPS X3, -8(X14*8), X6            // c4a2619234f5f8ffffff
	VGATHERDPS X3, 0(X14*8), X6             // c4a2619234f500000000
	VGATHERDPS X3, 664(X14*8), X6           // c4a2619234f598020000
	VGATHERQPS X2, (BP)(X7*2), X1           // c4e269934c7d00
	VGATHERQPS X2, (BP)(Y7*2), X1           // c4e26d934c7d00
	VGATHERQPS X12, (R13)(X14*2), X11       // c40219935c7500
	VGATHERQPS X12, (R13)(Y14*2), X11       // c4021d935c7500
	VGATHERQPS X2, (BP)(X7*2), X1           // c4e269934c7d00
	VGATHERQPS X2, (BP)(Y7*2), X1           // c4e26d934c7d00
	VGATHERQPS X12, (R13)(X14*2), X11       // c40219935c7500
	VGATHERQPS X12, (R13)(Y14*2), X11       // c4021d935c7500
	VPGATHERDD X2, (BP)(X7*2), X1           // c4e269904c7d00
	VPGATHERDD Y2, (BP)(Y7*2), Y1           // c4e26d904c7d00
	VPGATHERDD X12, (R13)(X14*2), X11       // c40219905c7500
	VPGATHERDD Y12, (R13)(Y14*2), Y11       // c4021d905c7500
	VPGATHERDD X3, 8(X4*1), X6              // c4e26190342508000000
	VPGATHERDD X3, -8(X4*1), X6             // c4e261903425f8ffffff
	VPGATHERDD X3, 0(X4*1), X6              // c4e26190342500000000
	VPGATHERDD X3, 664(X4*1), X6            // c4e26190342598020000
	VPGATHERDD X3, 8(X4*8), X6              // c4e2619034e508000000
	VPGATHERDD X3, -8(X4*8), X6             // c4e2619034e5f8ffffff
	VPGATHERDD X3, 0(X4*8), X6              // c4e2619034e500000000
	VPGATHERDD X3, 664(X4*8), X6            // c4e2619034e598020000
	VPGATHERDD X3, 8(X14*1), X6             // c4a26190343508000000
	VPGATHERDD X3, -8(X14*1), X6            // c4a261903435f8ffffff
	VPGATHERDD X3, 0(X14*1), X6             // c4a26190343500000000
	VPGATHERDD X3, 664(X14*1), X6           // c4a26190343598020000
	VPGATHERDD X3, 8(X14*8), X6             // c4a2619034f508000000
	VPGATHERDD X3, -8(X14*8), X6            // c4a2619034f5f8ffffff
	VPGATHERDD X3, 0(X14*8), X6             // c4a2619034f500000000
	VPGATHERDD X3, 664(X14*8), X6           // c4a2619034f598020000
	VPGATHERDD X2, (BP)(X7*2), X1           // c4e269904c7d00
	VPGATHERDD Y2, (BP)(Y7*2), Y1           // c4e26d904c7d00
	VPGATHERDD X12, (R13)(X14*2), X11       // c40219905c7500
	VPGATHERDD Y12, (R13)(Y14*2), Y11       // c4021d905c7500
	VPGATHERDD X3, 8(X4*1), X6              // c4e26190342508000000
	VPGATHERDD X3, -8(X4*1), X6             // c4e261903425f8ffffff
	VPGATHERDD X3, 0(X4*1), X6              // c4e26190342500000000
	VPGATHERDD X3, 664(X4*1), X6            // c4e26190342598020000
	VPGATHERDD X3, 8(X4*8), X6              // c4e2619034e508000000
	VPGATHERDD X3, -8(X4*8), X6             // c4e2619034e5f8ffffff
	VPGATHERDD X3, 0(X4*8), X6              // c4e2619034e500000000
	VPGATHERDD X3, 664(X4*8), X6            // c4e2619034e598020000
	VPGATHERDD X3, 8(X14*1), X6             // c4a26190343508000000
	VPGATHERDD X3, -8(X14*1), X6            // c4a261903435f8ffffff
	VPGATHERDD X3, 0(X14*1), X6             // c4a26190343500000000
	VPGATHERDD X3, 664(X14*1), X6           // c4a26190343598020000
	VPGATHERDD X3, 8(X14*8), X6             // c4a2619034f508000000
	VPGATHERDD X3, -8(X14*8), X6            // c4a2619034f5f8ffffff
	VPGATHERDD X3, 0(X14*8), X6             // c4a2619034f500000000
	VPGATHERDD X3, 664(X14*8), X6           // c4a2619034f598020000
	VPGATHERQD X2, (BP)(X7*2), X1           // c4e269914c7d00
	VPGATHERQD X2, (BP)(Y7*2), X1           // c4e26d914c7d00
	VPGATHERQD X12, (R13)(X14*2), X11       // c40219915c7500
	VPGATHERQD X12, (R13)(Y14*2), X11       // c4021d915c7500
	VPGATHERQD X2, (BP)(X7*2), X1           // c4e269914c7d00
	VPGATHERQD X2, (BP)(Y7*2), X1           // c4e26d914c7d00
	VPGATHERQD X12, (R13)(X14*2), X11       // c40219915c7500
	VPGATHERQD X12, (R13)(Y14*2), X11       // c4021d915c7500
	VPGATHERQQ X0, 0(X1*1), X2              // c4e2f991140d00000000
	VPGATHERQQ Y0, 0(Y1*1), Y2              // c4e2fd91140d00000000
	VPGATHERQQ X8, 0(X9*1), X10             // c422b991140d00000000
	VPGATHERQQ Y8, 0(Y9*1), Y10             // c422bd91140d00000000
	VPGATHERQQ X0, 0(X1*4), X2              // c4e2f991148d00000000
	VPGATHERQQ Y0, 0(Y1*4), Y2              // c4e2fd91148d00000000
	VPGATHERQQ X8, 0(X9*4), X10             // c422b991148d00000000
	VPGATHERQQ Y8, 0(Y9*4), Y10             // c422bd91148d00000000
	// AVX2GATHER: test SP/BP base with different displacements.
	VPGATHERQQ X0, (SP)(X1*1), X2           // c4e2f991140c
	VPGATHERQQ X0, 16(SP)(X1*1), X2         // c4e2f991540c10
	VPGATHERQQ X0, 512(SP)(X1*1), X2        // c4e2f991940c00020000
	VPGATHERQQ X0, (R12)(X1*1), X2          // c4c2f991140c
	VPGATHERQQ X0, 16(R12)(X1*1), X2        // c4c2f991540c10
	VPGATHERQQ X0, 512(R12)(X1*1), X2       // c4c2f991940c00020000
	VPGATHERQQ X0, (BP)(X1*1), X2           // c4e2f991540d00
	VPGATHERQQ X0, 16(BP)(X1*1), X2         // c4e2f991540d10
	VPGATHERQQ X0, 512(BP)(X1*1), X2        // c4e2f991940d00020000
	VPGATHERQQ X0, (R13)(X1*1), X2          // c4c2f991540d00
	VPGATHERQQ X0, 16(R13)(X1*1), X2        // c4c2f991540d10
	VPGATHERQQ X0, 512(R13)(X1*1), X2       // c4c2f991940d00020000
	VPGATHERQQ Y0, (SP)(Y1*1), Y2           // c4e2fd91140c
	VPGATHERQQ Y0, 16(SP)(Y1*1), Y2         // c4e2fd91540c10
	VPGATHERQQ Y0, 512(SP)(Y1*1), Y2        // c4e2fd91940c00020000
	VPGATHERQQ Y0, (R12)(Y1*1), Y2          // c4c2fd91140c
	VPGATHERQQ Y0, 16(R12)(Y1*1), Y2        // c4c2fd91540c10
	VPGATHERQQ Y0, 512(R12)(Y1*1), Y2       // c4c2fd91940c00020000
	VPGATHERQQ Y0, (BP)(Y1*1), Y2           // c4e2fd91540d00
	VPGATHERQQ Y0, 16(BP)(Y1*1), Y2         // c4e2fd91540d10
	VPGATHERQQ Y0, 512(BP)(Y1*1), Y2        // c4e2fd91940d00020000
	VPGATHERQQ Y0, (R13)(Y1*1), Y2          // c4c2fd91540d00
	VPGATHERQQ Y0, 16(R13)(Y1*1), Y2        // c4c2fd91540d10
	VPGATHERQQ Y0, 512(R13)(Y1*1), Y2       // c4c2fd91940d00020000
	// Test low-8 register for /is4 "hr" operand.
	VPBLENDVB X0, (BX), X1, X2              // c4e3714c1300
	// <XMM0>/Yxr0 tests.
	SHA256RNDS2 X0, (BX), X2   // 0f38cb13
	SHA256RNDS2 X0, (R11), X2  // 410f38cb13
	SHA256RNDS2 X0, X2, X2     // 0f38cbd2
	SHA256RNDS2 X0, X11, X2    // 410f38cbd3
	SHA256RNDS2 X0, (BX), X11  // 440f38cb1b
	SHA256RNDS2 X0, (R11), X11 // 450f38cb1b
	SHA256RNDS2 X0, X2, X11    // 440f38cbda
	SHA256RNDS2 X0, X11, X11   // 450f38cbdb
	// Rest SHA instructions tests.
	SHA1MSG1 (BX), X2        // 0f38c913
	SHA1MSG1 (R11), X2       // 410f38c913
	SHA1MSG1 X2, X2          // 0f38c9d2
	SHA1MSG1 X11, X2         // 410f38c9d3
	SHA1MSG1 (BX), X11       // 440f38c91b
	SHA1MSG1 (R11), X11      // 450f38c91b
	SHA1MSG1 X2, X11         // 440f38c9da
	SHA1MSG1 X11, X11        // 450f38c9db
	SHA1MSG2 (BX), X2        // 0f38ca13
	SHA1MSG2 (R11), X2       // 410f38ca13
	SHA1MSG2 X2, X2          // 0f38cad2
	SHA1MSG2 X11, X2         // 410f38cad3
	SHA1MSG2 (BX), X11       // 440f38ca1b
	SHA1MSG2 (R11), X11      // 450f38ca1b
	SHA1MSG2 X2, X11         // 440f38cada
	SHA1MSG2 X11, X11        // 450f38cadb
	SHA1NEXTE (BX), X2       // 0f38c813
	SHA1NEXTE (R11), X2      // 410f38c813
	SHA1NEXTE X2, X2         // 0f38c8d2
	SHA1NEXTE X11, X2        // 410f38c8d3
	SHA1NEXTE (BX), X11      // 440f38c81b
	SHA1NEXTE (R11), X11     // 450f38c81b
	SHA1NEXTE X2, X11        // 440f38c8da
	SHA1NEXTE X11, X11       // 450f38c8db
	SHA1RNDS4 $0, (BX), X2   // 0f3acc1300
	SHA1RNDS4 $0, (R11), X2  // 410f3acc1300
	SHA1RNDS4 $1, X2, X2     // 0f3accd201
	SHA1RNDS4 $1, X11, X2    // 410f3accd301
	SHA1RNDS4 $2, (BX), X11  // 440f3acc1b02
	SHA1RNDS4 $2, (R11), X11 // 450f3acc1b02
	SHA1RNDS4 $3, X2, X11    // 440f3accda03
	SHA1RNDS4 $3, X11, X11   // 450f3accdb03
	SHA256MSG1 (BX), X2      // 0f38cc13
	SHA256MSG1 (R11), X2     // 410f38cc13
	SHA256MSG1 X2, X2        // 0f38ccd2
	SHA256MSG1 X11, X2       // 410f38ccd3
	SHA256MSG1 (BX), X11     // 440f38cc1b
	SHA256MSG1 (R11), X11    // 450f38cc1b
	SHA256MSG1 X2, X11       // 440f38ccda
	SHA256MSG1 X11, X11      // 450f38ccdb
	SHA256MSG2 (BX), X2      // 0f38cd13
	SHA256MSG2 (R11), X2     // 410f38cd13
	SHA256MSG2 X2, X2        // 0f38cdd2
	SHA256MSG2 X11, X2       // 410f38cdd3
	SHA256MSG2 (BX), X11     // 440f38cd1b
	SHA256MSG2 (R11), X11    // 450f38cd1b
	SHA256MSG2 X2, X11       // 440f38cdda
	SHA256MSG2 X11, X11      // 450f38cddb
	// Test VPERMQ with both uint8 and int8 immediate args
	VPERMQ $-40, Y8, Y8 // c443fd00c0d8
	VPERMQ $216, Y8, Y8 // c443fd00c0d8
	// Test that VPERMPD that shares ytab list with VPERMQ continues to work too.
	VPERMPD $-40, Y7, Y7 // c4e3fd01ffd8
	VPERMPD $216, Y7, Y7 // c4e3fd01ffd8
	// Check that LEAL is permitted to use overflowing offset.
	LEAL 2400959708(BP)(R10*1), BP // 428dac15dcbc1b8f
	LEAL 3395469782(AX)(R10*1), AX // 428d8410d6c162ca
	// Make sure MOV CR/DR continues to work after changing it's movtabs.
	MOVQ CR0, AX // 0f20c0
	MOVQ CR0, DX // 0f20c2
	MOVQ CR4, DI // 0f20e7
	MOVQ AX, CR0 // 0f22c0
	MOVQ DX, CR0 // 0f22c2
	MOVQ DI, CR4 // 0f22e7
	MOVQ DR0, AX // 0f21c0
	MOVQ DR6, DX // 0f21f2
	MOVQ DR7, SI // 0f21fe
	// Test other movtab entries.
	PUSHQ GS // 0fa8
	PUSHQ FS // 0fa0
	POPQ FS  // 0fa1
	POPQ GS  // 0fa9
	// All instructions below semantically have unsigned operands,
	// but previous assembler permitted negative arguments.
	// This behavior is preserved for compatibility reasons.
	VPSHUFD $-79, X7, X7         // c5f970ffb1
	RORXL $-1, (AX), DX          // c4e37bf010ff
	RORXQ $-1, (AX), DX          // c4e3fbf010ff
	VPSHUFD $-1, X1, X2          // c5f970d1ff
	VPSHUFD $-1, Y1, Y2          // c5fd70d1ff
	VPSHUFHW $-1, X1, X2         // c5fa70d1ff
	VPSHUFHW $-1, Y1, Y2         // c5fe70d1ff
	VPSHUFLW $-1, X1, X2         // c5fb70d1ff
	VPSHUFLW $-1, Y1, Y2         // c5ff70d1ff
	VROUNDPD $-1, X1, X2         // c4e37909d1ff
	VROUNDPS $-1, Y1, Y2         // c4e37d08d1ff
	VPSLLD $-1, X1, X2           // c5e972f1ff
	VPSLLD $-1, Y1, Y2           // c5ed72f1ff
	VPSLLDQ $-1, X1, X2          // c5e973f9ff
	VPSLLDQ $-1, Y1, Y2          // c5ed73f9ff
	VPSLLQ $-1, X1, X2           // c5e973f1ff
	VPSLLQ $-1, Y1, Y2           // c5ed73f1ff
	VPSRLD $-1, X1, X2           // c5e972d1ff
	VPSRLD $-1, Y1, Y2           // c5ed72d1ff
	VPSRLDQ $-1, X1, X2          // c5e973d9ff
	VPSRLDQ $-1, Y1, Y2          // c5ed73d9ff
	VPSRLQ $-1, X1, X2           // c5e973d1ff
	VPSRLQ $-1, Y1, Y2           // c5ed73d1ff
	VPEXTRW $-1, X1, (AX)        // c4e3791508ff
	VPEXTRW $-1, X1, AX          // c4e37915c8ff
	VEXTRACTF128 $-1, Y1, X2     // c4e37d19caff
	VEXTRACTI128 $-1, Y1, X2     // c4e37d39caff
	VAESKEYGENASSIST $-1, X1, X2 // c4e379dfd1ff
	VPCMPESTRI $-1, X1, X2       // c4e37961d1ff
	VPCMPESTRM $-1, X1, X2       // c4e37960d1ff
	VPCMPISTRI $-1, X1, X2       // c4e37963d1ff
	VPCMPISTRM $-1, X1, X2       // c4e37962d1ff
	VPERMPD $-1, Y1, Y2          // c4e3fd01d1ff
	VPERMILPD $-1, X1, X2        // c4e37905d1ff
	VPERMILPD $-1, Y1, Y2        // c4e37d05d1ff
	VPERMILPS $-1, X1, X2        // c4e37904d1ff
	VPERMILPS $-1, Y1, Y2        // c4e37d04d1ff
	VCVTPS2PH $-1, X1, X2        // c4e3791dcaff
	VCVTPS2PH $-1, Y1, X2        // c4e37d1dcaff
	VPSLLW $-1, X1, X2           // c5e971f1ff
	VPSLLW $-1, Y1, Y2           // c5ed71f1ff
	VPSRAD $-1, X1, X2           // c5e972e1ff
	VPSRAD $-1, Y1, Y2           // c5ed72e1ff
	VPSRAW $-1, X1, X2           // c5e971e1ff
	VPSRAW $-1, Y1, Y2           // c5ed71e1ff
	VPSRLW $-1, X1, X1           // c5f171d1ff
	VPSRLW $-1, Y1, Y2           // c5ed71d1ff
	VEXTRACTPS $-1, X1, AX       // c4e37917c8ff
	VPEXTRB $-1, X1, AX          // c4e37914c8ff
	VPEXTRD $-1, X1, AX          // c4e37916c8ff
	VPEXTRQ $-1, X1, AX          // c4e3f916c8ff
	// End of tests.
	RET
