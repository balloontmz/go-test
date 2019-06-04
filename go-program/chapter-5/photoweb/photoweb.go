package main

import (
	"io"
	"os"
	"net/http"
	"log"
	"html/template"
	"io/ioutil"
	"path"
)

const (
	ListDir = 0x001
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates map[string]*template.Template

func init()  {
	// templates := make(map[string]*template.Template)
	
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

func renderHTML(w http.ResponseWriter, tmpl string, locals map[string]interface{})  {
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		renderHTML(w, "upload", nil)
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		t, err := ioutil.TempFile(UPLOAD_DIR, filename)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request)  {
	imageID := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageID
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request)  {
	
}

// 处理器错误处理
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		defer func ()  {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// ...
			}
		}()
		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flag int)  {
	
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", mux)
	
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}