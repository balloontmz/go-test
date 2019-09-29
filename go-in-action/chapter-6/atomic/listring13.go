package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64

	wg sync.WaitGroup
)

func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wg.Wait()

	// 显示最终的值
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)

		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()
	}
}
