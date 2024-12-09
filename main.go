package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	p := Person{"xx", 11}
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.ValueOf(p))
	fmt.Println(reflect.TypeOf(p).Field(1).Tag.Get("json"))
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
