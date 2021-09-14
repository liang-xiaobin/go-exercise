package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data *byte
	var in interface{}
	in = data
	fmt.Println(IsNil(in))
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}