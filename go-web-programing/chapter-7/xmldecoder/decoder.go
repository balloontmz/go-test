package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type (
	Post struct {
		XMLName xml.Name `xml:"post"`	// 这个类型？
		ID string `xml:"id,attr"`
		Content string `xml:"content"`
		Author Author `xml:"author"`
		Xml string `xml:",innerxml"`
		Comments []Comment `xml:"comments>comment"` // 访问子元素
	}
	Author struct {
		ID string `xml:"id,attr"`
		Name string `xml:",chardata"`
	}
	Comment struct {
		ID string `xml:"id,attr"`
		Content string `xml:"content"`
		Author string `xml:"author"`
	}
)

func main(){
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file", err)
		return
	}
	defer xmlFile.Close()

	// 向 NewDecoder 传递一个 io.Reader 创建一个 Decoder
	decoder := xml.NewDecoder(xmlFile) // 根据给定的 xml 数据生成解码器
	for{
		t, err := decoder.Token() // 每进行一次迭代，就从解码器中获取一个 token
		if err == io.EOF { // 所有 token 读取完毕，返回 eof
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se :=t.(type) {
		case xml.StartElement: // 判断是否是起始标签
			if se.Name.Local == "comment" { // 是否是指定标签
				var comment Comment
				decoder.DecodeElement(&comment, &se)
			}
		}
	}
}