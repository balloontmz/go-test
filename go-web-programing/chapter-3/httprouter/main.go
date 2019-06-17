package main

import (
	"fmt"
	"net/http"

	// "golang.org/x/net/http2"
	"github.com/julienschmidt/httprouter"	//第三方多路复用器
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func main()  {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux, // 调用自定义的多路复用器
	}

	server.ListenAndServe()
}