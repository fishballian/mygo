package mlib

func kmp(s, p string) int {
	n := len(p)
	fail := make([]int, n)
	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j > 0 && p[i] != p[j] {
			j = fail[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		fail[i] = j
	}
	match := 0
	for i := 0; i < len(s); i++ {
		for match > 0 && s[i] != p[match] {
			match = fail[match-1]
		}
		if s[i] == p[match] {
			match++
			if match == n {
				return i - n + 1
			}
		}
	}
	return -1
}

func zKmp(s, p string) []int {
	p += s
	n := len(p)
	l, r := 0, 0
	z := make([]int, n)
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && p[i+z[i]] == p[z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	return z
}
