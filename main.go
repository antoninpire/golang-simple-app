package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Example struct {
	ID   string `json:"id"`
	Data int    `json:"data"`
}

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		data := Example{
			ID:   "some_random_id",
			Data: 100,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
