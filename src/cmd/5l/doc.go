// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

5l is a modified version of the Plan 9 linker.  The original is documented at

	http://plan9.bell-labs.com/magic/man2html/1/2l

Its target architecture is the ARM, referred to by these tools as arm.
It reads files in .5 format generated by 5g, 5c, and 5a and emits
a binary called 5.out by default.

Major changes include:
	- support for segmented stacks (this feature is implemented here, not in the compilers).


Original options are listed in the link above.

Options new in this version:

-L dir1 -L dir2
	Search for libraries (package files) in dir1, dir2, etc.
	The default is the single location $GOROOT/pkg/$GOOS_arm.
-r dir1:dir2:...
	Set the dynamic linker search path when using ELF.
-V
	Print the linker version.

*/
package documentation
