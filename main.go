package main

import (
	"FileStore-Server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/uploadfile", handler.UploadFile)
	http.HandleFunc("/api/get_fileinfo_one", handler.GetFileInfoOne)
	http.HandleFunc("/api/downloadfile", handler.DownloadFile)
	http.HandleFunc("/api/rename_file", handler.RenameFile)
	http.HandleFunc("/api/delete_file", handler.DeleteFile)

	http.HandleFunc("/api/user/signup", handler.Sginup)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Connection failed")
	}
}
