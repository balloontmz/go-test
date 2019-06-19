// 采用 cookie 闪现消息
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMsg(w http.ResponseWriter, r *http.Request)  {
	msg := []byte("Hello World")
	// 设置临时 cookie，跳转即失效
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMsg(w http.ResponseWriter, r *http.Request)  {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie{
			fmt.Fprintln(w, "No message found")
		}
	} else {
		// 采用过期的持久 cookie 覆盖，然后自动删除 cookie 。--过去的时间。
		rc := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/set-msg", setMsg)
	http.HandleFunc("/show-msg", showMsg)
	server.ListenAndServe()
}