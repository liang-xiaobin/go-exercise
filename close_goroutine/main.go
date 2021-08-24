package main

// main 测试goroutine泄露情况
func main() {
	ch := make(chan string, 6)
	go func() {
		for {
			ch <- "hello world"
		}
	}()
}
