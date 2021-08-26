package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 统计实例实际创建次数
var numCalcsCreated int32

// 创建实例
func createBuffer() interface{} {
	//使用原子加,处理并发问题
	atomic.AddInt32(&numCalcsCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func main() {
	// 创建实例
	bufferPool := &sync.Pool{
		New: createBuffer,
	}
	// 多goroutine并发测试
	numWorks := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorks)
	for i := 0; i < numWorks; i++ {
		go func() {
			defer wg.Done()
			// 申请一个buffer实例
			buffer := bufferPool.Get()
			//buffer := createBuffer()
			_ = buffer.(*[]byte)
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	fmt.Printf("%d buffer objects were created \n", numCalcsCreated)
}
