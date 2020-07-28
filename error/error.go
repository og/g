package ge

import (
	"reflect"
)

func IsOrangeError(err interface{}) (matched bool, suggest string) {
	rValue := reflect.ValueOf(err)
	rType := rValue.Type()
	var nameList []string
	for i:=0;i<rType.NumField();i++ {
		itemType := rType.Field(i)
		nameList = append(nameList, itemType.Name)
	}
	// 必须包含 Message 属性
	{
		pass := false
		for _, name := range nameList {
			if name == "message" {
				pass = true
				break
			}
		}
		if !pass {
			return false, "Missing message field"
		}
	}

}
