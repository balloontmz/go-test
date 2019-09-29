package main

import (
	"fmt"
	"runtime"
	"sync"
)

//main 是所有 GO 程序的入口
func main() {
	// 分配一个逻辑处理器给调度器使用
	// runtime.GOMAXPROCS(1)  // 单核心  -- 并发
	runtime.GOMAXPROCS(runtime.NumCPU()) // 多核心 -- 并行
	fmt.Print("当前 cpu 数量为：", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\n Terminating Program")

}
