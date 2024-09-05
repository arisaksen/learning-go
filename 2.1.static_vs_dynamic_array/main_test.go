package main

import (
	"testing"
)

func BenchmarkCompare(b *testing.B) {
	b.Run("DynamicArray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DynamicArray()
		}
		b.ReportAllocs()
	})

	b.Run("StaticArray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StaticArray()
		}
		b.ReportAllocs()
	})

}

func BenchmarkCompare2(b *testing.B) {

	b.Run("StaticArray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StaticArray()
		}
		b.ReportAllocs()
	})

	b.Run("DynamicArray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DynamicArray()
		}
		b.ReportAllocs()
	})

}
