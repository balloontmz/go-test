# go in action

## chapter-3

[go 环境配置--自动补全等](https://zhuanlan.zhihu.com/p/36453771)

```shell
echo "github.com/nsf/gocode
github.com/uudashr/gopkgs/cmd/gopkgs
github.com/ramya-rao-a/go-outline
github.com/acroca/go-symbols
golang.org/x/tools/cmd/guru
golang.org/x/tools/cmd/gorename
github.com/fatih/gomodifytags
github.com/haya14busa/goplay/cmd/goplay
github.com/josharian/impl
github.com/rogpeppe/godef
sourcegraph.com/sqs/goreturns
github.com/golang/lint/golint
github.com/cweill/gotests/gotests
github.com/derekparker/delve/cmd/dlv" | xargs go get -u -v
```

nil 切片和空切片

```go
// nil 切片
var slice []int

// 空切片
slice := make([]int, 0)

slice := []int{}
```

### 创建切片时的 3 个索引
创建切片时第三个索引用于指定生成的切片的容量--作用：当扩展切片时，不会改变原切片的数组，而是重新生成一个原始数组
`slice := source[2:3:4]`

如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个 append 操作 创建新的底层数组，与原有的底层数组分离。新切片与原有的底层数组分离后，可以安全地进行 后续修改.

range 遍历返回的是当前遍历元素的副本，该值的指针总是指向相同的地址，如果需要该元素的地址，应该采用 索引进行获取

切片遍历总是从头部开始，所以是有序的？？？map 是无序的吧。。。