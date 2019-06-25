package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type (
	Post struct {
		XMLName xml.Name `xml:"post"`	// 这个类型？
		ID string `xml:"id,attr"`
		Content string `xml:"content"`
		Author Author `xml:"author"`
	}
	Author struct {
		ID string `xml:"id,attr"`
		Name string `xml:",chardata"`
	}
)

func main() {
	post := Post{
		ID: "1",
		Content: "this is a test",
		Author: Author{
			ID: "2",
			Name: "tomtiddler",
		},
	}
	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t") // 美化的输出
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error Encoding XML to file", err)
		return
	}
}