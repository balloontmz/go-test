// 暂未详读，先敲完。
package main

import (
	"encoding/gob"
	"fmt"
	"bytes"
	"io/ioutil"
)

type (
	//Post 帖子结构体
	Post struct {
		ID int
		Content string
		Author string
	}
)

func store(data interface{}, filename string)  {
	buffer := new(bytes.Buffer) // 生成指定类型的空指针--生成拥有 Read 和 Write 方法的可变长字节缓冲区
	encoder := gob.NewEncoder(buffer) // 创建 gob 编码器
	err := encoder.Encode(data) // 将数据编码到缓冲区
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600) // 将字节流写入文件
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string)  {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer) //创建解码器
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main()  {
	post := Post{ID:1, Content:"testtest", Author: "tomtiddler"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println("读取到的文件内容为：", postRead)
}