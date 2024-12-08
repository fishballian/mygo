package main

var Prime []int
var memo []int

func init() {
	isPrime := make([]bool, 1e8+1)
	for i := 2; i <= 2e8; i++ {
		if !isPrime[i] {
			Prime = append(Prime, i)
			if palindrome(i) {
				memo = append(memo, i)
			}
		}
		for _, p := range Prime {
			if i*p > 2e8 {
				break
			}
			isPrime[i*p] = true
			if i%p == 0 {
				break
			}
		}
	}
}

func palindrome(num int) bool {
	val := num
	tmp := 0
	for val != 0 {
		tmp, val = tmp*10+val%10, val/10
	}
	return tmp == num
}
