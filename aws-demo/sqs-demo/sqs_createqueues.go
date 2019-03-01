// 创建队列
// 创建 fifo 队列必须 在队列名加上 .fifo
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// 创建 sqs 客户端
	svc := sqs.New(sess)

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String("SQS_QUEUE_EMAIL_DEV"),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),    // 60s 是啥
			"MessageRetentionPeriod": aws.String("86400"), // 一天
		},
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.QueueUrl)
}
