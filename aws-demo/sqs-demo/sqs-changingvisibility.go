// 设置队列的可见性超时
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
	var name string
	flag.StringVar(&name, "n", "", "Queue name")
	flag.Parse()

	if len(name) == 0 {
		flag.PrintDefaults()
		exitErrorf("Queue name required")
	}
	// 初始化 sdk 将从共享凭证文件 ~/.aws/credentials 加载凭证的会话
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// 创建 sqs 服务客户端
	svc := sqs.New(sess)
	// 从队列中获取消息，调用 ReceiveMessage。传入队列的 url 以返回队列中下一条消息的详细信息。如果没有收到消息，打印任何返回的错误，否则打印消息

	// 获取队列 url
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			exitErrorf("未找到队列 %q", name)
		}
		exitErrorf("无法连接到队列 %q, %v", name, err)
	}
	fmt.Print("队列 url 为", resultURL.QueueUrl)
	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl: resultURL.QueueUrl,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// 检查是否有消息
	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		return
	}

	// 如果返回了消息，则使用其接收的句柄将可见性超时设置为 30 s
	duration := int64(30)
	resultVisibility, err := svc.ChangeMessageVisibility(&sqs.ChangeMessageVisibilityInput{
		ReceiptHandle:     result.Messages[0].ReceiptHandle,
		QueueUrl:          resultURL.QueueUrl,
		VisibilityTimeout: &duration,
	})

	if err != nil {
		fmt.Println("Visibility Error", err)
		return
	}

	fmt.Println("Time Changed", resultVisibility)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
