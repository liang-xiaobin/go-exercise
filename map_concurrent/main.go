package main

import "fmt"

func main() {
	s := make(map[string]string)
	for i := 0; i < 9999; i++ {
		go func() {
			s["hello"] = "world"
		}()
	}
	fmt.Printf("len is %d", len(s))
}
