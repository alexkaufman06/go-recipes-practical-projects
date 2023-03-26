package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
)

var totalSize uint64

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	size := atomic.AddUint64(&totalSize, uint64(len(data)))

	// TODO: Work with data

	fmt.Fprintf(w, "size=%d\n", size)
	fmt.Printf("Total Size:%d\n", totalSize)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// curl --header "Content-Type: application/json" -d "{\"testing\":\"bark\"}" http://localhost:8080/upload
// curl --header "Content-Type: application/json" -d "{\"testing\":\"bark\", \"cat\":\"meow, meow\"}" http://localhost:8080/upload
