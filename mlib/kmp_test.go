package mlib

import (
	"math/rand"
	"strings"
	"testing"
)

func Test_kmp(t *testing.T) {
	for i := 0; i < 10000; i++ {
		s := randString(rand.Int()%40 + 1)
		p := randString(rand.Int()%10 + 1)
		if kmp(s, p) != strings.Index(s, p) {
			t.Fatal(s, p, kmp(s, p))
		}
	}
}
