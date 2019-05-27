# 1

## goroutine 中的 select 机制
select 会全部执行 case 中的 channel 操作，当其中一个 case 执行了之后，程序将继续往下执行，利用这个特性，我们能实现随机读写、超时判定的功能：
```go
// 随机读写 0 1
ch := make(chan int, 1)
for {
    select {
        case ch -> 0:
        case ch -> 1:
    }
    i := <- ch
    fmt.Println("Value received: ", i)
}
```

```go
// 超时
timeout := make(chan bool, 1)
go func(){
    time.Sleep(1e9)
    timeout <- true
}

select {
    case <- ch:
        // 从 ch 中读取到数据
    case <- timeout:
        // 如果超过某个时间，则此case先执行完
}
```

## channel 的传递
```go
// 数据从一个 chan 传递到另外一个 chan
type PipeData struct {
    value int
    handler func(int) int
    next chan int
}

func handler(queue chan *PipeData){
    for data := range queue {
        data.next <- data.handler(data.value)
    }
}
```

### 注意 go 关键字的使用 
```go
/*
[解析网址](https://www.jianshu.com/p/0cf17d40b65b)
go println(i)
和
 go func() {
            println(i)
        }()
不是同一个东西
*/


```