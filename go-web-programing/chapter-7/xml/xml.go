/*
将 xml 解析到指定的结构体，这种处理只适合小型 xml，对于大型 xml，需要采用 decoder 进行手动解码并迭代。
*/
package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type (
	Post struct {
		XMLName xml.Name `xml:"post"`	// 这个类型？
		ID string `xml:"id,attr"`
		Content string `xml:"content"`
		Author Author `xml:"author"`
		Xml string `xml:",innerxml"`
		Comments []Comment `xml:"comments>comment"`
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

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data", err)
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}