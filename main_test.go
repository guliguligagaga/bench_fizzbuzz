package main

import (
	"testing"
)

func BenchmarkM(b *testing.B) {
	b.N = 2 ^ 32 - 1
	b.Run("f1", func(b *testing.B) {
		b.ReportAllocs()
		f1(b.N)
	})
	b.Run("f2", func(b *testing.B) {
		b.ReportAllocs()
		f2(b.N)
	})
	b.Run("f3", func(b *testing.B) {
		b.ReportAllocs()
		f3(b.N)
	})
	b.Run("f4", func(b *testing.B) {
		b.ReportAllocs()
		f4(b.N)
	})
}
