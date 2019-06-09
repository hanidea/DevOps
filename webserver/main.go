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
	http.HandleFunc("/sayHello",sayHello)
	http.ListenAndServe(":8000",nil)
}