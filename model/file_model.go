package model

// 文件结构信息
type FileInfoModel struct {
	FileSha1 string
	FileName string
	FileSize int64
	FilePath string
	FileAt   string
}

// 开辟的临时存储空间
var fileInfoMap map[string]FileInfoModel

func init() {
	fileInfoMap = make(map[string]FileInfoModel)
}

// 新增/更新文件元信息
func UpdateFileInfo(fi FileInfoModel) {
	fileInfoMap[fi.FileSha1] = fi
}

// 获取文件信息
func GetFileInfo(fileSha1 string) FileInfoModel {
	return fileInfoMap[fileSha1]
}
