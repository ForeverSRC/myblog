package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const(
	DB_LOG="root:src201261730@(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	DRIVER="mysql"
)

var Db *gorm.DB


func RegisterDB(){
	var err error
	Db,err=gorm.Open(DRIVER,DB_LOG)
	if err!=nil{
		log.Panic("数据库打开错误")
	}
	Db.AutoMigrate(&Category{},&Topic{},&Comment{},&Todo{},&Album{},&Photo{})
}

