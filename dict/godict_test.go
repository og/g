package gdict

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type dictOrderStruct struct {
	Some string
	Status struct{
		SystemReject string		// 默认转换为小驼峰
		CheckPending string `dict:"check_pending"` // 通过 tag 自定义 value
	}
}
// var dictOrder = Gen(&dictOrderStruct{}).(dictOrderStruct)
var dictOrder = dictOrderStruct{}
func init () { Fill(&dictOrder) }
func DictOrder() dictOrderStruct {
	return dictOrder
}

func TestFill(t *testing.T) {
	assert.Equal(t, dictOrderStruct{
		Some: "some",
		Status: struct {
			SystemReject string
			CheckPending string `dict:"check_pending"`
		}{
			SystemReject: "systemReject",
			CheckPending: "check_pending",
		},
	}, DictOrder())
	assert.Equal(t, "check_pending", DictOrder().Status.CheckPending)
	assert.Equal(t, "systemReject", DictOrder().Status.SystemReject)
}
func TestCustomFill(t *testing.T) {
	type User struct {
		AppID string `db:"app_id"`
	}
	{
		var user User
		CustomFill(&user, Custom{StructTagName: "db"})
		assert.Equal(t, User{
			AppID: "app_id",
		}, user)
	}
	{
		type Order struct {
			Title string `db:"title"`
		}
		type OrderAttr struct {
			Title string
		}
		var orderAttr OrderAttr
		CustomFill(&orderAttr, Custom{StructTagName: "db", UseOtherStructFill:true, OtherStruct: Order{}})
		assert.Equal(t, OrderAttr{
			Title: "title",
		}, orderAttr)
	}
}

