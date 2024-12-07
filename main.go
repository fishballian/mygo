package main

func main() {
}

func countSubstrings1(s string) int {
	n := len(s)
	a := make([]byte, 2*n+1)
	for i := 0; i < n; i++ {
		a[2*i] = '#'
		a[2*i+1] = s[i]
	}
	a[2*n] = '#'
	d := make([]int, 2*n+1)
	l, r := 0, -1
	ans := 0
	for i := 0; i < 2*n+1; i++ {
		k := 1
		if i <= r {
			k = min(d[l+r-i], r-i+1)
		}
		for i-k >= 0 && i+k < 2*n+1 && a[i-k] == a[i+k] {
			k++
		}
		d[i] = k
		ans += k / 2
		k--
		if i+k > r {
			l, r = i-k, i+k
		}

	}
	return ans
}

func countSubstrings2(s string) int {
	n := len(s)
	t := "$#"
	for i := 0; i < n; i++ {
		t += string(s[i]) + "#"
	}
	n = len(t)
	t += "!"

	f := make([]int, n)
	iMax, rMax, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		// 初始化 f[i]
		if i <= rMax {
			f[i] = min(rMax-i+1, f[2*iMax-i])
		} else {
			f[i] = 1
		}
		// 中心拓展
		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		// 动态维护 iMax 和 rMax
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		// 统计答案, 当前贡献为 (f[i] - 1) / 2 上取整
		ans += f[i] / 2
	}
	return ans
}
