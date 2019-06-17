/*
package main

import (
    "fmt"
    "net/http"
)

type HelloHandler struct{}
func (h HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Handler!")
}

func hello (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    helloHandler := HelloHandler{}
    http.Handle("/hello1", helloHandler)
    http.HandleFunc("/hello2", hello)
    server.ListenAndServe()
}

[go 简单的路由定向](https://www.jianshu.com/p/3b5c4fc0695c)
*/
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {}

func (h *HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct {}

func (h *WorldHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	 fmt.Fprintf(w, "World!")
}

func main ()  {
	hello := HelloHandler{}
	world := WorldHandler{}

	// 使用服务器默认的 DefaultServeMux
	server := http.Server {
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}