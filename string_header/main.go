package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello world"
	s1 := "hello world"
	s2 := "hello world"[3:]
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Data)
}
