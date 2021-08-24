package main

import (
	"fmt"
	"unsafe"
)

type Num struct {
	i string
	j int64
}

func main() {
	n := Num{i: "Liang", j: 1}
	nPointer := unsafe.Pointer(&n)
	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "hello world"
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2
	fmt.Printf("n.i:%s, n.j:%d", n.i, n.j)
}
