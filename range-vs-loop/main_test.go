package main

import "testing"

var str = getString(100000)

func BenchmarkRangeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeString(str)
	}
}

func BenchmarkRangeIndexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeIndexString(str)
	}
}

func BenchmarkLoopString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopString(str)
	}
}

var slice = getSlice(100000)

func BenchmarkRangeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeSlice(slice)
	}
}

func BenchmarkLoopSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopSlice(slice)
	}
}
