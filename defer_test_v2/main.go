package main

import "fmt"

func main() {
	var w [6]struct{}
	for i := range w {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}
