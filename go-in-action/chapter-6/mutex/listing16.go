/*
互斥锁这个名字来自互斥(mutual exclusion)的概念。互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以 执行这个临界区代码。我们还可以用互斥锁来修正代码清单 6-9 中创建的竞争状态
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	//创建两个 goroutine
	go incCounter(1)
	go incCounter(2)

	//等待 goroutine 结束
	wg.Wait()
	fmt.Printf("Final Counter: %d \n", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
