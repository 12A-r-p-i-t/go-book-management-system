package router

import (
	"github.com/12A-r-p-i-t/bookManagementSystem/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getBooks", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/getBooks/{id}", controllers.GetABook).Methods("GET")
	r.HandleFunc("/createBook", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/updateBook/{id}", controllers.UpdateABook).Methods("PUT")
	r.HandleFunc("/deleteBook/{id}", controllers.DeleteABook).Methods("DELETE")
	return r
}
