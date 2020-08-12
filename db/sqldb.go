package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//Db 是mysql的句柄
	Db  *gorm.DB
	err error
)

//init 获得链接数据库的句柄
func init() {
	Db, err = gorm.Open("mysql", "db_write:tm19941015@tcp(47.56.89.150:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
	//defer Db.Close()
	if err != nil {
		panic(err.Error())
	}

}
