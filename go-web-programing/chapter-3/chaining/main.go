package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello!!!")
}

// 这个应该是中间件的底层实现了，这个是串联处理器函数--串联处理器的实现和其类似
func log(h http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}