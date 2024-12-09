package mlib

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func RandString(n int) string {
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

func readLine() {
	r := bufio.NewReaderSize(os.Stdin, 1024*1024)
	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(line), isPrefix)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func arrayEqual(a, b []int) bool {
	n := len(a)
	if len(b) != n {
		return false
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Duplicate[T any](v T, n int) []T {
	a := make([]T, n)
	for i := 0; i < n; i++ {
		a[i] = v
	}
	return a
}

func twoDimensionArrayShape[T int | int32 | int64 | byte | uint | bool](m, n int, defaultValue T) [][]T {
	arr := make([][]T, m)
	for i := range arr {
		arr[i] = make([]T, n)
		for j := range arr[i] {
			arr[i][j] = defaultValue
		}
	}
	return arr
}

func RecoverPrint() {
	v := recover()
	fmt.Println("recover:", v)
}
