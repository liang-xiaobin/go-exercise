package main

import "fmt"

func main() {
	const N = 1024
	var a [N]int
	m := append(a[:N-1:N], 9, 9)
	k := append(a[:N:N], 9)
	x := cap(m)
	y := cap(k)
	fmt.Println(x, y)
}
