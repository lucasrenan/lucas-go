package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	db     = map[string]string{}
	dbLock sync.RWMutex
	nextID int
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

// CRUD
func newHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		URL string `json:"url"`
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&request); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	dbLock.Lock()
	defer dbLock.Unlock()
	key := fmt.Sprintf("%x", nextID)
	db[key] = request.URL
	nextID++

	reply := map[string]string{
		"id":  key,
		"url": request.URL,
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(reply)
}

func main() {
	http.HandleFunc("/_/ping", pingHandler)
	http.HandleFunc("/v1/new", newHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
