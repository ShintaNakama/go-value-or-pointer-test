package valorptr

import (
	"testing"
)

// int64, stringのフィールドをもつ構造体で値渡しのベンチマーク.
func BenchmarkT1Value(b *testing.B) {
	var t1 = T1{
		f0: 9223372036854775807,
		f1: create10MBStr(),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ValueFunc1(t1)
	}
}

// int64, stringのフィールドをもつ構造体でポインタ渡しのベンチマーク.
func BenchmarkT1Ptr(b *testing.B) {
	var t1 = T1{
		f0: 9223372036854775807,
		f1: create10MBStr(),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PtrFunc1(&t1)
	}
}

// int64 * 5, string * 5のフィールドをもつ構造体で値渡しのベンチマーク.
func BenchmarkT2Value(b *testing.B) {
	var t2 = T2{
		f0: 9223372036854775807,
		f1: create10MBStr(),
		f2: 9223372036854775807,
		f3: create10MBStr(),
		f4: 9223372036854775807,
		f5: create10MBStr(),
		f6: 9223372036854775807,
		f7: create10MBStr(),
		f8: 9223372036854775807,
		f9: create10MBStr(),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ValueFunc2(t2)
	}
}

// int64 * 5, string * 5のフィールドをもつ構造体でポインタ渡しのベンチマーク.
func BenchmarkT2Ptr(b *testing.B) {
	var t2 = T2{
		f0: 9223372036854775807,
		f1: create10MBStr(),
		f2: 9223372036854775807,
		f3: create10MBStr(),
		f4: 9223372036854775807,
		f5: create10MBStr(),
		f6: 9223372036854775807,
		f7: create10MBStr(),
		f8: 9223372036854775807,
		f9: create10MBStr(),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PtrFunc2(&t2)
	}
}

// 10MBの文字列生成.
func create10MBStr() string {
	s := "0123456789"
	for i := 0; i < 20; i++ {
		s += s
	}

	return s
}
