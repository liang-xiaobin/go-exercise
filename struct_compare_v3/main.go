package main

import "fmt"

type Value struct {
	Name   string
	GoodAt []string
}

func main() {
	v1 := Value{Name: "test", GoodAt: []string{"1", "2", "3"}}
	v2 := Value{Name: "test", GoodAt: []string{"1", "2", "3"}}
	if v1 == v2 {
		fmt.Println("equal")
		return
	}
	fmt.Println("not equal")
}
