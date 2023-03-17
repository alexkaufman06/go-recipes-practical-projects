package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// isbn -> book
var booksDB = map[string]Book{
	"0062225677": {
		Title:  "The Colour of Magic",
		Author: "Terry Pratchett",
		ISBN:   "0062225677",
	},
	"0765394855": {
		Title:  "Old Man's War",
		Author: "John Scalzi",
		ISBN:   "0765394855",
	},
}

func getBook(isbn string) (Book, error) {
	book, ok := booksDB[isbn]
	if !ok {
		return Book{}, fmt.Errorf("unknown ISBN: %q", isbn)
	}

	return book, nil
}

// Book is information about book
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func handleGetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]

	book, err := getBook(isbn)
	if err != nil {
		log.Printf("error - get: unknown ISBN - %q", isbn)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.Printf("error - json: %s", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{isbn}", handleGetBook).Methods("GET")

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// curl http://localhost:8080/books/0062225677
}
