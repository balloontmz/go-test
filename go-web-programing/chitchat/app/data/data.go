package data

import (
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql 驱动
)

var Db *sql.DB

// 只要该包被引入了，该函数就会被调用
func init()  {
	var err error
	// 创建连接池
	Db, err = sql.Open("mysql", "root:ls950322@tcp(127.0.0.1:3306)/chitchat")
	if err != nil {
		log.Fatal(err)
	}
	return
}


