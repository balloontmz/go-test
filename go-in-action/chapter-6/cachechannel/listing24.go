/*
由于程序和 Go 语言的调度器带有随机成分，这个程序每次执行得到的输出会不一样。不过，通过有缓冲的通道，使用所有 4 个 goroutine 来完成工作，这个流程不会变。从输出可以看到每 个 goroutine 是如何接收从通道里分发的工作。
在 main 函数的第 31 行，创建了一个 string 类型的有缓冲的通道，缓冲的容量是 10。在 第 34 行，给 WaitGroup 赋值为 4，代表创建了 4 个工作 goroutine。之后在第 35 行到第 37 行， 创建了 4 个 goroutine，并传入用来接收工作的通道。在第 40 行到第 42 行，将 10 个字符串发送 到通道，模拟发给 goroutine 的工作。一旦最后一个字符串发送到通道，通道就会在第 46 行关闭， 而 main 函数就会在第 49 行等待所有工作的完成。
第 46 行中关闭通道的代码非常重要。当通道关闭后，goroutine 依旧可以从通道接收数据， 但是不能再向通道里发送数据。能够从已经关闭的通道接收数据这一点非常重要，因为这允许通 道关闭后依旧能取出其中缓冲的全部值，而不会有数据丢失。从一个已经关闭且没有数据的通道 里获取数据，总会立刻返回，并返回一个通道类型的零值。如果在获取通道时还加入了可选的标 志，就能得到通道的状态信息。
在 worker 函数里，可以在第 58 行看到一个无限的 for 循环。在这个循环里，会处理所有 接收到的工作。每个 goroutine 都会在第 60 行阻塞，等待从通道里接收新的工作。一旦接收到返 回，就会检查 ok 标志，看通道是否已经清空而且关闭。如果 ok 的值是 false，goroutine 就会 终止，并调用第 56 行通过 defer 声明的 Done 函数，通知 main 有工作结束。
如果 ok 标志是 true，表示接收到的值是有效的。第 71 行和第 72 行模拟了处理的工作。 一旦工作完成，goroutine 会再次阻塞在第 60 行从通道获取数据的语句。一旦通道被关闭，这个 从通道获取数据的语句会立刻返回，goroutine 也会终止自己。
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的 goroutine 数量
	taskLoad         = 10 //要处理的工作的重量
)

var wg sync.WaitGroup

func init() {
	// 初始化随机数种子 -- 干嘛用的？？？
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作
	wg.Add(numberGoroutines)
	for gr := 0; gr < numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 0; post < taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// 当所有工作都处理完时关闭通道 以便所有 goroutine 退出
	close(tasks)
	// 等待所有工作完成
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已经关闭
			fmt.Printf("Worker: %d: Shutting Down \n", worker)
			// 此时执行 defer，关闭 goroutine
			return
		}
		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
