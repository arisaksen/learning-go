package main

import "testing"

func BenchmarkGenerateTestData(b *testing.B) {
	test := []int{1, 2, 3, 4, 5}

	b.Run("Remove element from list", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			removeFromSliceReturnUnordered(test, 1)
		}
		b.ReportAllocs()
	})

	b.Run("Remove element from list ordered", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			removeFromSliceReturnOrdered(test, 1)
		}
		b.ReportAllocs()
	})

}
