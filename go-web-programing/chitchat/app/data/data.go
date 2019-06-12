package data

import (
	"log"
	"database/sql"
)

var Db *sql.DB

// 只要该包被引入了，该函数就会被调用
func init()  {
	var err error
	Db, err := sql.Open("mysql", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}


