package database

import (
	"fmt"
	u "github.com/user/utils"
	"gorm.io/gorm"
	gorm_mysql "gorm.io/driver/mysql"
)


var DB *gorm.DB

func init(){
	dsn,err := u.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase fail", err)
	}

    db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{});
    if err != nil {
        panic(err)
        return
    }
    DB = db
	/* 连接数据库 */
	/*
	database, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Open mysql failed,err:%v\n", err)
        return
    }
	DB = database
    */
}
