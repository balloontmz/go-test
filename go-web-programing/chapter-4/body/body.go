package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength // 主体数据的字节长度
	body := make([]byte, len)// 根据长度创建字节数组
	r.Body.Read(body) // 将主体读入字节数组 
	fmt.Fprintln(w, string(body))
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	/*
	curl -id "first_name=tom&last_name=tiddler" 127.0.0.1:8080/body
	*/
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}