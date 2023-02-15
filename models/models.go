package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var dbConnection *gorm.DB

func init() {
	fmt.Println("Starting MYSQL DB...")
	db, err := gorm.Open("mysql", "root:Arpit123@tcp(127.0.0.1:3306)/Books?charset=utf8&parseTime=True")
	dbConnection = db
	db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal("Failed to start the DB in models.go ", err)
	}
	fmt.Println("Succesfully connected to DB")
}

func (b *Book) GetAllBooks() []Book {
	var books []Book
	err := dbConnection.Find(&books)
	if err.Error != nil {
		log.Fatal("Failed to retreive books from database")
	}
	fmt.Println("Successfully retreived books from Database")
	return books
}

func (b *Book) GetABook(id uint) Book {
	var book Book
	err := dbConnection.First(&book, id)
	if err.Error != nil {
		log.Fatal(err, "Failed to retreive the book with given id : ", id)
	}
	fmt.Println("Successfully retreived the book with given id :", id)
	return book
}

func (b *Book) CreateABook() uint {
	err := dbConnection.Create(&b)
	if err.Error != nil {
		log.Fatal(err, "Failed to create a book")
	}
	fmt.Println("Successfully created a book with given id :", b.ID)
	return b.ID
}

func (b *Book) UpdateABook(id uint, updatedValues Book) uint {
	err := dbConnection.Model(&b).Where("id = ?", id).Update(updatedValues)
	if err.Error != nil {
		log.Fatal(err, "Failed to update the book")
	}
	fmt.Println("Successfully updated the book with given id :", id)
	return id
}

func (b *Book) DeleteABook(id uint) uint {
	err := dbConnection.Delete(&b, id)
	if err.Error != nil {
		log.Fatal(err, "Failed to delete a book")
	}
	fmt.Println("Successfully deleted a book with given id :", b.ID)
	return b.ID
}
