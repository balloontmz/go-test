// 根据命令行参数获取队列
// 并根据获取到的队列拉取消息
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	// 获取从命令传递的队列名称和超时设置
	var name string
	var timeout int64
	flag.StringVar(&name, "n", "", "Queue name")
	flag.Int64Var(&timeout, "t", 20, "(Optional) tomeout in seconds for long polling")
	flag.Parse()

	if len(name) == 0 {
		flag.PrintDefaults()
		exitErrorf("Queue name required")
	}

	// 初始化 sdk 将从 ~/.aws/credentials 加载凭据的 session，本地环境变量优于从文件加载
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// 创建 sqs 客户端
	svc := sqs.New(sess)

	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{ // 获取队列 url
		QueueName: aws.String(name),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			exitErrorf("未找到队列 %q", name)
		}
		exitErrorf("无法连接到队列 %q, %v", name, err)
	}

	// 调用 ReceiveMessage 从队列中获取最新消息
	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{ // 此处构造方式和 recieveanddeletemessage.go 方式有差别，可用于比较
		QueueUrl: resultURL.QueueUrl,
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
		}),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: aws.Int64(timeout),
	})

	if err != nil {
		exitErrorf("无法从队列获取消息 %q, %v", name, err)
	}

	fmt.Printf("收到 %d 条消息\n", len(result.Messages))
	if len(result.Messages) > 0 {
		fmt.Println(result.Messages)
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
