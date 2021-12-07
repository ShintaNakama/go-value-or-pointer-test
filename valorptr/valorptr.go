package valorptr

type T1 struct {
	f0 int64
	f1 string
}

type T2 struct {
	f0 int64
	f1 string
	f2 int64
	f3 string
	f4 int64
	f5 string
	f6 int64
	f7 string
	f8 int64
	f9 string
}

var num int64
var str string

//go:noinline
func ValueFunc1(t1 T1) {
	num = t1.f0
	str = t1.f1
}

//go:noinline
func PtrFunc1(t1 *T1) {
	num = t1.f0
	str = t1.f1
}

//go:noinline
func ValueFunc2(t2 T2) {
	num = t2.f0
	str = t2.f1
	num = t2.f2
	str = t2.f3
	num = t2.f4
	str = t2.f5
	num = t2.f6
	str = t2.f7
	num = t2.f8
	str = t2.f9
}

//go:noinline
func PtrFunc2(t2 *T2) {
	num = t2.f0
	str = t2.f1
	num = t2.f2
	str = t2.f3
	num = t2.f4
	str = t2.f5
	num = t2.f6
	str = t2.f7
	num = t2.f8
	str = t2.f9
}
