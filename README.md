
go 相关的练习
=============
## aws 相关

### sqs 相关
```
VisibilityTimeout  // 可见性超时
WaitTimeSeconds     // 此属性大于1 小于 20 代表设置为长轮询
ReceiveMessageWaitTimeSeconds  //（队列属性）此值设为 0，并且消息属性  WaitTimeSeconds 设为0，代表短轮询
```
[aws sqs 用户凭证配置文档](https://docs.aws.amazon.com/zh_cn/sdk-for-go/v1/developer-guide/configuring-sdk.html)

[可见性超时文档--目测比较重要的一个属性](https://docs.aws.amazon.com/zh_cn/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)

[创建 sqs 队列](aws-demo/sqs-demo/sqs_createqueues.go)

[创建并删除 sqs 队列](aws-demo/sqs-demo/sqs-createanddeletequeue)

[根据名字获取队列 url](aws-demo/sqs-demo/sqs-getqueueurl)

[获取队列 url 列表](aws-demo/sqs-demo/sqs-listqueue)

[采用长轮询接收消息](aws-demo/sqs-demo/sqs-longpolling-receive-message)

[采用短轮询接收消息并删除](aws-demo/sqs-demo/sqs-recieveanddeletemessage)

[发送消息](aws-demo/sqs-demo/sqs-sendmessage)

[sqs 队列中的可见性超时]

### queue 相关

队列测试
`go test '../go-test/queue/queue-test-1/queue_test.go'`

### 测试上传


### 编程语言自己实现自己的解释器 -- 自举

### 类型的方法--接收者的方法

1. 定义的接收者为指针，调用者可以是指针或者值，会自动转换

2. 定义的接收者为值，调用者可以是指针或者值，会自动转换

**但是不应该混用**，比如指针类型实现了某个接口，那么该值类型并未实现该接口。

### 从一个接口中取到某个类型的值的方法：

```go
1. t := i.(T)  // 从 interface i 中获取 类型 T 的值并赋值给 t，如果 i 中不存在该值，则触发一个 panic

2. t, ok := i.(T) // 从 interface i 中获取 类型 T 的值并赋值给 t，把取值结果赋值给 ok。如果没有取到 T 的值，则 ok 值为 false，t 的值为 零值。
```

### 程序的集合是可数的

所有程序的集合虽然是无穷的。但是程序是由有限种字符有限次排列成的，那么每个程序就能唯一对应一个整数

## 参考书
Go in action
The Go Programming Language