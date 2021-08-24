package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 6)
	go func() {
		for {
			// v, ok := <-ch
			// if !ok {
			// 	fmt.Println("end")
			// 	return
			// }
			// fmt.Println(v)
			for v := range ch {
				fmt.Println(v)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	ch <- "hello world1"
	ch <- "hello world2"
	close(ch)
	time.Sleep(time.Second)
}
