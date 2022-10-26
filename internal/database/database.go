package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConn() *gorm.DB {
	//dsn := "host=localhost user=root password=password dbname=course port=3306"
	dsn := "root:password@tcp(127.0.0.1:3306)/course?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("can't connect db")
	}
	return db
}
