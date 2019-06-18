package main 

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type (
	Post struct {
		User string
		Threads []string
	}
)

func writeExample(w http.ResponseWriter, r *http.Request)  {
	str := `
	<html>
		<head><title>Go</title></head>
		<body><h1>Hello</h1></body>
	</html>
	`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request){ // curl -i 127.0.0.1:8000/redirect
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)//在执行完毕后就不再允许对首部进行写入。所以需要先写入首部再写入状态码。
}

func jsonExample(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "tomtiddler",
		Threads: []string{"first", "second", "third"},
	}
	data, _ := json.Marshal(post)
	w.Write(data)
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)//curl -i 127.0.0.1:8000/writeheader
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}