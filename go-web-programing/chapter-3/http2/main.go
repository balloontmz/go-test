package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

type MyHandler struct {}

// 结构和指针 和 其方法的区别，还需要再看文档。忘记了，好像在调用时能相互转换？ 
func (h *MyHandler) ServerHTTP (w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World!")
}

func main()  {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}

	http2.ConfigureServer(&server, &http2.Server{}) // 此行是重点
	
	server.ListenAndServeTLS("cert.pem", "key.pem")
}