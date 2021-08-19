package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			printMemStats()
			time.Sleep(5 * time.Second)
		}
	}()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8888")
	if err != nil {
		log.Printf("router init err:(%v)", err)
	}
}

// printMemStats print runtime info
// Alloc: golang语言框架堆空间分配的字节数
// HeapIdle: 申请但是未分配的堆内存或者回收了的堆内存(空闲)字节数
// HeapReleased: 返回给OS的堆内存
// HeapInuse uint64: 正在使用的堆内存字节数
// GCSys uint64: 垃圾回收标记元信息使用的内存
// Sys uint64: 服务现在系统使用的内存
func printMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf(
		"====> Alloc:%d(bytes), HeapIdle:%d(bytes), "+
			"HeapReleased:%d(bytes), HeapInuse:%d(bytes), "+
			"GCSys:%d(bytes), Sys:%d(bytes)",
		ms.Alloc, ms.HeapIdle, ms.HeapReleased, ms.HeapInuse, ms.GCSys, ms.Sys,
	)
}
