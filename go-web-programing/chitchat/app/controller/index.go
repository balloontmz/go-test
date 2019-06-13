package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"chitchat/app/utility"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求首页")
	// threads, err := data.Threads()
	// if err == nil {

	// }

	_, err := utility.Session(w, r)
	threads, err := "  ", nil

	if err != nil { // 根据 err 判断是否登录
		generateHTML(w, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, threads, "layout", "public.navbar", "index") 
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string)  {
	var files []string
	for _, file := range fn {
		// templates 放在项目运行根目录下
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))// Must 函数能捕捉语法分析过程中的错误信息，ParseFiles 对模板进行语法分析
	templates.ExecuteTemplate(w, "layout", data)
}