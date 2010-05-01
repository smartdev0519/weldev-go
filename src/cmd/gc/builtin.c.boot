char *runtimeimport =
	"package runtime\n"
	"func \"\".mal (? int32) *any\n"
	"func \"\".panicindex ()\n"
	"func \"\".panicslice ()\n"
	"func \"\".throwreturn ()\n"
	"func \"\".throwinit ()\n"
	"func \"\".panic (? interface { })\n"
	"func \"\".recover (? *int32) interface { }\n"
	"func \"\".printbool (? bool)\n"
	"func \"\".printfloat (? float64)\n"
	"func \"\".printint (? int64)\n"
	"func \"\".printuint (? uint64)\n"
	"func \"\".printcomplex (? complex128)\n"
	"func \"\".printstring (? string)\n"
	"func \"\".printpointer (? any)\n"
	"func \"\".printiface (? any)\n"
	"func \"\".printeface (? any)\n"
	"func \"\".printslice (? any)\n"
	"func \"\".printnl ()\n"
	"func \"\".printsp ()\n"
	"func \"\".printf ()\n"
	"func \"\".catstring (? string, ? string) string\n"
	"func \"\".cmpstring (? string, ? string) int\n"
	"func \"\".slicestring (? string, ? int, ? int) string\n"
	"func \"\".slicestring1 (? string, ? int) string\n"
	"func \"\".indexstring (? string, ? int) uint8\n"
	"func \"\".intstring (? int64) string\n"
	"func \"\".slicebytetostring (? []uint8) string\n"
	"func \"\".sliceinttostring (? []int) string\n"
	"func \"\".stringtoslicebyte (? string) []uint8\n"
	"func \"\".stringtosliceint (? string) []int\n"
	"func \"\".stringiter (? string, ? int) int\n"
	"func \"\".stringiter2 (? string, ? int) (retk int, retv int)\n"
	"func \"\".slicecopy (to any, fr any, wid uint32) int\n"
	"func \"\".ifaceI2E (iface any) any\n"
	"func \"\".ifaceE2I (typ *uint8, iface any) any\n"
	"func \"\".ifaceT2E (typ *uint8, elem any) any\n"
	"func \"\".ifaceE2T (typ *uint8, elem any) any\n"
	"func \"\".ifaceE2I2 (typ *uint8, iface any) (ret any, ok bool)\n"
	"func \"\".ifaceE2T2 (typ *uint8, elem any) (ret any, ok bool)\n"
	"func \"\".ifaceT2I (typ1 *uint8, typ2 *uint8, elem any) any\n"
	"func \"\".ifaceI2T (typ *uint8, iface any) any\n"
	"func \"\".ifaceI2T2 (typ *uint8, iface any) (ret any, ok bool)\n"
	"func \"\".ifaceI2I (typ *uint8, iface any) any\n"
	"func \"\".ifaceI2Ix (typ *uint8, iface any) any\n"
	"func \"\".ifaceI2I2 (typ *uint8, iface any) (ret any, ok bool)\n"
	"func \"\".ifaceeq (i1 any, i2 any) bool\n"
	"func \"\".efaceeq (i1 any, i2 any) bool\n"
	"func \"\".ifacethash (i1 any) uint32\n"
	"func \"\".efacethash (i1 any) uint32\n"
	"func \"\".makemap (key *uint8, val *uint8, hint int64) map[any] any\n"
	"func \"\".mapaccess1 (hmap map[any] any, key any) any\n"
	"func \"\".mapaccess2 (hmap map[any] any, key any) (val any, pres bool)\n"
	"func \"\".mapassign1 (hmap map[any] any, key any, val any)\n"
	"func \"\".mapassign2 (hmap map[any] any, key any, val any, pres bool)\n"
	"func \"\".mapiterinit (hmap map[any] any, hiter *any)\n"
	"func \"\".mapiternext (hiter *any)\n"
	"func \"\".mapiter1 (hiter *any) any\n"
	"func \"\".mapiter2 (hiter *any) (key any, val any)\n"
	"func \"\".makechan (elem *uint8, hint int64) chan any\n"
	"func \"\".chanrecv1 (hchan <-chan any) any\n"
	"func \"\".chanrecv2 (hchan <-chan any) (elem any, pres bool)\n"
	"func \"\".chansend1 (hchan chan<- any, elem any)\n"
	"func \"\".chansend2 (hchan chan<- any, elem any) bool\n"
	"func \"\".closechan (hchan any)\n"
	"func \"\".closedchan (hchan any) bool\n"
	"func \"\".newselect (size int) *uint8\n"
	"func \"\".selectsend (sel *uint8, hchan chan<- any, elem any) bool\n"
	"func \"\".selectrecv (sel *uint8, hchan <-chan any, elem *any) bool\n"
	"func \"\".selectdefault (sel *uint8) bool\n"
	"func \"\".selectgo (sel *uint8)\n"
	"func \"\".makeslice (typ *uint8, nel int64, cap int64) []any\n"
	"func \"\".sliceslice1 (old []any, lb int, width int) []any\n"
	"func \"\".sliceslice (old []any, lb int, hb int, width int) []any\n"
	"func \"\".slicearray (old *any, nel int, lb int, hb int, width int) []any\n"
	"func \"\".closure ()\n"
	"func \"\".int64div (? int64, ? int64) int64\n"
	"func \"\".uint64div (? uint64, ? uint64) uint64\n"
	"func \"\".int64mod (? int64, ? int64) int64\n"
	"func \"\".uint64mod (? uint64, ? uint64) uint64\n"
	"func \"\".float64toint64 (? float64) int64\n"
	"func \"\".int64tofloat64 (? int64) float64\n"
	"func \"\".complex128div (num complex128, den complex128) complex128\n"
	"\n"
	"$$\n";
char *unsafeimport =
	"package unsafe\n"
	"type \"\".Pointer *any\n"
	"func \"\".Offsetof (? any) int\n"
	"func \"\".Sizeof (? any) int\n"
	"func \"\".Alignof (? any) int\n"
	"func \"\".Typeof (i interface { }) interface { }\n"
	"func \"\".Reflect (i interface { }) (typ interface { }, addr \"\".Pointer)\n"
	"func \"\".Unreflect (typ interface { }, addr \"\".Pointer) interface { }\n"
	"func \"\".New (typ interface { }) \"\".Pointer\n"
	"func \"\".NewArray (typ interface { }, n int) \"\".Pointer\n"
	"\n"
	"$$\n";
