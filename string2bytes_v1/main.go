package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func StringToByte(key *string) []byte {
	strPtr := (*reflect.SliceHeader)(unsafe.Pointer(key))
	strPtr.Cap = strPtr.Len
	b := *(*[]byte)(unsafe.Pointer(strPtr))
	return b
}

// main 测试go中踩内存的情况
func main() {
	decryptContent := "/AvYEjm4g6xJ3LVrk2/Adk"
	iv := decryptContent[0:16]
	key := decryptContent[2:18]
	fmt.Println(&iv)
	fmt.Println(&key)
	ivBytes := StringToByte(&iv)
	keyBytes := StringToByte(&key)
	fmt.Println(string(ivBytes))
	fmt.Println(string(keyBytes))
}
