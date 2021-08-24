package main

import (
	"context"
	"fmt"
	"time"
)

//main 使用context关闭goroutine
func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("hello world")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	<-ch
	fmt.Println("end")
}
