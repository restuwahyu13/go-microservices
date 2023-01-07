package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{"name": "Service Products"}
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":4002", nil)
}
