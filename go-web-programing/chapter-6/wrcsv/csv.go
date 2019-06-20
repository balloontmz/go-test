package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

type (
	//Post 帖子结构体
	Post struct {
		ID int
		Content string
		Author string
	}
)

func main()  {
	// 问题：读写操作没有指定分隔符？-- 默认了逗号分隔？还是 csv 就是逗号分隔的。
	// 写文件
	//---------------------------------------------------------------
	csvFile, err := os.Create("post.csv")//创建 csv 文件
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{ID:1, Content:"test", Author: "tom"},
		Post{ID:2, Content:"test", Author: "tiddler"},
		Post{ID:3, Content:"test", Author: "tom"},
		Post{ID:4, Content:"test", Author: "balloon"},
	}

	writer := csv.NewWriter(csvFile)  // 创建一个写入器？

	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush() // 将写入器中的缓存写入文件？--清空缓冲区。
	//---------------------------------------------------------------

	// 读文件
	//---------------------------------------------------------------
	file, err := os.Open("post.csv") // 打开一个文件句柄？
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file) // 创建一个读取器？
	reader.FieldsPerRecord = -1 // 指定每条记录的需要读取的字段数量，负数没有限制。
	record, err := reader.ReadAll() // 返回由一系列切片组成的切片
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0) // 第一个元素
		post := Post{ID: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].ID)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
	//---------------------------------------------------------------
}