package main

import (
	"fmt"
	"sync"
)

type Man struct {
	Animal
	Name string
	Age  int
}

type Animal struct {
	Age int
}

func (a *Animal) GetAge() int {
	return a.Age
}

func main() {
	p := new(sync.Pool)
	p.Put(1)
	fmt.Println(p.Get())

}
