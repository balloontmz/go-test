package controller

import (
	"net/http"
	"time"

	"chitchat/app/data"
)

type (
	User struct {
		Password string
	}
	Session struct {
		Id       int
		Uuid     string
		Email    string
		UserID   int
		CreateAt time.Time
	}
)

func Auth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email")) // 根据 email 获取用户
	passwdHash := data.Encrypt(r.PostFormValue("password"))
	if user.Password == passwdHash {
		// ...
		// session := user.CreateSession()
		session := Session{Uuid: "123456"}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true, // 只能通过 http 和 https 访问，无法通过非 http 方法访问 -- 不懂
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
