package utils

import "strconv"

func ConvertToUint(s string) uint {
	value, _ := strconv.ParseUint(s, 10, 64)
	return uint(value)
}

func ConvertToFloat64(s string) float64 {
	value, _ := strconv.ParseFloat(s, 64)
	return value
}
