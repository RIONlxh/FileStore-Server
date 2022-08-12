package db

import (
	mysql "FileStore-Server/db/mysql"
	"database/sql"
	"fmt"
)

type FileInfo struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullStringgit
}

func FileInfoDB(filesize int64, fileaddr string, filesha1 string, filename string) bool {
	stmt, err := mysql.DBConn().Prepare("insert into file_info(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("database prepare failed! ")
		return false
	}
	ret, err := stmt.Exec(filesha1, filename, filesize, fileaddr)
	if err != nil {
		fmt.Printf("sql exect failed! ,Err : %s", err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File Sha1 :%s has exists! ", filesha1)
			return false
		}
		return true
	}
	return false
}
