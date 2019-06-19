package main

import (
	"fmt"
)

type (
	//Post 帖子结构体
	Post struct {
		ID int
		Content string
		Author string
	}
)

//PostByID 根据帖子 id 存储帖子指针
var PostByID map[int] *Post
//PostByAuthor 根据帖子 author 存储帖子指针切片
var PostByAuthor map[string] []*Post

func store(post Post)  {
	PostByID[post.ID] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main(){
	PostByID = make(map[int] *Post)//初始化 PostByID 的值
	PostByAuthor = make(map[string] []*Post)

	post1 := Post{ID:1, Content:"test", Author: "tom"}
	post2 := Post{ID:2, Content:"test", Author: "tiddler"}
	post3 := Post{ID:3, Content:"test", Author: "tom"}
	post4:= Post{ID:4, Content:"test", Author: "balloon"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	//ID 存储的是结构体
	fmt.Println(PostByID[1])
	fmt.Println(PostByID[2])

	//Author 存储的是切片。post 变量的生命周期是循环内。
	for _, post := range PostByAuthor["tom"] {
		fmt.Println(post)
	}

	for _, post := range PostByAuthor["balloon"] {
		fmt.Println(post)
	}
}