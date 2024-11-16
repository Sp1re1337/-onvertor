package utils

import (
	"fmt"
	"linux/handlers"
)

func GenerateCacheKey(temp handlers.Temperature) string {
	return fmt.Sprintf("%f:%s:%s", temp.Value, temp.UnitFrom, temp.UnitTo)
}