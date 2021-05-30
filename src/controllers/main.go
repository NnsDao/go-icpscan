package controllers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	connection, err := gorm.Open(mysql.Open("root:dai2310ll@*@tcp(cdb-m0nw6jnk.bj.tencentcdb.com:10024)/icpscan?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = connection
}
