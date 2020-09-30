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
)

var DB *sql.DB

func init(){
	var dsn string
	dsn = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	/* 连接数据库 */
	database, err := sql.Open("mysql", dsn)
	fmt.Println(database)
	if err != nil {
		fmt.Println("Open mysql failed,err:%v\n", err)
		return
	}
	DB = database
}
