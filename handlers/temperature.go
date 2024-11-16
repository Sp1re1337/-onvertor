package handlers

import (
	"fmt"
	"encoding/json"
	"linux/cache"
	"net/http"
)

type Temperature struct {
  Value    float64 `json:"value"`
	UnitFrom string  `json:"unit_from"`
	UnitTo   string  `json:"unit_to"`
	Result   float64 `json:"result"`
}

func GenerateCacheKey(temp Temperature) string {
	return fmt.Sprintf("%f:%s:%s", temp.Value, temp.UnitFrom, temp.UnitTo)
}

func ConvertTemperature(w http.ResponseWriter, r *http.Request) {
	var temp Temperature
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	cacheKey := GenerateCacheKey(temp)

	if cachedResult, found := cache.Get(cacheKey); found {
		temp.Result = cachedResult
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(temp)
		return
	}

	switch {
	case temp.UnitFrom == "C" && temp.UnitTo == "F":
		temp.Result = (temp.Value * 9 / 5) + 32
	case temp.UnitFrom == "F" && temp.UnitTo == "C":
		temp.Result = (temp.Value - 32) * 5 / 9
	default:
		http.Error(w, "Invalid units", http.StatusBadRequest)
		return
 }

 cache.Set(cacheKey, temp.Result)


 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(temp)
}
