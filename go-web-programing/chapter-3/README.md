# chapter-3 接收请求

## 处理器和处理器函数

处理器和处理器函数--处理器函数是拥有和处理器相同行为的函数。（分别对应 http.Handle 和 http.HandlerFunc）

处理器函数只是创建处理器的一个简便方法，`helloHandler := HandlerFunc(hello)` --全局函数？？？--将处理器函数转换为处理器。

处理器函数虽然简便，但是某些情况下并不能完全替代处理器。这是因为现在某些情况下，代码可能已经包含了某个接口或者某种类型，而此时只要实现一个 ServeHTTP 函数就能实现 处理器接口了。 

## 串联处理器和处理器函数

串联处理器和处理器函数是一种非常常见的惯用法，很多 web 应用框架都使用这一技术

## 路由匹配

*最小惊讶原则(The Principle of Least Surprise)*

路由如果 以 `/` 结尾，将能部分匹配路由，否则，只能完全匹配。

go 的 处理器接口能非常方便地实现多路复用器。-- HttpRouter(轻量级第三方多路复用器)

## 生成 ssl 证书

CA: VeriSign Thawte Comodo SSL -- CA 提供商

自行生成证书的办法有很多中，其中一种就是使用 Go 标准库的 crypto 包群(library group)

暂时搁置，进行下一步。

## 使用 HTTP/2

依赖 https 暂时搁置，暂存。

## 使用 curl