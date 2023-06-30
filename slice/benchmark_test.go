package slice

import "testing"

func BenchmarkSum15(b *testing.B) {
	c := New[int](15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(c)
	}
}

func BenchmarkShift(b *testing.B) {
	c := New[int](15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shift(c)
	}
}

func BenchmarkShift1000(b *testing.B) {
	c := New[int](15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			Shift(c)
		}
	}
}

func BenchmarkShift1000000(b *testing.B) {
	c := New[int](15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			Shift(c)
		}
	}
}

func BenchmarkShiftN1000(b *testing.B) {
	c := New[int](15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftN(c, 1000)
	}
}

func BenchmarkShiftN1000000(b *testing.B) {
	c := New[int](15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftN(c, 1000000)
	}
}

func BenchmarkShiftNOverflow(b *testing.B) {
	c := New[int](10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftN(c, 1024)
	}
}

func BenchmarkShiftNFOverflow(b *testing.B) {
	c := New[int](10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftNF(c, 1024)
	}
}

func BenchmarkCeilIndexOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 5; j++ {
			CeilIndexOf(j)
		}
	}
}
