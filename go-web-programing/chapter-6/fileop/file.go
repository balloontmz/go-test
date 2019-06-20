package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	// 第一种读写文件的方法
	//------------------------------------------------------------------
	data := []byte("Hello World!\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Println("第一次读取 data1 的结果为：", string(read1))
	//------------------------------------------------------------------

	// 第二种读写文件的方法 -- 更麻烦，也更灵活。
	//------------------------------------------------------------------
	file1, _ := os.Create("data2")
	defer file1.Close()

	// 写入一个刚创建的文件
	bytes, _ := file1.Write(data)
	fmt.Println("将数据写入文件", data)

	file2, _ := os.Open("data2") // 读取刚写入的文件
	defer file2.Close()

	read2 := make([]byte, len(data)) // 切片本身本质是一个指针类型。
	bytes, _ = file2.Read(read2)
	fmt.Println("文件中读取到的字节为：", bytes) // 返回的是字节长度？
	fmt.Println("read2为：", read2)
}