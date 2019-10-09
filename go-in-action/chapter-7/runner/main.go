package main

import (
	"log"
	"os"
	"time"

	"runner/runner"
)

const timeout = 4 * time.Second

/*
代码清单 7-13 的第 16 行是 main 函数。在第 20 行，使用 timeout 作为超时时间传给 New 函数，并返回了一个指向 Runner 类型的指针。之后在第 23 行，
使用 createTask 函数创建了 几个任务，并被加入 Runner 里。在第 42 行声明了 createTask 函数。这个函数创建的任务只 是休眠了一段时间，用来模拟正在进行工作。
增加完任务后，在第 26 行调用了 Start 方法，main 函数会等待 Start 方法的返回。

当 Start 返回时，会检查其返回的 error 接口值，并存入 err 变量。如果确实发生了错 误，代码会根据 err 变量的值来判断方法是由于超时终止的，还是由于收到了中断信号终止。
 如果没有错误，任务就是按时执行完成的。如果执行超时，程序就会用错误码 1 终止。如果接收 到中断信号，程序就会用错误码 2 终止。其他情况下，程序会使用错误码 0 正常终止。
*/
func main() {
	log.Println("start work.")

	// 为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminationg due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
