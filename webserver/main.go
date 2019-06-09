package main

import (
	"io"
	"net/http"
	"strings"
	"crypto/md5"
	"time"
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	// io.WriteString(w,"<h1> Hello, James devops 0609</h1>")
	w.Write([]byte("hello world"))
}

func main() {
	//1.实现读取文件handler
	fileHandler := http.FileServer(http.Dir("./video"))
	http.Handle("/video/",http.StripPrefix("/video/", fileHandler))
	//2. 上传文件handler
	http.HandleFunc("/api/upload",uploadHandler)
	//3. 显示list
	http.HandleFunc("/api/list",getFileListHandler)

	http.HandleFunc("/sayHello",sayHello)
	http.ListenAndServe(":8000",nil)
}

//业务逻辑
func uploadHandler(w http.ResponseWriter, r *http.Request){
	//限制客户端上传视频文件大小
	r.Body= http.MaxBytesReader(w, r.Body, 10*1024*1024)
	err := r.ParseMultipartForm(10*1024*1024)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//获取上传的文件
	file, fileHeader, err := r.FormFile("uploadFile")

	//检查文件类型
	ret	:= strings.HasSuffix(fileHeader.Filename,".mp4")
	if ret == false {
		http.Error(w,"not mp4",http.StatusInternalServerError)
		return
	}

	//获取随机名称
	md5Byte := md5.Sum([]byte(fileHeader.Filename + time.Now().String()))
	md5Str := fmt.Sprintf("%x",md5Byte)
	newFileName := md5Str + ".mp4"

	// 写入文件
	dst, err := os.Create("./video/"+newFileName)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	defer file.Close()
	if _, err := io.Copy(dst, file); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

//获取视频文件列表
func getFileListHandler(w http.ResponseWriter, r *http.Request){
	files, _ := filepath.Glob("video/*")
	var ret []string
	for _, file :=range files {
		ret = append(ret, "http://" + r.Host + filepath.Base(file))
	}
	retJson, _ := json.Marshal(ret)
	w.Write(retJson)
	return
}
