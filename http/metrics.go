package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// DB is a database
type DB struct{}

// Add adds a metric to the database
func (db *DB) Add(m MetricA) string {
	return "7b91dbb8289845c291325cb295b2364b" // FIXME
}

var (
	db *DB
)

// MetricA is an application metric
type MetricA struct {
	Time   time.Time `json:"time"`
	Host   string    `json:"host"`
	CPU    float64   `json:"cpu"`    // CPU load
	Memory float64   `json:"memory"` // MB
}

func handleMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var m MetricA
	const maxSize = 1 << 20 // MB
	dec := json.NewDecoder(io.LimitReader(r.Body, maxSize))
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := db.Add(m)
	log.Printf("metric: %+v (id=%s)", m, id)

	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"id": id,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error reply: %s", err)
	}
}

func main() {
	http.HandleFunc("/metric", handleMetric)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}

	// curl --header "Content-Type: application/json" -d "{\"host\":\"GO\", \"cpu\": 5.0, \"memory\": 7.2}" http://localhost:8080/metric
}
