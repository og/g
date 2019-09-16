package greq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostJSON(t *testing.T) {
	type dataStruct struct {
		Name string
		ID int `json:"id"`
	}
	var data dataStruct
	_, e := PostJSON(PostJSONProps{
		Path: "https://jsonplaceholder.typicode.com/posts",
		Query: struct {
			Name string `url:"name"`
		}{"some"},
		Data: struct {Name string}{"nimo"},
		Receiver: &data,
	})
	assert.Equal(t, false, e.Fail)
	assert.Equal(t, dataStruct{Name:"nimo", ID:101}, data)
}

func TestGetJSON(t *testing.T) {
	type dataStruct struct {
		UserID int `json:"userId"`
		ID int `json:"id"`
		Title string `json:"title"`
		Body string `json:"body"`
	}
	var data dataStruct

	_, e := GetJSON(GetJSONProps{
		Path: "https://jsonplaceholder.typicode.com/posts/1",
		Query: struct {
			Name string `url:"name"`
		}{"some"},
		Receiver: &data,
	})
	assert.Equal(t, false, e.Fail)
	assert.Equal(t, dataStruct{UserID:1, ID:1, Title:"sunt aut facere repellat provident occaecati excepturi optio reprehenderit", Body:"quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"}, data)
}