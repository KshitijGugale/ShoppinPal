package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hope page")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/book/{author}/{title}/{isbn}/{release_date}", AddBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

type Book struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	ISBN        string `json:"isbn"`
	ReleaseDate string `json:"release_date"`
}

type Books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := Books{
		Book{
			Author:      "A1",
			Title:       "Golang Book",
			ISBN:        "ABCD",
			ReleaseDate: "01-01-2020",
		},
	}
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(mux.Vars(r))
	vars := mux.Vars(r)
	author := vars["author"]
	title := vars["title"]
	isbn := vars["isbn"]
	releaseDate := vars["release_date"]

	book := Book{
		Author: author,
		Title: title,
		ISBN: isbn,
		ReleaseDate: releaseDate,
	}
	fmt.Println(book)

	// TODO: ADD logic to append book to Books
}

func main() {
	handleRequest()
}
