package main

import (
	"fmt"
	"net/http"

	"chitchat/app/controller"
)



func main()  {
	
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", controller.Index)
	// 以下暂未实现
	// mux.HandleFunc("/err", controller.Err)

	// mux.HandleFunc("/login", controller.Login)
	// mux.HandleFunc("/logout", controller.Logout)
	// mux.HandleFunc("/signup", controller.SignUp)
	// mux.HandleFunc("/signup_account", controller.SignUpAcc)
	mux.HandleFunc("/authenticate", controller.Auth)

	// mux.HandleFunc("/thread/new", controller.New)
	// mux.HandleFunc("/thread/create", controller.Create)
	// mux.HandleFunc("/thread/post", controller.Post)
	// mux.HandleFunc("/thread/read", controller.Read)

	server := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	fmt.Println("项目已启动")
	server.ListenAndServe()
}