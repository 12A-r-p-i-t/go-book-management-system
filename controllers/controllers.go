package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/12A-r-p-i-t/bookManagementSystem/models"
	"github.com/gorilla/mux"
)

var b *models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "GET")
	books := b.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func GetABook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "GET")
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Fatal("Failed to convert to int", err)
	}
	book := b.GetABook(uint(id))
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "POST")
	json.NewDecoder(r.Body).Decode(&b)
	id := b.CreateABook()
	json.NewEncoder(w).Encode(id)
}

func UpdateABook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "PUT")
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Fatal("Failed to convert to int", err)
	}
	var change models.Book
	json.NewDecoder(r.Body).Decode(&change)
	fmt.Println(change)
	temp := b.UpdateABook(uint(id), change)
	json.NewEncoder(w).Encode(temp)
}

func DeleteABook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Fatal("Failed to convert to int", err)
	}
	temp := b.DeleteABook(uint(id))
	json.NewEncoder(w).Encode(temp)
}
