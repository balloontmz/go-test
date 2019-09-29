package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 捕获 counter 的值
		value := counter

		fmt.Println("id 为", id, "前置计算")

		// 当前 goroutine 从线程退出，并放回到队列  -- 手动操作
		runtime.Gosched()

		fmt.Println("id 为", id, "后置计算")

		// 增加本地变量 value 的值
		value++

		// 将该值保存会 counter
		counter = value
	}
}
