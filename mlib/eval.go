package mlib

import "strings"

func eval(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	n := len(s)
	var stack []int
	num := 0
	op := byte('+')
	for i := 0; i < n; i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == '(' || c == '[' || c == '{' {
			j, count := i+1, 1
			for count > 0 {
				if s[j] == ')' || s[j] == ']' || s[j] == '}' {
					count--
				} else if s[j] == '(' || s[j] == ']' || s[j] == '}' {
					count++
				}
				j++
			}
			num = eval(s[i+1 : j])
			i = j - 1
		}
		if !(c >= '0' && c <= '9') || i == n-1 {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				l := len(stack)
				stack[l-1] *= num
			case '/':
				l := len(stack)
				stack[l-1] /= num
			}
			op = c
			num = 0
		}
	}
	ans := 0
	for _, v := range stack {
		ans += v
	}
	return ans
}
