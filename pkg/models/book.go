package models

import (
	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	"database/sql"
	"fmt"

	// "math/rand"
	// "strconv"
	// "time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/meinantoyuriawan/go-bookstore/pkg/config"
)

// var db *sql.DB

type Book struct {
	// gorm.Model
	ID          int    `json:"id"`
	CreatedAt   string `json:"createdat"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	// db.AutoMigrate(&Book{})
// }

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	// db.Create(&b)
	// return b
	// insert query from book parameter

	Name := b.Name
	Author := b.Author
	Publication := b.Publication

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO bookdata (Name, Author, Publication) values (?, ?, ?)", Name, Author, Publication)
	if err != nil {
		panic(err)
	}
	fmt.Println("update success!")
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	// db.Find(&Books)
	// return Books

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM bookdata")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = Book{}
		var err = rows.Scan(&each.ID, &each.CreatedAt, &each.Name, &each.Author, &each.Publication)
		if err != nil {
			panic(err)
		}

		Books = append(Books, each)
	}
	return Books
}

// func GetBookById(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	db := db.Where("ID=?", Id).Find(&getBook)
// 	return &getBook, db
// }

func GetBookById(Id int64) (*Book, *sql.DB) {

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result = Book{}
	err = db.
		QueryRow("SELECT * FROM bookdata WHERE id = ?", Id).
		Scan(&result.ID, &result.CreatedAt, &result.Name, &result.Author, &result.Publication)
	if err != nil {
		panic(err)
	}
	// var getBook Book
	// db := db.Where("ID=?", Id).Find(&getBook)
	return &result, db
}

func DeleteBook(Id int64) Book {
	// 	var book Book
	// 	db.Where("ID=?", ID).Delete(book)
	// 	return book
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result = Book{}
	err = db.
		QueryRow("SELECT * FROM bookdata WHERE id = ?", Id).
		Scan(&result.ID, &result.CreatedAt, &result.Name, &result.Author, &result.Publication)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("delete from bookdata where id = ?", Id)
	if err != nil {
		panic(err)
	}
	return result
}

func (b *Book) UpdateBook(Id int64) *Book {
	// db.NewRecord(b)
	// db.Create(&b)
	// return b
	// insert query from book parameter

	Name := b.Name
	Author := b.Author
	Publication := b.Publication

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO bookdata (ID, Name, Author, Publication) values (?, ?, ?, ?)", Id, Name, Author, Publication)
	if err != nil {
		panic(err)
	}
	fmt.Println("update success!")
	return b
}
