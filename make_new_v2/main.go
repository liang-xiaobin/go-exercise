package main

import "fmt"

type T struct {
	Name string
}

func main() {
	v := new(T)
	v.Name = "liangxiaobin"
	fmt.Println(v)
	t := &T{}
	t.Name = "kkk"
	fmt.Println(t)
	_ = new(chan bool)
	//v2 := new(map[string]struct{})
	v2 := make(map[string]struct{}, 1)
	fmt.Println(v2)
}