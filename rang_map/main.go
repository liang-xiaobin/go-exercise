package main

import "log"

func main() {
	m := make(map[int]string)
	m[0] = "1"
	m[1] = "2"
	m[2] = "3"
	m[3] = "4"
	m[4] = "5"
	for k, v := range m {
		log.Printf("k:%v, v:%v", k, v)
	}
}
