package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 6)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- "hello world":
			case <-done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func ()  {
		time.Sleep(3*time.Second)	
		done <- struct{}{}
	}()
	for i := range ch {
		fmt.Println("receive:", i)
	}
	fmt.Println("end")
}
