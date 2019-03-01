// 获取队列列表
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	// go run sqs_listqueues.go
	// 在共享凭证文件 ~/.aws/credentials 中查找凭证，此方法优先级低于环境变量
	// https://docs.aws.amazon.com/zh_cn/sdk-for-go/v1/developer-guide/configuring-sdk.html
	// https://docs.aws.amazon.com/zh_cn/sdk-for-go/v1/developer-guide/sqs-example-create-queue.html
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// 创建一个 sqs 服务客户端
	svc := sqs.New(sess)

	// 列出在指定地区可用的队列
	result, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success")
	// 由于这些结果是指针，直接打印出来是没有用的
	for i, urls := range result.QueueUrls {
		// 避免不相关的 nil
		if urls == nil {
			continue
		}
		fmt.Printf("%d: %s\n", i, *urls)
	}
}
