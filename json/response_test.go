package gjson_test

import (
	gjson "github.com/og/x/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DictAuthStruct struct {
	NotLogin gjson.ResponseAuthCode
	NotAllowUpdate gjson.ResponseAuthCode
}
func DictAuth() DictAuthStruct {
	return DictAuthStruct{
		NotLogin: gjson.ResponseAuthCode {"NotLogin","账户信息登录状态过期，请重新登录！",},
		NotAllowUpdate: gjson.ResponseAuthCode {"NotAllowUpdate","没有更新的权限！",},
	}
}

func TestResponse_Auth(t *testing.T) {
	{
		sjson := `{"type":"auth","data":{},"code":"NotLogin","msg":"账户信息登录状态过期，请重新登录！"}`
		assert.Equal(t, sjson, gjson.String(
			gjson.Auth(DictAuth().NotLogin),
		))
	}
	{
		sjson := `{"type":"auth","data":{},"code":"NotAllowUpdate","msg":"没有更新的权限！"}`
		assert.Equal(t, sjson, gjson.String(
			gjson.Auth(DictAuth().NotAllowUpdate),
		))
	}
}

func TestResponse_FailMsg(t *testing.T) {
	{
		sjson := `{"type":"fail","data":{},"code":"","msg":"错误消息"}`
		assert.Equal(t, sjson, gjson.String(
			gjson.FailMsg("错误消息"),
		))
	}
}

type DictFailStruct struct {
	EmailExisting gjson.ResponseFailCode
}
func DictFail() DictFailStruct {
	return DictFailStruct{
		EmailExisting: gjson.ResponseFailCode {"EmailExisting","邮件已存在",},
	}
}

func TestResponse_FailCode(t *testing.T) {
	{
		sjson := `{"type":"fail","data":{},"code":"EmailExisting","msg":"邮件已存在"}`
		assert.Equal(t, sjson, gjson.String(
			gjson.FailCode(DictFail().EmailExisting),
		))
	}
}

func TestResponse_Pass(t *testing.T) {
	{
		sjson := `{"type":"pass","data":{},"code":"","msg":""}`
		assert.Equal(t, sjson, gjson.String(
			gjson.Pass(gjson.EmptyObject()),
		))
	}
	{
		type User struct{
			Name string `json:"name"`
			Age int `json:"age"`
		}
		user  := User{
			Name: "nimo",
			Age: 18,
		}
		sjson := `{"type":"pass","data":{"name":"nimo","age":18},"code":"","msg":""}`
		assert.Equal(t, sjson, gjson.String(
			gjson.Pass(user),
		))
	}
}
func TestResponse_Fail(t *testing.T) {
	{
		type User struct{
			Name string `json:"name"`
			Age int `json:"age"`
		}
		user  := User{
			Name: "nimo",
			Age: 18,
		}
		sjson := `{"type":"fail","data":{"name":"nimo","age":18},"code":"EmailExisting","msg":"邮件已存在"}`
		assert.Equal(t, sjson, gjson.String(
			gjson.Fail(user, DictFail().EmailExisting),
		))
	}
}