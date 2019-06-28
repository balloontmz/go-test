#chapter-9

## 并行和并发的区别

## goroutine

### 等待组 WaitGroup
```go
var sync.WaitGroup
wg.Add(2)
wg.Done()
wg.Done()
wg.Wait()
```
1. 声明一个等待组
2. 使用 Add 方法为等待组计数器设置值。
3. 当一个 goroutine 完成它的工作时，使用 Done 方法对等待组的计数器执行减一操作。
4. 调用 Wait 方法，该方法将一直阻塞，直到等待组计数器的值变为 0

## 通道
通过 channel 在多个 goroutine 中通信。 -- 引用类型。make 创建

**new 函数分配类型，返回指向新分配类型零值得指针**
**make 创建 slice map chan ，返回这些类型的引用** -- 这些类型内部存的是指针。--引用类型

### 通过通道实现消息传递

### 有缓存通道

### 从多个通道选择
select case 如果多个 case 都有值，就会默认选择一个 case 执行。

## web 应用并发

Upstart systemd 守护进程