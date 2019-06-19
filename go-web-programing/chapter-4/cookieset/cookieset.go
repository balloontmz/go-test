package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request){
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programing",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Manning Publication Co",
		HttpOnly: true,
	}
	// 两种添加 cookie 的方法。http 的传指针，RW 的传值。
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	// w.Header().Set("Set-Cookie", c1.String()) // 设置首部
	// w.Header().Add("Set-Cookie", c2.String())// 添加首部
}

func getCookie(w http.ResponseWriter, r *http.Request)  {
	/*
	Header["Cookie"] 返回一个切片，和 Req 的 Cookies 方法返回的一致。
	Req 的 Cookie 返回指定的 Cookie，如果不存在，会报错。
	*/
	h :=r.Header["Cookie"]
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, "------------------\n")
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/set-cookie", setCookie)
	http.HandleFunc("/get-cookie", getCookie)
	server.ListenAndServe()
}