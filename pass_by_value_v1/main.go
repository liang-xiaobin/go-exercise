package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["1"] = "a"
	fmt.Printf("main 内存地址:%p\n", &m)
	hello(m)
	fmt.Printf("%v", m)
}

func hello(p map[string]string) {
	fmt.Printf("hello 内存地址:%p\n", &p)
	p["1"] = "b"
}
