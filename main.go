package main

func main() {
	a := []int{1, 2, 4}
	for {
		a = append(a, 1)
		a = a[1:]
	}
}
