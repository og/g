package greq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	// {
	// 	res, err := do(Request{
	// 		Method: "中文",
	// 		Path: "",
	// 	})
	// 	_ = res
	// 	assert.Equal(t, true, err.Fail)
	// 	assert.Equal(t, "invalidMethod", err.Code)
	// 	assert.Equal(t, Dict().Error.Code.InvalidMethod, err.Code)
	// 	assert.Equal(t, `net/http: invalid method "中文"`, err.Msg)
	// 	assert.Equal(t, errors.New(`net/http: invalid method "中文"`), err.ToError())
	// }
	// {
	// 	res, err := do(Request{
	// 		Method: "GET",
	// 		Path: "///",
	// 	})
	// 	_ = res
	// 	assert.Equal(t, true, err.Fail)
	// 	assert.Equal(t, "invalidURL", err.Code)
	// 	assert.Equal(t, Dict().Error.Code.InvalidURL, err.Code)
	// 	assert.Equal(t, "parse \r\n: net/url: invalid control character in URL", err.Msg)
	// }
	{
		res, err := do(Request{
			Method: "GET",
			Path: "https://jsonplaceholder.typicode.com/todos/1",
			Headers: map[string]string {
				"Content-Type": "application/json",
			},
		})
		if err.Fail {
			panic(err.Msg)
		}

		assert.Equal(t, `{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}`, string(res.Content))
	}
	{
		res, err := do(Request{
			Method: "POST",
			Path: "https://jsonplaceholder.typicode.com/posts",
			Headers: map[string]string {
				"Content-Type": "application/json",
			},
			DataType: Dict().Request.DataType.JSON,
			JSONData: struct {Name string}{"nimo"},
		})
		_=res
		if err.Fail {
			panic(err.Msg)
		}
		assert.Equal(t, "{\n  \"Name\": \"nimo\",\n  \"id\": 101\n}", string(res.Content))
	}
}