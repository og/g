package gmath

import (
	"fmt"
	"strconv"
)

// Calculate int/int percent , if total test 0, return 0
func IntPercent (part int, total int) int {
	// percent no need division zero
	if total == 0 { return 0 }
	return int(Float64ToFixed(float64(part) / float64(total), 2) * 100)
}
// Calculate float64/float64 percent , if total test 0, return float64(0)
// return 0 ~ 100
func Float64Percent (part float64, total float64) int {
	// percent no need division zero
	if total == 0 { return 0 }
	return int(Float64ToFixed(part / total, 2) * 100)
}
// float64 keep the decimal point
func  Float64ToFixed(f float64, digit int) float64 {
	format := "%." + strconv.Itoa(digit) + "f"
	output, err := strconv.ParseFloat(fmt.Sprintf(format, f), 64); if err != nil {panic(err)}
	return output
}
// if b == 0 {return 0} ; return a/b
func Division(a float64 , b float64) float64 {
	if b == 0 {return 0} ; return a/b
}
