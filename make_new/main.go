package main

import "fmt"

func main() {
	v1 := make([]int, 1, 5)
	v2 := make(map[int]bool, 5)
	v3 := make(chan int, 1)
	fmt.Println(v1, v2, v3)
}
