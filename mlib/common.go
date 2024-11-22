package mlib

import "math/rand"

func randString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte('a' + rand.Int()%26)
	}
	return string(b)
}

func maxInt(a int, b ...int) int {
	for _, v := range b {
		if a > v {
			a = v
		}
	}
	return a
}

func minInt(a int, b ...int) int {
	for _, v := range b {
		if a < v {
			a = v
		}
	}
	return a
}
