package main

import (
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql 驱动
)

type (
	//Post 帖子结构体
	Post struct {
		ID int
		Content string
		Author string
	}
)

var Db *sql.DB

// 只要该包被引入了，该函数就会被调用
func init()  {
	var err error
	// sql.Open() 不建立到数据库的任何连接，也不验证驱动程序连接参数。相反，它只是准备数据库抽象以备后用。与底层数据存储的第一次实际连接将在第一次
	// 需要时被懒散地建立。如果您想立即检查数据库是否可用并可访问（例如，检查您是否可以建立网络连接并登录），请使用此操作db.Ping()并记住检查错误：
	// 创建连接池
	Db, err = sql.Open("mysql", "tomtiddler:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Posts(limit int) (posts []Post, err error) {
	// Query 方法返回一个 Rows 接口，Rows 是一个迭代器。调用 Next 方法获取下一行，迭代完毕返回 io.EOF
	rows, err := Db.Query("select postsId,content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next(){
		post := Post{}
		err = rows.Scan(&post.ID, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return 
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning postId"
	// 指向一个 sql.Stmt 接口的引用，位于 sql.Driver 包中。 该接口被数据库驱动实现
	stmt, err := Db.Prepare(statement) // 构建预处理语句
	if err != nil {
		return
	}
	defer stmt.Close()
	// QueryRaw 方法返回一个 sql.Raw 的引用。如果返回多个 Row，只会返回第一个 Raw，并且不会返回错误。所以能搭配 Raw 的方法使用。
	// 实现插入并且返回 ID 的功能。
	err = stmt.QueryRaw(post.Content, post.Author).Scan(&post.ID) // 执行插入语句并将返回值赋予给 post 的 ID
}

func GetPost(id int) (psot Post, err error) {
	post = Post{}
	// 也可以采用 sql.Stmt 实现，除了 stmt 能生成 sql 语句用于重复利用，两者差别不大
	err = Db.QueryRow("select postId, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return 
}

func (posst *Post) Update() (err error) {
	// 程序并不需要对接收者进行更新，也不需要对结果进行扫描，所以采用 Exec 函数
	// 只返回错误信息，为不为空都返回
	_, err = Db.Exec("update posts set content = $2, author = $3 where postId=$1", post.ID, post.Content, post.Author)
	return
}

func (posst *Post) Delete() (err error) {
	// 程序并不需要对接收者进行删除，也不需要对结果进行扫描，所以采用 Exec 函数
	// 只返回错误信息，为不为空都返回
	_, err = Db.Exec("delete from posts where postId=$1", post.ID)
	return
}