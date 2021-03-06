//连接池的意义： 主要是尽可能的减少创建资源的时间，对于需要复用的资源很好用。具体池子大小最好根据测试进行。。。毕竟在池内没有资源的时候，能创建资源进行使用。
package main

import (
	"io"
	"log"
	"math/rand"
	"pool/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25 // 要使用的 goroutine 的数量
	pooledResources = 2  // 池中的资源的数量
)

//dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

//idCounter 用来给每个连接分配一个独一无二的 id
var idCounter int32

//createConnection 是一个工厂函数，当需要一个新链接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池里的链接来完成查询
	for query := 0; query < maxGoroutines; query++ {
		// 每个 goroutine 需要自己复制一份要查询值的副本，不然所有的查询会共享同一个查询变量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	// 等待 goroutine 结束
	wg.Wait()

	log.Println("Shutdown Program.")
	// 关闭池 -- 同时会释放仍然存在的资源，一般应该是池子大小
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	// 从池中请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将该连接释放回池里
	defer p.Release(conn)

	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
