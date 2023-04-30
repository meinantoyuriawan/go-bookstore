package config

import (
	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// var (
// 	db *gorm.DB
// )

// var (
// 	db *sql.DB
// )

// func Connect() {
// 	// d, err := gorm.Open("mysql", "meinanto:meinantoyuriawan@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
// 	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	db = d
// }

func Connect() (*sql.DB, error) {
	d, err := sql.Open("sqlite3", "data.db")

	if err != nil {
		panic(err)
	}
	return d, nil
}

// func GetDB() *sql.DB {
// 	return db
// }
