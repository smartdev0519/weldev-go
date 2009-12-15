// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package md4

import (
	"fmt"
	"io"
	"testing"
)

type md4Test struct {
	out string
	in  string
}

var golden = []md4Test{
	md4Test{"31d6cfe0d16ae931b73c59d7e0c089c0", ""},
	md4Test{"bde52cb31de33e46245e05fbdbd6fb24", "a"},
	md4Test{"ec388dd78999dfc7cf4632465693b6bf", "ab"},
	md4Test{"a448017aaf21d8525fc10ae87aa6729d", "abc"},
	md4Test{"41decd8f579255c5200f86a4bb3ba740", "abcd"},
	md4Test{"9803f4a34e8eb14f96adba49064a0c41", "abcde"},
	md4Test{"804e7f1c2586e50b49ac65db5b645131", "abcdef"},
	md4Test{"752f4adfe53d1da0241b5bc216d098fc", "abcdefg"},
	md4Test{"ad9daf8d49d81988590a6f0e745d15dd", "abcdefgh"},
	md4Test{"1e4e28b05464316b56402b3815ed2dfd", "abcdefghi"},
	md4Test{"dc959c6f5d6f9e04e4380777cc964b3d", "abcdefghij"},
	md4Test{"1b5701e265778898ef7de5623bbe7cc0", "Discard medicine more than two years old."},
	md4Test{"d7f087e090fe7ad4a01cb59dacc9a572", "He who has a shady past knows that nice guys finish last."},
	md4Test{"a6f8fd6df617c72837592fc3570595c9", "I wouldn't marry him with a ten foot pole."},
	md4Test{"c92a84a9526da8abc240c05d6b1a1ce0", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	md4Test{"f6013160c4dcb00847069fee3bb09803", "The days of the digital watch are numbered.  -Tom Stoppard"},
	md4Test{"2c3bb64f50b9107ed57640fe94bec09f", "Nepal premier won't resign."},
	md4Test{"45b7d8a32c7806f2f7f897332774d6e4", "For every action there is an equal and opposite government program."},
	md4Test{"b5b4f9026b175c62d7654bdc3a1cd438", "His money is twice tainted: 'taint yours and 'taint mine."},
	md4Test{"caf44e80f2c20ce19b5ba1cab766e7bd", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	md4Test{"191fae6707f496aa54a6bce9f2ecf74d", "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	md4Test{"9ddc753e7a4ccee6081cd1b45b23a834", "size:  a.out:  bad magic"},
	md4Test{"8d050f55b1cadb9323474564be08a521", "The major problem is with sendmail.  -Mark Horton"},
	md4Test{"ad6e2587f74c3e3cc19146f6127fa2e3", "Give me a rock, paper and scissors and I will move the world.  CCFestoon"},
	md4Test{"1d616d60a5fabe85589c3f1566ca7fca", "If the enemy is within range, then so are you."},
	md4Test{"aec3326a4f496a2ced65a1963f84577f", "It's well we cannot hear the screams/That we create in others' dreams."},
	md4Test{"77b4fd762d6b9245e61c50bf6ebf118b", "You remind me of a TV show, but that's all right: I watch it anyway."},
	md4Test{"e8f48c726bae5e516f6ddb1a4fe62438", "C is as portable as Stonehedge!!"},
	md4Test{"a3a84366e7219e887423b01f9be7166e", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	md4Test{"a6b7aa35157e984ef5d9b7f32e5fbb52", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	md4Test{"75661f0545955f8f9abeeb17845f3fd6", "How can you write a big system without C++?  -Paul Glick"},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		c := New()
		for j := 0; j < 2; j++ {
			io.WriteString(c, g.in)
			s := fmt.Sprintf("%x", c.Sum())
			if s != g.out {
				t.Errorf("md4[%d](%s) = %s want %s", j, g.in, s, g.out)
				t.FailNow()
			}
			c.Reset()
		}
	}
}
