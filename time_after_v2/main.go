package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		in := 1
		for {
			in++
			ch <- in
		}
	}()

	for {
		select {
		case _ = <-ch:
			continue
		case <-time.After(3 * time.Minute):
			fmt.Printf("现在是: %d", time.Now().Unix())
		}
	}
}
