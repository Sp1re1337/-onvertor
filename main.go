package main

import (
	"log"
	"net/http"

	"linux/handlers"
)

func main() {
	http.HandleFunc("/convert", handlers.ConvertTemperature)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}