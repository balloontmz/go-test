// 创建队列 并且 删除队列
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

	// 创建队列
	result1, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String("SQS_QUEUE_TEMP"),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),    // 60s 是啥
			"MessageRetentionPeriod": aws.String("86400"), // 一天
		},
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result1.QueueUrl)

	// 删除队列
	result3, err := svc.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: aws.String(*result1.QueueUrl),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", result3)
}
