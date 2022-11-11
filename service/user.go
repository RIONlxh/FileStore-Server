package service

import (
	"FileStore-Server/db/mysql"
	"fmt"
)

func SignUp(username string, password string) error {
	sql := "insert into user (`user_name`, `user_pwd`) values (?, ?)"
	// 创建连接
	stmt, err := mysql.DBConn().Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 执行sql
	result, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 获取结果
	if row, err := result.RowsAffected(); row > 0 && err == nil {
		return nil
	}
	return err
}
