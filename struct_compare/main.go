package main

import "fmt"

type Value struct {
	Name   string
	Gender string
}

func main() {
	v1 := Value{Name: "test", Gender: "man"}
	v2 := Value{Name: "test", Gender: "man"}
	if v1 == v2 {
		fmt.Println("equal")
		return
	}
	fmt.Println("not equal")
}
