package main

import "fmt"

func main() {
	var v interface{}
	var data *byte
	var in interface{}
	v = (*int)(nil)
	fmt.Println(v == nil)
	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)
	in = data
	fmt.Println(in, in == nil)
}
