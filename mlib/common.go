package mlib

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func randString(n int) string {
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
