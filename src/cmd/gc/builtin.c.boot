char *runtimeimport =
	"package runtime\n"
	"import runtime \"runtime\"\n"
	"func @\"\".new(typ *byte) *any\n"
	"func @\"\".panicindex()\n"
	"func @\"\".panicslice()\n"
	"func @\"\".throwreturn()\n"
	"func @\"\".throwinit()\n"
	"func @\"\".panicwrap(? string, ? string, ? string)\n"
	"func @\"\".panic(? interface {})\n"
	"func @\"\".recover(? *int32) interface {}\n"
	"func @\"\".printbool(? bool)\n"
	"func @\"\".printfloat(? float64)\n"
	"func @\"\".printint(? int64)\n"
	"func @\"\".printuint(? uint64)\n"
	"func @\"\".printcomplex(? complex128)\n"
	"func @\"\".printstring(? string)\n"
	"func @\"\".printpointer(? any)\n"
	"func @\"\".printiface(? any)\n"
	"func @\"\".printeface(? any)\n"
	"func @\"\".printslice(? any)\n"
	"func @\"\".printnl()\n"
	"func @\"\".printsp()\n"
	"func @\"\".goprintf()\n"
	"func @\"\".concatstring()\n"
	"func @\"\".append()\n"
	"func @\"\".appendslice(typ *byte, x any, y []any) any\n"
	"func @\"\".appendstr(typ *byte, x []byte, y string) []byte\n"
	"func @\"\".cmpstring(? string, ? string) int\n"
	"func @\"\".slicestring(? string, ? int, ? int) string\n"
	"func @\"\".slicestring1(? string, ? int) string\n"
	"func @\"\".intstring(? int64) string\n"
	"func @\"\".slicebytetostring(? []byte) string\n"
	"func @\"\".slicerunetostring(? []rune) string\n"
	"func @\"\".stringtoslicebyte(? string) []byte\n"
	"func @\"\".stringtoslicerune(? string) []rune\n"
	"func @\"\".stringiter(? string, ? int) int\n"
	"func @\"\".stringiter2(? string, ? int) (retk int, retv rune)\n"
	"func @\"\".copy(to any, fr any, wid uint32) int\n"
	"func @\"\".slicestringcopy(to any, fr any) int\n"
	"func @\"\".convI2E(elem any) any\n"
	"func @\"\".convI2I(typ *byte, elem any) any\n"
	"func @\"\".convT2E(typ *byte, elem any) any\n"
	"func @\"\".convT2I(typ *byte, typ2 *byte, elem any) any\n"
	"func @\"\".assertE2E(typ *byte, iface any) any\n"
	"func @\"\".assertE2E2(typ *byte, iface any) (ret any, ok bool)\n"
	"func @\"\".assertE2I(typ *byte, iface any) any\n"
	"func @\"\".assertE2I2(typ *byte, iface any) (ret any, ok bool)\n"
	"func @\"\".assertE2T(typ *byte, iface any) any\n"
	"func @\"\".assertE2T2(typ *byte, iface any) (ret any, ok bool)\n"
	"func @\"\".assertI2E(typ *byte, iface any) any\n"
	"func @\"\".assertI2E2(typ *byte, iface any) (ret any, ok bool)\n"
	"func @\"\".assertI2I(typ *byte, iface any) any\n"
	"func @\"\".assertI2I2(typ *byte, iface any) (ret any, ok bool)\n"
	"func @\"\".assertI2T(typ *byte, iface any) any\n"
	"func @\"\".assertI2T2(typ *byte, iface any) (ret any, ok bool)\n"
	"func @\"\".ifaceeq(i1 any, i2 any) bool\n"
	"func @\"\".efaceeq(i1 any, i2 any) bool\n"
	"func @\"\".ifacethash(i1 any) uint32\n"
	"func @\"\".efacethash(i1 any) uint32\n"
	"func @\"\".makemap(mapType *byte, hint int64) map[any]any\n"
	"func @\"\".mapaccess1(mapType *byte, hmap map[any]any, key any) any\n"
	"func @\"\".mapaccess2(mapType *byte, hmap map[any]any, key any) (val any, pres bool)\n"
	"func @\"\".mapassign1(mapType *byte, hmap map[any]any, key any, val any)\n"
	"func @\"\".mapassign2(mapType *byte, hmap map[any]any, key any, val any, pres bool)\n"
	"func @\"\".mapiterinit(mapType *byte, hmap map[any]any, hiter *any)\n"
	"func @\"\".mapdelete(mapType *byte, hmap map[any]any, key any)\n"
	"func @\"\".mapiternext(hiter *any)\n"
	"func @\"\".mapiter1(hiter *any) any\n"
	"func @\"\".mapiter2(hiter *any) (key any, val any)\n"
	"func @\"\".makechan(chanType *byte, hint int64) chan any\n"
	"func @\"\".chanrecv1(chanType *byte, hchan <-chan any) any\n"
	"func @\"\".chanrecv2(chanType *byte, hchan <-chan any) (elem any, received bool)\n"
	"func @\"\".chansend1(chanType *byte, hchan chan<- any, elem any)\n"
	"func @\"\".closechan(hchan any)\n"
	"func @\"\".selectnbsend(chanType *byte, hchan chan<- any, elem any) bool\n"
	"func @\"\".selectnbrecv(chanType *byte, elem *any, hchan <-chan any) bool\n"
	"func @\"\".selectnbrecv2(chanType *byte, elem *any, received *bool, hchan <-chan any) bool\n"
	"func @\"\".newselect(size int) *byte\n"
	"func @\"\".selectsend(sel *byte, hchan chan<- any, elem *any) bool\n"
	"func @\"\".selectrecv(sel *byte, hchan <-chan any, elem *any) bool\n"
	"func @\"\".selectrecv2(sel *byte, hchan <-chan any, elem *any, received *bool) bool\n"
	"func @\"\".selectdefault(sel *byte) bool\n"
	"func @\"\".selectgo(sel *byte)\n"
	"func @\"\".block()\n"
	"func @\"\".makeslice(typ *byte, nel int64, cap int64) []any\n"
	"func @\"\".growslice(typ *byte, old []any, n int64) []any\n"
	"func @\"\".sliceslice1(old []any, lb uint64, width uint64) []any\n"
	"func @\"\".sliceslice(old []any, lb uint64, hb uint64, width uint64) []any\n"
	"func @\"\".slicearray(old *any, nel uint64, lb uint64, hb uint64, width uint64) []any\n"
	"func @\"\".closure()\n"
	"func @\"\".int64div(? int64, ? int64) int64\n"
	"func @\"\".uint64div(? uint64, ? uint64) uint64\n"
	"func @\"\".int64mod(? int64, ? int64) int64\n"
	"func @\"\".uint64mod(? uint64, ? uint64) uint64\n"
	"func @\"\".float64toint64(? float64) int64\n"
	"func @\"\".float64touint64(? float64) uint64\n"
	"func @\"\".int64tofloat64(? int64) float64\n"
	"func @\"\".uint64tofloat64(? uint64) float64\n"
	"func @\"\".complex128div(num complex128, den complex128) complex128\n"
	"\n"
	"$$\n";
char *unsafeimport =
	"package unsafe\n"
	"import runtime \"runtime\"\n"
	"type @\"\".Pointer uintptr\n"
	"func @\"\".Offsetof(? any) uintptr\n"
	"func @\"\".Sizeof(? any) uintptr\n"
	"func @\"\".Alignof(? any) uintptr\n"
	"func @\"\".Typeof(i interface {}) interface {}\n"
	"func @\"\".Reflect(i interface {}) (typ interface {}, addr @\"\".Pointer)\n"
	"func @\"\".Unreflect(typ interface {}, addr @\"\".Pointer) interface {}\n"
	"func @\"\".New(typ interface {}) @\"\".Pointer\n"
	"func @\"\".NewArray(typ interface {}, n int) @\"\".Pointer\n"
	"\n"
	"$$\n";
