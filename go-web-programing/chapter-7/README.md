# chapter-7

soap 风格和 restful 风格

## web 服务

web 服务是与其他软件程序进行交互的采用 http 进行交互的程序。

soap 通常是 rpc 的，功能驱动。restful 服务通常是数据驱动。

restful：速度快，构建简单

soap：安全且健壮。

## 基于 soap 的 web 服务

## 基于 rest 的 web 服务

Representational State Transfer -- 具象状态传输 -- 设计理念

PUT 方法是幂等的，POST 方法不是。多次请求，PUT 只会创建一次资源，POST 会有多次。

REST 完成理论上的任意动作：
1. 把过程抽象成对象，或者把动作转换为名词，然后将其用作为资源
2. 将动作用作资源的属性。

## 分析和创建 xml

### 分析 xml
反引号包围的信息为结构标签。结构标签为键值对，键不能包含空格、引号和冒号。值用引号包围。处理 xml 时，键总是 xml

xml 部分使用规则：
1. 通过创建 xml.Name 字段，将 xml 元素的名字存储在这个字段里面，在一般情况下，结构的名字就是元素的名字
2. 通过创建一个与 xml 元素同名的字段，并使用 `` "`xml:"<name>,attr`"`` 作为该字段的结构标签，可以将元素的 `<name>`属性的值存储在这个字段里面。
3. 通过创建一个与 xml 元素标签同名的字段，并使用 `` "`xml:",chardata"`" `` 作为该字段的结构标签，可以将 xml 元素的字符数据存储到这个字段里面。
4. 通过定义一个任意名字的字段，并使用`` "`xml:",innerxml"`" `` 作为该字段的结构标签，可以将 xml 元素的原始 xml 存储到这个字段里面。
5. 没有模式标志（比如,arrt、,chardata、或者,innerxml）的结构字段将于同名的 xml 元素匹配。
6. 使用 `` "`xml:"a>b>c"`" `` 这样的结构标签可以在不指定树状结构的情况下直接获取指定的 xml 元素，其中 a 和 b 为中间元素，而 c 则是想要获取的节点元素。

对于直接将 xml 解析到结构体的方式，只适合较小的 xml 文件，对于流文件传输的较大文件，需要采用 decoder 来代替 Unmarshal 来进行手动解码。

手动解码 xml 需要做更多的工作，所以并不适用于处理小型 xml 文件。

### 创建 xml
MarshalIndent 能生成更好看的 xml！！！--接收两个额外参数，一个前缀、一个缩进（制表符"\t"）

xml.Header 添加在 xml 字符串之前为输出的 xml 添加声明。

也可以采用 Encoder 将 xml输出到文件当中

## 分析和创建 json

### 分析 json

1. 创建一些用于包含 json 数据的结构
2. 通过 json.Unmarshal 函数，把 json 数据解封到结构里面。

通用规则：将字段的结构标签设置为：``" `json:"<name>"` "``

根据输入决定采用 Decoder 还是 Unmarshal。如果 json 来自 io.Reader 流，那么使用 Decoder 更好。如果 json 数据来源于字符串或者内存中的某个地方，那么使用 Unmarshal 更好。

MarshalIndent 将结构封装成格式化的 json，更美观。

### web 服务

### 总结：
1. 编写 web 服务是 go 目前非常常见的用途之一
2. web 服务主要分为两种类型：基于 soap 的 web 服务和基于 rest 的 web 服务
3. 创建和分析 xml 以及 json 的步骤都是相似的。