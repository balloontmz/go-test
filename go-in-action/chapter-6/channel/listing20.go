package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	// 用处？
	rand.Seed(time.Now().UnixNano())
}

// 网球比赛
func main() {
	// 创建一个无缓冲通道
	court := make(chan int)

	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 启动两个选手
	go player("Nadal", court)
	go player("Djokovic", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won \n", name)
			return
		}

		// 选择随机数，然后用这个数判断是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed Number %d \n", name, n)

			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加 1
		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		court <- ball
	}
}
