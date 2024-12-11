package mlib

func GenFac(n, mod int) ([]int, []int) {
	fac := make([]int, n+1)
	inv := make([]int, n+1)
	fac[0] = 1
	inv[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
		inv[i] = PowMod(fac[i], mod-2, mod)
	}
	return fac, inv
}

func GenComb(n, mod int) [][]int {
	comb := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		comb[i] = make([]int, n+1)
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % mod
		}
	}
	return comb
}

func PowMod(a, b, mod int) int {
	res := 1
	t := a
	for b > 0 {
		if b&1 == 1 {
			res = res * t % mod
		}
		t = t * t % mod
		b >>= 1
	}
	return res
}
