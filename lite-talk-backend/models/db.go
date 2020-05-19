package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open("mysql", "root:mysql666@(127.0.0.1)/lite_talk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("MySql Connect Error: %v", err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}
