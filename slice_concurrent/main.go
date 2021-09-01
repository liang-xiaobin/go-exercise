package main

import "fmt"

func main() {
	var s []string
	for i := 0; i < 9999; i++ {
		go func() {
			s = append(s, "hello world")
		}()
	}
	fmt.Printf("len is %d", len(s))
}
