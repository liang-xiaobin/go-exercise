package main

import "log"

type MyErr struct {
	Msg string
}

// main 测试interface{}只有在类型和值都为nil的情况下, interface的nil判断为true
func main() {
	var e error
	e = GetErr()
	log.Println(e == nil)
}

func GetErr() *MyErr {
	return nil
}

func (m *MyErr) Error() string {
	return "hello world"
}
