package gmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)
func ExampleStringKeys() {
	sMap := map[string]string{
		"name": "nimo",
		"title": "abc",
		"lang": "go",
	}

	for _, key := range UnsafeKeys(sMap).String() {
		log.Print(key+":"+sMap[key])
	}
	// lang:go
	// name:nimo
	// title:abc
}
func Test_StringKeys(t *testing.T) {
	// ExampleStringKeys()
	sMap := make(map[string]string)
	sMap["name"] = "nimo"
	sMap["title"] = "abc"
	assert.Equal(t, []string{"name", "title"}, StringStringKeys(sMap))
	sMap["lang"] = "go"
	assert.Equal(t, []string{"lang", "name", "title"},StringStringKeys(sMap))
	sMap["1"] = "1"
	assert.Equal(t, []string{"1", "lang", "name", "title"},StringStringKeys(sMap))
}


func ExampleIntKeys() {
	iMap := map[int]string{
		6: "nimo",
		2: "abc",
		9: "go",
	}
	for _, key := range IntStringKeys(iMap) {
		log.Print(key, ":", iMap[key])
	}
	// 2:abc
	// 6:nimo
	// 9:go
}
func Test_IntKeys(t *testing.T) {
	// ExampleIntKeys()
	sMap := make(map[int]string)
	sMap[1] = "nimo"
	sMap[3] = "abc"
	assert.Equal(t, []int{1, 3},IntStringKeys(sMap))
	sMap[2] = "go"
	assert.Equal(t, []int{1, 2, 3},IntStringKeys(sMap))
	sMap[8] = "1"
	assert.Equal(t, []int{1, 2, 3, 8},IntStringKeys(sMap))
}


func ExampleFloat64Keys() {
	fMap := map[float64]string{
		6.1: "nimo",
		2.2: "abc",
		9.3: "go",
	}

	for _, key := range Float64StringKeys(fMap) {
		log.Print(key, ":", fMap[key])
	}
	// 2.2:abc
	// 6.1:nimo
	// 9.3:go
}
func Test_Float64Keys(t *testing.T) {
	// ExampleFloat64Keys()
	sMap := make(map[float64]string)
	sMap[1.1] = "nimo"
	sMap[3.2] = "abc"
	assert.Equal(t, []float64{1.1, 3.2},Float64StringKeys(sMap))
	sMap[2.3] = "go"
	assert.Equal(t, []float64{1.1, 2.3, 3.2},Float64StringKeys(sMap))
	sMap[8.4] = "1"
	assert.Equal(t, []float64{1.1, 2.3, 3.2, 8.4},Float64StringKeys(sMap))
}
func TestMap(t *testing.T) {
	bjson, _ := json.Marshal(Any{"name":"nimo"})
	sjson := string(bjson)
	assert.Equal(t, `{"name":"nimo"}`,sjson)
}
