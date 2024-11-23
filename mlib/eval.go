package mlib

func eval(s string) int {
	stack := []int{}
	num := 0
	n := len(s)
	op := byte('+')
	for i := 0; i < n; i++ {
		b := s[i]
		if b == ' ' {
			continue
		}
		if b >= '0' && b <= '9' {
			num = num*10 + int(b-'0')
		} else if b == '(' || b == '[' || b == '{' {
			j := i + 1
			count := 1
			for count > 0 {
				if s[j] == ')' || s[j] == ']' || s[j] == '}' {
					count--
				} else if s[j] == '(' || s[j] == ']' || s[j] == '}' {
					count++
				}
				j++
			}
			num = eval(s[i+1 : j-1])
			i = j - 1
		}
		if !(b >= '0' && b <= '9') || i == n-1 {
			if op == '+' {
				stack = append(stack, num)
			} else if op == '-' {
				stack = append(stack, -num)
			} else if op == '*' {
				l := len(stack)
				stack[l-1] *= num
			} else if op == '/' {
				l := len(stack)
				stack[l-1] /= num
			}
			op = b
			num = 0
		}
	}
	ans := 0
	for _, v := range stack {
		ans += v
	}
	return ans
}
