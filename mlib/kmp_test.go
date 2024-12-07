package mlib

import (
	"math/rand"
	"strings"
	"testing"
)

func Test_kmp(t *testing.T) {
	for i := 0; i < 10000; i++ {
		s := RandString(rand.Int()%40 + 1)
		p := RandString(rand.Int()%10 + 1)
		if kmp(s, p) != strings.Index(s, p) {
			t.Fatal(s, p, kmp(s, p))
		}
		z := zKmp(s, p)
		s2 := p + s
		for i := 1; i < len(z); i++ {
			if s2[:z[i]] != s2[i:i+z[i]] {
				t.Fatal(s, p, z, i)
			}
			if i+z[i] < len(s) && s2[z[i]] == s2[i+z[i]] {
				t.Fatal(s, p, z, i)
			}
		}
	}
}
