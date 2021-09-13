package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		fmt.Println("groutines: ", runtime.NumGoroutine())
	}()
	var ch chan int
	go func() {
		<- ch
	}()
	time.Sleep(time.Second)
}