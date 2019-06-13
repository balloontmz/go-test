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
	// sql.Open() 不建立到数据库的任何连接，也不验证驱动程序连接参数。相反，它只是准备数据库抽象以备后用。与底层数据存储的第一次实际连接将在第一次
	// 需要时被懒散地建立。如果您想立即检查数据库是否可用并可访问（例如，检查您是否可以建立网络连接并登录），请使用此操作db.Ping()并记住检查错误：
	// 创建连接池
	Db, err = sql.Open("mysql", "root:ls950322@tcp(127.0.0.1:3306)/chitchat")
	if err != nil {
		log.Fatal(err)
	}
	return
}


