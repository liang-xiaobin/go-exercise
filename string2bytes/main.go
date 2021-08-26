package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello world"
	v := string2bytes1(s)
	fmt.Println(v)
}

// string2bytes Data作为值拷贝的情况，这就会导致无法保证它所引用的数据不会被垃圾回收(GC)
func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

//string2bytes1 在程序中必须保留一个单独的、正确类型的指向底层数据的指针
func string2bytes1(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var b []byte
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pbytes.Data = stringHeader.Data
	pbytes.Len = stringHeader.Len
	pbytes.Cap = stringHeader.Len
	return b
}

// 如果只是期望单纯的转换，对容
func string2bytes2(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
