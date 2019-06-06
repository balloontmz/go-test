package controller

import (
	"net/http"
)

type (
	User struct {
		Password string
	}
)

func Auth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// user, _ := data.UserByEmail(r.PostFormValue("email")) // 根据 email 获取用户
	user := User{"123456"}
	passwdHash := "123456"
	if user.Password == passwdHash {
		// ...

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
