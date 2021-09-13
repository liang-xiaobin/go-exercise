package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		ch <- "hello world"
	}()
	select {
	case _ = <-ch:
	case <-time.After(time.Second * 1):
		fmt.Println("hello world")
	}
}
