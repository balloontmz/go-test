/*
在 goroutine 之间同步数据，来模拟接力 比赛。在接力比赛里，4 个跑步者围绕赛道轮流跑(如代码清单 6-22 所示)。第二个、第三个和 第四个跑步者要接到前一位跑步者的接力棒后才能起跑。
比赛中最重要的部分是要传递接力棒， 要求同步传递。在同步接力棒的时候，参与接力的两个跑步者必须在同一时刻准备好交接。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 为最后一个跑步者计数加 1
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

	//开始比赛
	baton <- 1

	// 等待比赛结束
	wg.Wait()
}

//Runner 跑步者
func Runner(baton chan int) {
	var newRunner int
	//等待接力棒
	runner := <-baton
	// 开始绕着跑到跑步
	fmt.Printf("Runner %d Running With Baton \n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line \n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over \n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d \n",
		runner,
		newRunner)

	baton <- newRunner

}
