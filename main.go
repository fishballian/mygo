package main

import (
	"fmt"
	"mygo/mlib"
)

func main() {
	for i := 0; i < 101; i++ {
		fmt.Println(i, mlib.CallM(i))
	}
}
