package ogjson

import (
	gtest "github.com/og/x/test"
	"log"
	"testing"
)

type User struct {
	Name string
	Age int
}
var user = User {
	Name: "nimo",
	Age: 27,
}
func ExampleString() {
	type User struct {
		Name string
		Age int
	}
	user := User {
		Name: "nimo",
		Age: 27,
	}
	String(user)
	// {"Name":"nimo","Age":27}
}
func ExampleParse() {
	type User struct {
		Name string
		Age int
	}
	var user User
	// In the value pointed to by user. If v test nil or not a pointer,
	// Parse returns an InvalidUnmarshalError.
	Parse(`{"Name":"nimo","Age":27}`, &user)
}

var userJSON = `{"Name":"nimo","Age":27}`
func TestString(t *testing.T) {
	as := gtest.NewAS(t)
	{
		json := String(user)
		as.Equal(userJSON, json)
	}
	{
		json, err := StringWithErr(user)
		as.Equal(userJSON, json)
		as.Equal(nil, err)
	}
	{
		json, err := StringWithErr(log.Print)
		as.Equal("", json)
		if err == nil {
			panic("ByteWithErr(log.Print) should return error")
		}
	}
}

func TestStringSpace(t *testing.T) {
	as := gtest.NewAS(t)
	{
		json := StringSpace(user, 2)
		as.Equal("{\n  \"Name\": \"nimo\",\n  \"Age\": 27\n}", json)
	}
	{
		json, err := StringSpaceWithErr(user, 2)
		as.Equal("{\n  \"Name\": \"nimo\",\n  \"Age\": 27\n}", json)
		as.Equal(nil, err)
	}
	{
		json, err := StringSpaceWithErr(log.Print, 2)
		as.Equal("", json)
		if err == nil {
			panic("StringIndentWithErr(log.Print) should return error")
		}
	}
}

func TestBytes(t *testing.T) {
	as := gtest.NewAS(t)
	{
		json := Bytes(user)
		as.Equal([]byte(userJSON), json)
	}
	{
		user := User{
			Name: "nimo",
			Age: 27,
		}
		json, err := BytesWithErr(user)
		as.Equal([]byte(userJSON), json)
		as.Equal(nil, err)
	}
	{
		json, err := BytesWithErr(log.Print)
		as.Equal([]byte(nil), json)
		if err == nil {
			panic("ByteWithErr(log.Print) should return error")
		}
	}
}


func TestParse(t *testing.T) {
	as := gtest.NewAS(t)
	{
		{
			var user User
			Parse(userJSON, &user)
			as.Equal(User{
				Name: "nimo",
				Age: 27,
			}, user)
		}
		{
			var user User
			Parse(userJSON, user)  // not pointer
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
		}
	}
	{
		{
			var user User
			err := ParseWithErr(userJSON, &user)
			as.Equal(User{
				Name: "nimo",
				Age: 27,
			}, user)
			as.Equal(nil, err)
		}

		{
			var user User
			err := ParseWithErr(``, &user)
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}

		{
			var user User
			err := ParseWithErr(userJSON, user) // not pointer
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
			as.Equal(nil, err)
		}

		{
			var user User
			err := ParseWithErr(``, user) // not pointer
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}
	}
}


func TestParseByte(t *testing.T) {
	as := gtest.NewAS(t)
	{
		{
			var user User
			ParseBytes([]byte(userJSON), &user)
			as.Equal(User{
				Name: "nimo",
				Age: 27,
			}, user)
		}
		{
			var user User
			ParseBytes([]byte(userJSON), user)  // not pointer
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
		}
	}
	{
		{
			var user User
			ParseBytes([]byte(userJSON), &user)
			as.Equal(User{
				Name: "nimo",
				Age: 27,
			}, user)
		}

		{
			var user User
			err := ParseBytesWithErr([]byte(``), &user)
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}

		{
			var user User
			err := ParseBytesWithErr([]byte(userJSON), user) // not pointer
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
			as.Equal(nil, err)
		}

		{
			var user User
			err := ParseBytesWithErr([]byte(``), user) // not pointer
			as.Equal(User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}
	}
}

func TestStringUnfold(t *testing.T) {
	as := gtest.NewAS(t)
	{
		var user User
		userUnfoldJSON := `{
  "Name": "",
  "Age": 0
}`
		as.Equal(userUnfoldJSON, StringUnfold(user))
	}
}
func TestEmptyListMap (t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(`{"List":[],"Map":{}}`, String(struct {
		List []string
		Map map[string]interface{}
	}{}))
}

func TestStringConvInt (t *testing.T) {
	as := gtest.NewAS(t)
	query := struct {
		Page int `json:"page"`
	}{}
	Parse(`{"page": "2"}`,&query)
	as.Equal(2, query.Page)
}
func TestStringConvIntAndFloat (t *testing.T) {
	as := gtest.NewAS(t)
	{
		query := struct {
			Page int `json:"page"`
		}{}
		Parse(`{"page": "2"}`,&query)
		as.Equal(2, query.Page)
	}
	{
		query := struct {
			Page float64 `json:"page"`
		}{}
		Parse(`{"page": "2.2"}`,&query)
		as.Equal(2.2, query.Page)
	}
}
// func TestInterface (t *testing.T) {
// 	{
// 		data := struct {
// 			Date SecondTime
// 			Name string
// 		}{}
// 		Parse(`{"Date":"2020-02-28 20:48:45"}`, &data)
// 		as.Equal(data.Date.Format(gtime.Second), "2020-02-28 20:48:45")
// 		as.Equal(String(data), `{"Date":"2020-02-28 20:48:45","Name":""}`)
// 	}
// 	{
// 		data := struct {
// 			Date MinuteTime
// 		}{}
// 		Parse(`{"Date":"2020-02-28 20:48"}`, &data)
// 		as.Equal(data.Date.Format(gtime.Second), "2020-02-28 20:48:00")
// 		as.Equal(String(data), `{"Date":"2020-02-28 20:48"}`)
// 	}
// 	{
// 		data := struct {
// 			Date HourTime
// 		}{}
// 		Parse(`{"Date":"2020-02-28 20"}`, &data)
// 		as.Equal(data.Date.Format(gtime.Second), "2020-02-28 20:00:00")
// 		as.Equal(String(data), `{"Date":"2020-02-28 20"}`)
// 	}
// 	{
// 		data := struct {
// 			Date DayTime
// 		}{}
// 		Parse(`{"Date":"2020-02-28"}`, &data)
// 		as.Equal(data.Date.Format(gtime.Second), "2020-02-28 00:00:00")
// 		as.Equal(String(data), `{"Date":"2020-02-28"}`)
// 	}
// 	{
// 		data := struct {
// 			Date MonthTime
// 		}{}
// 		Parse(`{"Date":"2020-02"}`, &data)
// 		as.Equal(data.Date.Format(gtime.Second), "2020-02-01 00:00:00")
// 		as.Equal(String(data), `{"Date":"2020-02"}`)
// 	}
// 	{
// 		data := struct {
// 			Date YearTime
// 		}{}
// 		Parse(`{"Date":"2020"}`, &data)
// 		as.Equal(data.Date.Format(gtime.Second), "2020-01-01 00:00:00")
// 		as.Equal(String(data), `{"Date":"2020"}`)
// 	}
// }

func Test_ParseSliceNil(t *testing.T) {
	as := gtest.NewAS(t)
	{
		data := struct {
			List []string
		}{}
		Parse(`{"List":[]}`, &data)
		// binding 等库依赖了这个特性，所以不要改变这个特性 @nimoc
		as.Equal(data.List, []string{})
	}
	{
		data := struct {
			List []string
		}{}
		Parse(`{}`, &data)
		// binding 等库依赖了这个特性，所以不要改变这个特性 @nimoc
		as.Equal(data.List, []string(nil))
	}
}