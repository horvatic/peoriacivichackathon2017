package main

import (
	"json"
	"net/http"
)

func main() {
	http.HandleFunc("/events", json.EventData)
	http.ListenAndServe(":8080", nil)
}
