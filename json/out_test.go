package ogjson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ctrl (title string, name string, age int) Response {
	if title == "" {
		return FailMsg("标题不能为空")
	}
	out := service(name, age)
	// do some
	return out.Response
}
func service (name string, age int) (out Out) {
	if name == "" {
		return OutFailMsg("姓名不能为空")
	}
	out = model(age) ; if out.Fail { return out }
	// do some
	return
}

func model (age int) (out Out) {
	if age < 18 {
		return OutFailMsg("必须成年")
	} else {
		return
	}
}

func TestOut (t *testing.T) {
	assert.Equal(t, Response{
		Type: "fail",
		Data:map[string]interface{}{},
		Code: "",
		Msg: "标题不能为空",
	}, ctrl("","",0))
	assert.Equal(t, Response{
		Type: "fail",
		Data:map[string]interface{}{},
		Code: "",
		Msg: "姓名不能为空",
	}, ctrl("title","",0))
	assert.Equal(t, Response{
		Type: "fail",
		Data:map[string]interface{}{},
		Code: "",
		Msg: "必须成年",
	}, ctrl("title","name",0))
	assert.Equal(t, Response{
		Type: "",
		Data:interface{}(nil),
		Code: "",
		Msg: "",
	}, ctrl("title","name",20))
}