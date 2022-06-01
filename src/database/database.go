package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseHandler() *gorm.DB {
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:4306)"
	DBNAME := "hs_reservation"
	OPTION := "parseTime=true&loc=Local"

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

func NewTestDatabaseHandler() *gorm.DB {
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:4306)"
	DBNAME := "hs_reservation_test"
	OPTION := "parseTime=true&loc=Local"

	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
