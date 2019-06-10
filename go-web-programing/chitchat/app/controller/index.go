package controller

import (
	"net/http"
	"text/template"

	"chitchat/app/utility"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// threads, err := data.Threads()
	// if err == nil {

	// }

	_, err := utility.Session(w, r)

	// templates 放在项目运行根目录下
	public_tmpl_files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}

	private_tmpl_files := []string{
		"templates/layout.html",
		"templates/private.navbar.html",
		"templates/index.html",
	}

	var templates *template.Template

	if err != nil { // 根据 err 判断是否登录
		templates = template.Must(template.ParseFiles(public_tmpl_files...))
	} else {
		templates = template.Must(template.ParseFiles(private_tmpl_files...)) // Must 函数能捕捉语法分析过程中的错误信息，ParseFiles 对模板进行语法分析
	}

	// threads, err := data.Threads()
	threads, err := "  ", nil
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
