package ge

import (
	"time"
)

// Check (err error) equal `if err != nil { panic(err) }`
func Check (err error) {
	if err != nil { panic(err) }
}

func Int(i int, err error) int {
	Check(err)
	return i
}
func IntList(i []int, err error) []int {
	Check(err)
	return i
}
func Int32List(i []int32, err error) []int32 {
	Check(err)
	return i
}

func Float64(i float64, err error) float64 {
	Check(err)
	return i
}
func Float64List(i []float64, err error) []float64 {
	Check(err)
	return i
}
func Float32(i float32, err error) float32 {
	Check(err)
	return i
}
func Float32List(i []float32, err error) []float32 {
	Check(err)
	return i
}

func String(s string, err error) string {
	Check(err)
	return s
}
func StringList(s []string, err error) []string {
	Check(err)
	return s
}

func Bool(b bool, err error) bool {
	Check(err)
	return b
}
func BoolList(b []bool, err error) []bool {
	Check(err)
	return b
}

func Any(v interface{}, err error) interface{} {
	Check(err)
	return v
}
func AnyList(v []interface{}, err error) []interface{} {
	Check(err)
	return v
}
func Time(v time.Time, err error) time.Time {
	Check(err)
	return v
}
func Func(closeFunc func () error) {
	Check(closeFunc())
}