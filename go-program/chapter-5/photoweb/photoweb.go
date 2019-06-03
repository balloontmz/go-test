package main

import (
	"io"
	"net/http"
	"log"
	"html/template"
)

const (
	ListDir = 0x001
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./views"
)

func init()  {
	templates := make(map[string]*template.Template)
	
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		io.WriteString(w, `
		
		`)
		return 
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}