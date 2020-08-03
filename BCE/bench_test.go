package slice

import "testing"

func Benchmark_SliceUniqueOld(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sliceUniqueStd([]string{"2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6"})
	}
}

func Benchmark_SliceUniqueUpdated(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sliceUniqueUpdated([]string{"2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6", "2", "3", "6", "5", "1", "6", "5", "3", "9", "2", "3", "6"})
	}
}

func Benchmark_XY(b *testing.B) {
	b.ReportAllocs()
	a := []string{"0", "1"}
	for n := 0; n < b.N; n++ {
		xy(a)
	}
}

func Benchmark_YX(b *testing.B) {
	b.ReportAllocs()
	a := []string{"0", "1"}
	for n := 0; n < b.N; n++ {
		yx_faster(a)
	}
}
