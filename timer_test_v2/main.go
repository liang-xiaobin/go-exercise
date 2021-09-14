package main

import (
	"fmt"
	"time"
)

func main() {
	v := make(chan struct{})
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("hello world")
		v <- struct{}{}
	})
	defer timer.Stop()
	fmt.Println("hello world1")
	<-v
}
