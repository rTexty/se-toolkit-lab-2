package main

import (
	"encoding/json"
	"net/http"
)

func handleOutcomes(w http.ResponseWriter, r *http.Request) {
	data,err := loadOutcomes()
	if err != nil{
		http.Error(w, "error while loading data", 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)

}