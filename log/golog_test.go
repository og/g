package l

import (
	"encoding/json"
	"testing"
)
func ExampleV_struct () {
	example := struct {
		String string
		Number int
		Float float64
	}{
		String: "abc",
		Number: 1,
		Float: 1.2422414,
	}
	V(example)
}
func ExampleV_sub_struct() {
	type Object struct {
		String string
		Number int
		Float float64
	}
	example := struct {
		String string
		Object Object
	}{
		String: "abc",
		Object: Object {
			String: "abc",
			Number: 1,
			Float: 1.2422414,
		},
	}
	V(example)
}
func ExampleV_struct_slice() {
	type Object struct {
		String string
		Number int
		Float float64
	}
	type Example struct {
		String string
		Array []Object
	}
	example := Example {
		String: "abc",
		Array: []Object{
			Object {
				String: "abc",
				Number: 1,
				Float: 1.2422414,
			},
			Object {
				String: "abc",
				Number: 1,
				Float: 1.2422414,
			},
		},
	}
	V(example)
}
func ExampleV_jsonString() {
	example := struct {
		String string
		JSON string
	}{
		String: "string",
		JSON: `{"name":"nimo","age": 18}`,
	}
	V(example)
}
func ExampleV_arg() {
	example := struct {
		String string
	}{
		String: "string",
	}
	V(example, "github.com/og/golog")
}
func ExampleV_chinese () {
	V("你好")
}
func ExampleV_JSON () {
	type User struct {
		Name string
		Age int
	}
	user := User{"nimo", 27,}
	bjson, _ := json.MarshalIndent(user, "", "  ")
	V(string(bjson))
}
func TestV(t *testing.T) {
	ExampleV_struct()
	ExampleV_sub_struct()
	ExampleV_struct_slice()
	ExampleV_jsonString()
	ExampleV_arg()
	ExampleV_chinese()
	ExampleV_JSON()
}
