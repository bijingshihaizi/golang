package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"syscall"
)

const(
	USERNAME = ""
	PASSWORD = ""
	NETWORK = ""
	SERVER = ""
	PORT = 3306
	DATABASE = ""

	USERNAME1 = ""
	PASSWORD1 = ""
	SERVER1 = ""
	DATABASE1 = ""
)

var DB *sql.DB

func init(){
	var dsn string
	v, _ := syscall.Getenv("ENV_NAME")
	if v == "prod" {
		dsn = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME1, PASSWORD1, NETWORK, SERVER1, PORT, DATABASE1)
	}else{
		dsn = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	}
	/* 连接数据库 */
	database, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Open mysql failed,err:%v\n", err)
		return
	}
	DB = database
}
