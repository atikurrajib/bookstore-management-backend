package models

import (
	"github.com/atikurrajib/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

func (b *Book) UpdateBook(ID int64) *Book {
	var book Book
	db.Where("ID=?", ID).Find(&book)
	book.Name = b.Name
	book.Author = b.Author
	book.Publication = b.Publication
	db.Save(&book)
	return &book
}

func (b *Book) GetBookById(ID int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", ID).Find(&book)
	return &book, db
}

func (b *Book) GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
