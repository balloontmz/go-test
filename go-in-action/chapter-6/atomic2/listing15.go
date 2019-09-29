package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	go doWork("A")
	go doWork("B")

	// 给定 goroutine 执行时间
	time.Sleep(1 * time.Second)

	// 该停止工作了，安全地设置 shutdown 标志
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// 等待 goroutine 结束
	wg.Wait()
}

func doWork(name string) {
	// 在函数退出时调用 DOne 来通知 main 函数工作已完成
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work \n", name)
		time.Sleep(250 * time.Millisecond)

		// 要停止工作了吗？
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down \n", name)
			break
		}
	}
}
