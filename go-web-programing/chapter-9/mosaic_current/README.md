# mosaic 生成马赛克图片应用
使用并发特性提高 go web 应用的性能

本应用创建一个对图片进行马赛克处理，一次生成马赛克图片的程序。

马赛克处理指的是将图片分割成多个（通常是相同大小的）矩形截面，然后使用一些瓷砖图片（tile picture）去代替截面原有的照片。 -- 拷贝的瓷砖图片尺寸为 150x150

接收用户上传的图片，然后据此生成马赛克图片。

## 创建马赛克图片
创建马赛克图片的第一步是定义一个马赛克算法。

1. 通过扫描图片目录，并使用图片的文件名作为键、图片的平均颜色作为值，构建出一个由瓷砖图片组成的散列（map），也就是一个瓷砖图片数据库。通过计算图片中每个像素红绿蓝三种颜色的总和，并将它们除以像素的总数量，我们得到一个三元组，而这个三元组就是图片的平均颜色。

2. 根据瓷砖图片的大小，将目标图片分割成一系列尺寸更小的子图片。

3. 对于目标图片切割出的每张子图片，将它们位于左上方的第一个像素设定为该图片的平均颜色。--分隔是根据尺寸还是像素？？？填充是根据尺寸还是像素？？？尺寸和像素的关系是什么？？？--子图片已分隔，可以认为元素是够小的。 -- 当目标文件够大但是尺寸正常时，为什么会分割得很小很小

4. 根据子图片的平均颜色，在瓷砖图片数据库中找出一张平均颜色与之最为接近的瓷砖图片，然后在目标图片的相应位置上使用瓷砖图片去代替原有的子图片。为了找出最接近的平均颜色，程序需要将子图片的平均颜色以及瓷砖图片的平均颜色都转换成三维空间的一个点，并计算这两个点之间的欧几里得距离。

5. 当一张瓷砖图片被选中之后，程序就会把这张图片从瓷砖图片数据库中移除以此保证马赛克图片中的每张瓷砖图片都是独一无二，与众不同的。--此处需要注意瓷砖数据库的量需要够多。

## 总结
1. 并发能显著提升性能
2. 需要注意在并行执行完成之后再进行时间统计
3. 瓷砖图片需要对其进行重新指定尺寸 -- 契合图片的分割大小。
GOMAXPROCS 指定运行程序采用的 cpu 数量

1. go web 服务器本身是并发的，服务器会把接收到的每条请求都放到独立的 goroutine 里运行
2. 并发和并行是两个相辅相成的概念，但他们并不相同。并发指的是两个或多个任务在同一时间段内启动、运行和结束，并且这些任务可能会彼此互动，而并行则是单纯地运行多个任务
3. go 通过 goroutine 和通道这两个重要的特性直接支持并发，但 go并不直接支持并行。
4. goroutine 用于编写并发程序，而通道则用于为不同的 goroutine 之间提供通行功能。
5. 无缓存通道都是同步的，尝试向一个已经包含数据的无缓存通道推入新的数据将被阻塞；但是，有缓冲通道在填满之前都是异步的。
6. select 语句可以以先到先服务的方式，从多个通道中选出一个已经准备好执行接收操作的通道。--需要回看，如果没有 default 会不会阻塞。--select 的机制！！！
7. WaitGroup 同样可以用于对多个通道进行同步
8. 并发程序的性能一般会比非并发程序的性能好，而具体提升多少则取决于使用的算法。
9. 在条件允许的情况下，并发的文本、应用将自动获得并行的优势。