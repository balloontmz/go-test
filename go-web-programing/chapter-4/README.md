# chapter-4 
浏览器向服务器发送请求时， Fragment （哈希片段）字段将会被剔除 。

## 请求和响应

### Request 结构

URL 结构

Header 字段

Body 字段

Form 等字段

### Url 结构

### 请求首部

设置首部和添加首部的区别： 设置首部会将之前有的值置空并添加，添加首部会将值添加到已有值后方。

### 请求主体

请求主体放在 Request.Body 中，由于 Go 提供诸如 FormValue 这样的方法提取表单，所以一般情况并不需要提取未经处理的表单。

## Go 和 HTML 表单

在绝大多数情况下，POST 请求都是通过 HTML 表单发送的？？？--json？

如果把表单的 enctype 属性的值设置为 x-www-form-urlencoded ，那么浏览器将把 HTML 的数据编码成一个连续的 “长查询字符串”。但是，如果我们把 enctype 设置为 multipart/form-data ，那么表单中的数据将被转换成一条 MIME 报文。如果表单只是传递简单的文本数据，第一种更为合适，如果需要传递文件等大量数据，第二种更为合适。

GET 方法也是能发送表单请求的。但是由于 GET 没有请求体，但是表单数据会以键值对的形式存在于 URL 里面。

### Form 字段

r.Form 能访问表单或者 url 提供的键值，如果只想访问表单，可以通过 PostForm 字段。PostForm 只支持 x-www-form-urlencoded 不支持 form-data。要获取 formm-data 的表单值，可以通过 MultiPartForm 字段 -- 不包括 url

FormValue 能获取指定表单键的值，即使有多个值，也会只获取第一个。

### 文件

### 处理带有 json 主体的 post 请求

## ResponseWriter

ResponseWriter 其实是 http.response 非导出结构的指针接口。

Write
writeHeader
Header

Write 方法接收一个字节数组作为参数，并为数组中的字节写入 HTTP 响应的主体中。如果用户在使用 write 方法执行写入操作的时候没有为首部设置相应的内容类型，那么响应的内容类型通过检测被写入的前 512 个字节决定。

WriteHeader 为响应写入 HTTP 响应状态码

Header 为响应写入首部。

WriteHeader 在执行完毕后就不在允许对首部进行写入。

## Cookie

cookie 是一种存储在客户端、体积较小的信息。cookie 分为会话 cookie 和持久 cookie。

cookie 用于客服 http 的无状态属性。

Expires 用于指定过期时间，没有设置该属性的 cookie 称为临时 cookie，当浏览器关闭时该 cookie 会自动删除。设置了该属性的 cookie 会在到期时间到达时自动删除。Expires 和 Maxage 都是用于设置过期时间，但是 Expires 应用更为广泛，HTTP 1.1 删除了Expires，推荐使用 Maxage。Expires 指定的是一个时间戳，Maxage 指定的是从页面被打开到 cookie 过期的秒数。

### 将 cookie 发送给浏览器 和接收请求的 cookie

### 使用 cookie 实现闪现消息

通过一个临时会话 cookie 传递消息，通过找到该 cookie 显示消息。通过设置过期的持久会话删除 cookie。    