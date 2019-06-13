package data

import (
	"time"
)

type (
	Thread struct {
		ID int
		Uuid string
		Topic string
		UserID int
		CreatedAt time.Time
	}
)

func Threads() (threads []Thread, err error) {
	// 通过数据库连接池与数据库进行连接，并向数据库发送一个 SQL 查询
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	// 遍历行，将数据放入结构体
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.ID, &th.Uuid, &th.Topic, &th.UserID, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

//NumReplies 获取回复数，模版中的 Thread 对象能直接调用该对象的方法。
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where  thread_id = $1", thread.ID)
	if err != nil {
		return
	}
	for rows.Next(){
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}