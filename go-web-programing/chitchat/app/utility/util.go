package utility

import (
	"errors"
	"net/http"

	"chitchat/app/data"
)

func Session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie") // 从请求中获取 cookie，如果不存在，直接返回
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok { // 检查数据库 session 是否存在
			err = errors.New("非法会话")
		}
	}
	return
}
