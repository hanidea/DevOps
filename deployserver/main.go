package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"<h1> Hello, James 3000</h1>")
}

func main() {
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":3000",nil)
}