package controller

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	// threads, err := data.Threads()
	var err error
	threads, err := "  ", nil
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
