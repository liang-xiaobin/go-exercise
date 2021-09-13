package main

import (
	"fmt"
	"runtime"
	"time"
)

// main goroutine
func main() {
	// 模拟单核 CPU
	runtime.GOMAXPROCS(1)
	// 模拟goroutine 死循环
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("hello world")
}
