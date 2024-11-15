package main

import (
	"log"
	"encoding/json"
	"net/http"
)

type Temperature struct {
	Value    float64 `json:"value"`
	UnitFrom string `json:"unit_from"`
	UnitTo   string `json:"unit_to"`
	Result   float64 `json:"rusult"`
}

func main() {
  http.HandleFunc("/convert", convertTemperature)
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func convertTemperature(w http.ResponseWriter, r *http.Request) {
	var temp Temperature
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
	  http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	switch {
	case temp.UnitFrom == "C" && temp.UnitTo == "F":
		temp.Result = (temp.Value * 9 / 5) + 32
	case temp.UnitFrom == "F" && temp.UnitTo == "C":
		temp.Result = (temp.Value - 32) * 5 / 9
	default:
		http.Error(w, "invalid units", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temp)
}