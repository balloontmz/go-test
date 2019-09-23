package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher, 0)

//Run 执行搜索逻辑
func Run(searchTerm string)  {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err) // 打印错误并终止程序
	}

	// 创建一个无缓冲通道 接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup 以便处理所有的数据源
	var waitGroup sync.WaitGroup


	// 设置需要等待处理每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()  // 这个参数实际是
		}(matcher, feed)  // 此处参数采用传递的方式而不是闭包的方式 -- 如果采用闭包，goroutine 可能处理的是同一个数据
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等待所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式 通知 display 函数可以退出程序了
		close(results)
	}()

	// 启动函数，显示返回的结果，并且在最后一个结果显示完成之后返回
	Display(results)
}

//Register 调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher)  {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}