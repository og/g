package gconv

import (
	"strconv"
)

func Float64String(f float64) string {
	return strconv.FormatFloat(f,'f', -1, 64)
}
func Float32String (f float32) string {
	return strconv.FormatFloat(float64(f),'f',-1,32)
}
