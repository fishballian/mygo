package main

import (
	"mygo/mlib"
	"testing"
)

var s = mlib.RandString(1e6)

func BenchmarkCountSubstrings1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSubstrings1(s)
	}
}

func BenchmarkCountSubstrings2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSubstrings2(s)
	}
}
