package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"<h1> Hello, James 5</h1>")
}

func main() {
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":5000",nil)
}