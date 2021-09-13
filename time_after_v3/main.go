package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	timer := time.NewTimer(3 * time.Minute)
	defer timer.Stop()
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
		case <-timer.C:
			fmt.Printf("现在是: %d", time.Now().Unix())
		}
	}
}
