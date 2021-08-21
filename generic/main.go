package main

import "fmt"

func print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	print([]string{"您好","梁晓斌"})
	print([]int64{1,2,3,4})
}