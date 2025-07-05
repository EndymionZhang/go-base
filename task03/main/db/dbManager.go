package db

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:13306)/go_gorm?charset=utf8mb4&parseTime=True&loc=Local"

func GetGormDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CloseGorm(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		return
	}
}

func GetSqlxDb() *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CloseSqlx(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err.Error())
	}
}
