package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 初始化底层数组
	s := [4]string{"1", "2", "3", "4"}
	s1 := s[0:1]
	s2 := s[:]
	// 构造sliceHeader
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	sh2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	fmt.Println(sh1.Len, sh1.Cap, sh1.Data)
	fmt.Println(sh2.Len, sh2.Cap, sh2.Data)
}
