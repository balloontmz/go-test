package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // 得到 x 的地址
	fmt.Println("p 的类型是：", p.Type())
	fmt.Println("p 是否为可设属性", p.CanSet()) // 此处是否

	v := p.Elem()       // 获取指针内值
	fmt.Println("v 是否为可设属性", v.CanSet())  // 此属性为可设属性

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}