package main

import (
	//"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	// io.WriteString(w,"<h1> Hello, James devops 0609</h1>")
	w.Write([]byte("hello world"))
}

func main() {
	//1.实现读取文件handler
	fileHandler := http.FileServer(http.Dir("./video"))
	http.Handle("/video/",http.StripPrefix("/video/", fileHandler))
	http.HandleFunc("/sayHello",sayHello)
	http.ListenAndServe(":8000",nil)
}