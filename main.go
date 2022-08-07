package main

import (
	"FileStore-Server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/uploadfile", handler.UploadFile)
	http.HandleFunc("/api/get_fileinfo_one", handler.GetFileInfoOne)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Connection failed")
	}
}
