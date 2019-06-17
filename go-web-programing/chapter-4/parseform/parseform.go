package main

import (
	"fmt"
	"net/http"
)

/*
!!!
Form  URL
PostForm  Body
MultipartForm  Both
*/

func progress(w http.ResponseWriter, r *http.Request){
	r.ParseForm() // 调用方法对请求进行语法分析
	fmt.Fprintln(w, r.Form) // 根据调用的方法访问对应属性
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/progress", progress)
	server.ListenAndServe()
}