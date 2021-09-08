package main

import "fmt"

type People struct{}

func main() {
	a := &People{}
	b := &People{}
	fmt.Println(a == b)
}
