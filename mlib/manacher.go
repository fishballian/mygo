package mlib

func manacher(s string) ([]int, []int) {
	return calcD1(s), calcD2(s)
}

func calcD1(s string) []int {
	n := len(s)
	d1 := make([]int, n)
	l, r := 0, -1
	for i := 0; i < n; i++ {
		k := 1
		if r >= i {
			k = min(d1[l+r-i], r-i+1)
		}
		for i-k >= 0 && i+k < n && s[i-k] == s[i+k] {
			k++
		}
		d1[i] = k
		k--
		if i+k > r {
			l, r = i-k, i+k
		}
	}
	return d1
}

func calcD2(s string) []int {
	n := len(s)
	d2 := make([]int, n)
	l, r := 0, -1
	for i := 0; i < n; i++ {
		k := 0
		if r >= i {
			k = min(d2[l+r-i+1], r-i+1)
		}
		for i-k-1 >= 0 && i+k < n && s[i-k-1] == s[i+k] {
			k++
		}
		d2[i] = k
		k--
		if i+k > r {
			l = i - k - 1
			r = i + k
		}
	}
	return d2
}
