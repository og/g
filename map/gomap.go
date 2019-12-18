package gmap

import (
	"reflect"
	"sort"
)


type KeyList struct {
	value []reflect.Value
}
func (kList KeyList) String() (keys []string) {
	for _, key := range kList.value {
		keys = append(keys, key.String())
	}
	sort.Strings(keys)
	return
}
func (kList KeyList) Int() (keys []int) {
	for _, key := range kList.value {
		keys = append(keys, int(key.Int()))
	}
	sort.Ints(keys)
	return
}
func (kList KeyList) Float64() (keys []float64) {
	for _, key := range kList.value {
		keys = append(keys, float64(key.Float()))
	}
	sort.Float64s(keys)
	return
}

func Keys(data interface{}) (keys KeyList){
	mapKeys := reflect.ValueOf(data).MapKeys()
	for _, key := range mapKeys {
		keys.value = append(keys.value, key)
	}
	return
}
// Map equal `map[string]interface{}`
type Map map[string]interface{}

