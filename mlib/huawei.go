package mlib

import "strings"

func findBest(a []int, pMax int) int {
	n := len(a)
	ans := 0
	var dfs func(i, s int)
	dfs = func(i, s int) {
		if i == n {
			return
		}
		dfs(i+1, s)
		s += a[i]
		if s > pMax {
			return
		}
		if pMax-s < pMax-ans {
			ans = s
		}
		dfs(i+1, s)
	}
	dfs(0, 0)
	return ans
}

func findBest2(a []int, pMax int) int {
	f := make([]bool, pMax+1)
	f[0] = true
	for _, v := range a {
		for i := pMax; i >= v; i-- {
			if f[i-v] {
				f[i] = true
			}
		}
	}
	for i := pMax; i >= 0; i-- {
		if f[i] {
			return i
		}
	}
	return 0
}

func format(k int, s string) string {
	a := []byte{}
	parts := strings.Split(s, "-")
	a = append(a, parts[0]...)
	a = append(a, '-')
	s2 := strings.Join(parts[1:], "")
	for i := 0; i < len(s2); i += k {
		t := caseTrans(s2[i:min(i+k, len(s2))])
		a = append(a, t...)
		if i+k < len(s2) {
			a = append(a, '-')
		}
	}
	return string(a)
}

func caseTrans(s string) string {
	n := len(s)
	upper := 0
	lower := 0
	for i := 0; i < n; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			upper++
		} else if s[i] >= 'a' && s[i] <= 'z' {
			lower++
		}
	}
	a := []byte(s)
	if upper > lower {
		for i := 0; i < n; i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				a[i] = a[i] + 'A' - 'a'
			}
		}
	} else if upper < lower {
		for i := 0; i < n; i++ {
			if s[i] >= 'A' && s[i] <= 'Z' {
				a[i] = a[i] + 'a' - 'A'
			}
		}
	}
	return string(a)
}

func findMaxMine(grid []string) int {
	m, n := len(grid), len(grid[0])
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || y < 0 || x == m || y == n || vis[x][y] || grid[x][y] == '0' {
			return 0
		}
		vis[x][y] = true
		res := 0
		res += int(grid[x][y] - '0')
		for _, d := range dir {
			res += dfs(x+d[0], y+d[1])
		}
		return res
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' || vis[i][j] {
				continue
			}
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
