package handler

import (
	"FileStore-Server/db"
	"FileStore-Server/model"
	"FileStore-Server/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		contByte, err := ioutil.ReadFile("./static/html/index.html")
		if err != nil {
			fmt.Println(err)
			io.WriteString(w, "Server Reponse Failed")
			return
		}
		io.WriteString(w, string(contByte))
	}
	if r.Method == "POST" {
		// 1.接收文件流
		fileStream, fileHandler, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("File Upload Failed! Error: %s", err.Error())
			return
		}
		defer fileStream.Close()

		// 1.2 model 初始化赋值
		fileModel := model.FileInfoModel{
			FileName: fileHandler.Filename,
			FilePath: "./tmp/" + fileHandler.Filename,
			FileAt:   time.Now().Format("2006-01-02 15:04:05"),
		}

		// 2.创建文件夹
		newFile, err := os.Create(fileModel.FilePath) // 注意此处的相对路径，Create() 只能创建文件，不能创建文件夹
		if err != nil {
			fmt.Printf("Dir Creating Failed! Error: %s \n", err.Error())
			return
		}

		defer newFile.Close()

		// 3.拷贝文件
		fileModel.FileSize, err = io.Copy(newFile, fileStream)
		if err != nil {
			fmt.Printf("Copy File Failed! Error: %s \n", err.Error())
			return
		}

		// 3.1 计算sha1值，并更新file map
		newFile.Seek(0, 0)
		fileModel.FileSha1 = util.FileSha1(newFile)
		newFileModel := model.UpdateFileInfo(fileModel)

		byteData, err := json.Marshal(newFileModel)
		ret := db.FileInfoDB(newFileModel.FileSize, newFileModel.FilePath, newFileModel.FileSha1, newFileModel.FileName)
		if ret == false {
			io.WriteString(w, "File Update Failed")
		}
		// 4.返回文件流
		//io.WriteString(w, "Upload Success!")
		w.Write(byteData)

	}
}

func GetFileInfoOne(w http.ResponseWriter, r *http.Request) {
	// 解析params参数
	r.ParseForm()
	fileSha1 := r.Form["filesha1"][0]

	// 获取 FileInfoModel struct
	fmi := model.GetFileInfo(fileSha1)

	// 解析成json数据
	jsonByte, err := json.Marshal(fmi)
	if err != nil {
		fmt.Printf("Json Failed! %s", err.Error())
		return
	}

	// 写入response中
	w.Write(jsonByte)

}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	// 1.获取入参
	r.ParseForm()
	fileHash := r.Form.Get("filehash")
	fileModel := model.GetFileInfo(fileHash)

	// 2.查找文件是否存在
	file, err := os.Open(fileModel.FilePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 3.读取文件
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 4.设置响应头
	w.Header().Set("Content-Type", "application/octect")
	w.Header().Set("content-disposition", "attachment; filename=\""+fileModel.FileName+"\"")
	w.Write(fileBytes)
}

func RenameFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileModel := model.GetFileInfo(fileHash)
	curFileModel.FileName = newFileName
	model.UpdateFileInfo(curFileModel)

	// 并未修改存储的文件名
	byteData, err := json.Marshal(curFileModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(byteData)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		filehash := r.Form.Get("filehash") // 只能解析 params传参，无法解析 from-data
		fmt.Println(filehash)
		fileModel := model.GetFileInfo(filehash)
		os.Remove(fileModel.FilePath)
		model.DeleteFileInfo(filehash)
		io.WriteString(w, "delete success")
	}
}
