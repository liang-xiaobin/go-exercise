package main

import "fmt"

//MyType 有限的枚举类型测试
type MyType int

//使用常量模拟不完全枚举
const (
	A MyType = iota
	B
	C
	D
)

func (f MyType) String() string {
	return [...]string{"A", "B", "C", "D"}[f]
}

func main() {
	var f MyType = A
	fmt.Println(f)
	switch f {
	case A:
		fmt.Println("this is A")
	case B:
		fmt.Println("this is B")
	default:
		fmt.Println("this is other")
	}
}
