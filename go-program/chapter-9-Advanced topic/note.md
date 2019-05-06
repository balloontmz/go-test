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