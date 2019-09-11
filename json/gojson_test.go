package gjson

import (
	"github.com/stretchr/testify/assert"
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
	// In the value pointed to by user. If v is nil or not a pointer,
	// Parse returns an InvalidUnmarshalError.
	Parse(`{"Name":"nimo","Age":27}`, &user)
}

var userJSON = `{"Name":"nimo","Age":27}`
func TestString(t *testing.T) {
	{
		json := String(user)
		assert.Equal(t, userJSON, json)
	}
	{
		json, err := StringWithErr(user)
		assert.Equal(t, userJSON, json)
		assert.Equal(t, nil, err)
	}
	{
		json, err := StringWithErr(log.Print)
		assert.Equal(t, "", json)
		if err == nil {
			panic("ByteWithErr(log.Print) should return error")
		}
	}
}

func TestStringSpace(t *testing.T) {
	{
		json := StringSpace(user, 2)
		assert.Equal(t, "{\n  \"Name\": \"nimo\",\n  \"Age\": 27\n}", json)
	}
	{
		json, err := StringSpaceWithErr(user, 2)
		assert.Equal(t, "{\n  \"Name\": \"nimo\",\n  \"Age\": 27\n}", json)
		assert.Equal(t, nil, err)
	}
	{
		json, err := StringSpaceWithErr(log.Print, 2)
		assert.Equal(t, "", json)
		if err == nil {
			panic("StringIndentWithErr(log.Print) should return error")
		}
	}
}

func TestByte(t *testing.T) {
	{
		json := Byte(user)
		assert.Equal(t, []byte(userJSON), json)
	}
	{
		user := User{
			Name: "nimo",
			Age: 27,
		}
		json, err := ByteWithErr(user)
		assert.Equal(t, []byte(userJSON), json)
		assert.Equal(t, nil, err)
	}
	{
		json, err := ByteWithErr(log.Print)
		assert.Equal(t, []byte(nil), json)
		if err == nil {
			panic("ByteWithErr(log.Print) should return error")
		}
	}
}


func TestParse(t *testing.T) {
	{
		{
			var user User
			Parse(userJSON, &user)
			assert.Equal(t, User{
				Name: "nimo",
				Age: 27,
			}, user)
		}
		{
			var user User
			Parse(userJSON, user)  // not pointer
			assert.Equal(t, User{
				Name: "",
				Age: 0,
			}, user)
		}
	}
	{
		{
			var user User
			err := ParseWithErr(userJSON, &user)
			assert.Equal(t, User{
				Name: "nimo",
				Age: 27,
			}, user)
			assert.Equal(t, nil, err)
		}

		{
			var user User
			err := ParseWithErr(``, &user)
			assert.Equal(t, User{
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
			assert.Equal(t, User{
				Name: "",
				Age: 0,
			}, user)
			assert.Equal(t, nil, err)
		}

		{
			var user User
			err := ParseWithErr(``, user) // not pointer
			assert.Equal(t, User{
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
	{
		{
			var user User
			ParseByte([]byte(userJSON), &user)
			assert.Equal(t, User{
				Name: "nimo",
				Age: 27,
			}, user)
		}
		{
			var user User
			ParseByte([]byte(userJSON), user)  // not pointer
			assert.Equal(t, User{
				Name: "",
				Age: 0,
			}, user)
		}
	}
	{
		{
			var user User
			ParseByte([]byte(userJSON), &user)
			assert.Equal(t, User{
				Name: "nimo",
				Age: 27,
			}, user)
		}

		{
			var user User
			err := ParseByteWithErr([]byte(``), &user)
			assert.Equal(t, User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}

		{
			var user User
			err := ParseByteWithErr([]byte(userJSON), user) // not pointer
			assert.Equal(t, User{
				Name: "",
				Age: 0,
			}, user)
			assert.Equal(t, nil, err)
		}

		{
			var user User
			err := ParseByteWithErr([]byte(``), user) // not pointer
			assert.Equal(t, User{
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
	{
		var user User
		userUnfoldJSON := `{
  "Name": "",
  "Age": 0
}`
		assert.Equal(t, userUnfoldJSON, StringUnfold(user))
	}
}