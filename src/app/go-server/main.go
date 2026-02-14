package main

import "net/http"

func main() {
	http.HandleFunc("/outcomes", handleOutcomes)
	http.ListenAndServe(":8080", nil)
}