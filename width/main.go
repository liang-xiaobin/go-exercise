package main

import (
	"fmt"
	"time"
	"unsafe"
)
type S struct {
	A struct{}
	B struct{}
}

type Set map[string]struct{}

func (s Set) Append(k string) {
	s[k] = struct{}{}
}

func (s Set) Remove(k string) {
	delete(s, k)
}

func (s Set) Exist(k string) bool {
	_, ok := s[k]
	return ok
}

type T struct {}

func (t *T) Call() {
	fmt.Println("hello world")
}

func main() {
	var a int 
	var b string
	var c bool
	var d [3]int32
	var e []string
	var f map[string]bool
	fmt.Println(
		unsafe.Sizeof(a),
		unsafe.Sizeof(b),
		unsafe.Sizeof(c),
		unsafe.Sizeof(d),
		unsafe.Sizeof(e),
		unsafe.Sizeof(f),
	)
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))
	var k S
	fmt.Println(unsafe.Sizeof(k))
	var v T
	v.Call()
	set := Set{}
	set.Append("hello")
	set.Append("world")
	set.Append("hello")
	set.Append("world")
	set.Remove("hello")
	fmt.Println(set.Exist("hello"))
	ch := make(chan struct{})
	go func() {
		time.Sleep(1*time.Second)
		close(ch)
	}()
	fmt.Println("sleep ")
	<-ch
	fmt.Println("kkkkkk")
}