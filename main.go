package main

import (
	"database/sql"
	"fmt"
	"go-books-list/controllers"
	"go-books-list/driver"
	"go-books-list/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8000"
	}

	fmt.Println("Server has strated on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
