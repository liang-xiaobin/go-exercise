package main

import "fmt"

type Value1 struct {
	Name string
}

type Value2 struct {
	Name string
}

func main() {
	v1 := Value1{Name: "k"}
	v2 := Value2{Name: "k"}
	if v1 == Value1(v2) {
		fmt.Println("equal")
		return
	}
	fmt.Println("not equal")
}
