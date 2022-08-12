package mysql

import (
	"FileStore-Server/conf"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	MysqlSource := conf.MYSQLSOURCE
	db, _ = sql.Open("mysql", MysqlSource)
	db.SetMaxOpenConns(50)
	err := db.Ping()
	if err != nil {
		fmt.Printf("Failed to connection mysql,err : %s", err.Error())
		return
	}
}

func DBConn() *sql.DB {
	return db
}
