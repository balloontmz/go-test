//Package worker 适用于大量需要时间处理的任务
package worker

import (
	"sync"
)

//Worker 必须满足接口类型才能使用工作池
type Worker interface {
	Task()
}

//Pool 提供一个 goroutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

//New 创建一个新的工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task() // 此处会阻塞、并等待，除非通道关闭
			}
			p.wg.Done()
		}()
	}
	return &p
}

//Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//Shutdown 等待所有 goroutine 停止工作
func (p *Pool) Shutdown() {
	close(p.work) // 关闭通道相当于中断所有工作池 goroutine 的 for 循环、WaitGroup 执行了 Done 操作
	p.wg.Wait()
}
