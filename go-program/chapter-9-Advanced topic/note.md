# 进阶话题 Advanced topic

## 9.1 反射 reflection

反射是 Java 出现后迅速流行起来的概念

go 的反射实现了反射的大部分功能，但是相比 Java 还是没那么全。没有像 Java 内置类型工厂，无法通过类型字符串创建对象实例

### 9.1.1 基本概念

*基础的类型检查和转换应该可以通过`*.(type)`的方式进行转换，而不需要用到反射。反射或许用于更为复杂的场景。*

对所有接口进行反射，都可以得到一个`Type`和一个`Value`。`Type`为接口名，`Value`则为该变量实例本身的信息。

### 9.1.2 

1. 获取类型信息
```go
package main

import (
    "fmt"
    "reflect"
)

func main () {
    var x float64 =3.4
    fmt.Println("type:", reflect.TypeOf(x))
    /*
    结果为：
    type:float64
    */

    v := reflect.ValueOf(x)
    fmt.Println("type:", reflect.Type())
    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
    fmt.Println("value:", v.Float())
    /*
    结果为：
    type: float64
    kind is float64: true
    value:3.4
    */
}
```

2. 获取值类型

**go 默认是值传递的，所以采用发射取值的之后对其进行修改其实是无效的，反射本身也会报错。实践方式是反射该值的地址，然后修改该地址内的值！！！**

```go
var x float64 = 3.4
p := reflect.ValueOf(&x) // 得到 x 的地址
fmt.Println("p 的类型是：", p.Type())
fmt.Println("p 是否为可设属性", p.CanSet()) // 此处是否

v := p.Elem()       // 获取指针内值
fmt.Println("v 是否为可设属性", v.CanSet())  // 此属性为可设属性

v.SetFloat(7.1)
fmt.Println(v.Interface())
fmt.Println(x)
```

### 9.1.3 对结构的反射操作

```go
type T struct {
    A int
    B string
}
t := T{203, "mn203"}
s := reflect.ValueOf(&t).Elem()
typeOfT := s.Type()

for i := 0; i < s.NumField(); i++ {
    f := s.Field(i)
    fmt.Print("%d: %s %s = %v\n", i, typeOfT.field(i).Name, f.Type(), f.Interface())
}
/*
输出：
0: A int = 203
1: B string = mh203
*/
```

对于结构的反射操作并没有根本上的不同，只是用了 Field() 方法来按索引获取对应的成员。**在试图修改成员的值时，也需要注意可赋值属性。**

## 9.2 语言交互性--C

## 9.3 链接符号

链接符号关心的是如何将语言文法使用的符号转换为链接期间使用的符号，在常规情况下，链接期间使用的符号对我们不可见。但是一些特殊情况下，我们需要关系这一点。**比如：在用 gdb 调试的时候，要设置断点：b<函数名> , 这里的函数名是指“链接符号”**

```go
//假设在 qbox.us/mockfs 模块中，有如下几个函数：
func New(cfg Config) *MockFS
func (fs *MockFS) Mkdir(dir string) (code int, err error)
func (fs MockFS) Foo(bar Bar)

//它们的链接符号分别为：
qbox.us/mockfs.New
qbox.us/mockfs.*MockFS.Mkdir
qbox.us/mockfs.MockFS.Foo
```

## 9.4 goroutine 机理

从根本上来说 goroutine 就是一种 Go 语言版本的协程(coroutine)。因此，要理解 goroutine 的运作机理，关键是理解传统意义上协程的工作机理

### 9.4.1 协程

**特点：**

+ 能够在单一的系统线程中模拟多个任务的并发执行
+ 在一个特定的时间，只有一个任务在运行，即并非真正地并行
+ 被动的任务调度方式，即任务没有主动抢占时间片的说法。当一个任务正在执行时，外部没有办法中止它。要进行任务切换，只能通过由该任务自身调用 yield() 来主动出让 cpu 使用权
+ 每个协程都有自己的堆栈和局部变量

协程有三种状态：挂起、运行、和停止。挂起代表协程尚未执行完成，但出让了时间片

### 9.4.2 协程的 C 语言实现

`http://swtch.com/libtask` 协程库网址  

### 9.4.3 协程库概述

### 9.4.4 任务

### 9.4.5 任务调度

### 9.4.6 上下文切换

### 9.4.7 通信机制（channel）

**channel 的基本组成**
+ 内存缓存，用于存放元素
+ 发送队列
+ 接收队列

因为协程原则上不会出现多线程编程中经常遇到的资源竞争问题，所以这个 channel 的数据结构甚至在访问的时候都不用加锁。因为Go 语言支持多个核心并发执行多个 goroutine，会造成资源竞争，所以必要的位置还是需要加锁。

## 9.5 接口机制

Go：非侵入式的接口实现

主要用法：
+ 从类型赋值到接口
+ 接口之间相互赋值
+ 接口查询

### note

匿名组合：
关键点是方法的类似继承。即基类的方法会自动导入到子类。如果子类有同名方法，则会被覆盖。
```go
type Base struct {
    Name string
}

func (base *Base) Foo() {}
func (base *Base) Bar() {}

type Foo struct {
    Base
}

func (foo *Foo) Bar(){
    foo.Base.Bar()
}

foo = &Foo{Base{"test"}}
// 基类的方法被覆盖
foo.Bar() != foo.Base.Bar()

// 基类的方法被'继承'
foo.Foo() == foo.Base.Foo()
```

**[c 语言版本的接口实现，代码还需要仔细看。](./interface-2.c)**

### 9.5.2 接口查询

```go
// 查询 IReader 是否实现了 IReadWriter 接口
var reader IReader = NewReader()
if writer, ok := reader.(IReadWriter); ok {
    writer.Write()
}
```

接口匹配过程，按照接口信息表包含的方法是名进行逐一匹配，如果发现传入的信息 是接口方法表的超集，则表示接口查询成功。
**c 语言 -> 访问指针 . 访问成员**
```c
typedef struct _ITbl {
    InterfaceInfo* inter;
    TypeInfo* type;
} ITbl;

ITbl* MakeITbl(InterfaceInfo* intf, TypeInfo* ti) {
    size_t i, n =MemberCount(intf);
    ITbl* dest = (ITbl*)malloc(n * sizeof(void*) + sizeof(ITbl));
    void** addrs = (void**)(dest + 1);

    for (i=0;i<n;i++) {
        addrs[i] = MemberFind(ti, intf->tags[i]);
        if (addrs[i] == NULL){
            free(dest);
            return NULL
        }
    }

    dest->inter = intf;
    dest->type = ti;
}


```

### 9.5.3 接口赋值

接口赋值可以看成是接口查询的优化，因为其在编译器就能确定接口和类型是否匹配。省略了查询的步骤。
