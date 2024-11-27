package mlib

import (
	"math/rand"
	"testing"
)

func Test_manacher(t *testing.T) {
	for i := 0; i < 10000; i++ {
		n := rand.Int()%10 + 1
		s := randString(n)
		d1, d2 := manacher(s)
		//t.Log(s, d1, d2)
		c1 := make([]int, len(s))
		c2 := make([]int, len(s))
		for i := 0; i < n; i++ {
			k := 1
			for i-k >= 0 && i+k < n && s[i-k] == s[i+k] {
				k++
			}
			c1[i] = k
			k2 := 0
			for i-k2-1 >= 0 && i+k2 < n && s[i-k2-1] == s[i+k2] {
				k2++
			}
			c2[i] = k2
		}
		if !arrayEqual(d1, c1) {
			t.Fatal("manacher not equal", s, d1, c1)
		}
		if !arrayEqual(d2, c2) {
			t.Fatal("manacher not equal", s, d2, c2)
		}
	}
}
