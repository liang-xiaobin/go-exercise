package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	for {
		go func() {
			_, err := http.Get("https://www.qq.com")
			if err != nil {
				fmt.Printf("http.Get err: %v\n", err)
			}
		}()
		time.Sleep(time.Second * 1)
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
