package gconv

import (
	"strconv"
)

func StringInt(s string) (i int, err error) {
	return strconv.Atoi(s)
}
func StringInt64 (s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
func StringFloat64 (s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
func StringFloat32 (s string) (float32, error) {
	f64, err := strconv.ParseFloat(s, 32)
	return float32(f64), err
}