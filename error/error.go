package ge

import (
	"reflect"
	"strings"
)
type OrangeError interface {
	Has() bool
	Error() string
	Check()
}
func IsOrangeError(errPtr OrangeError) (matched bool, suggest string) {
	rValuePtr := reflect.ValueOf(errPtr)
	if rValuePtr.Kind() != reflect.Ptr {
		return false, "IsOrangeError(errPtr) errPtr must be pointer ,like `ge.IsOrangeError(&"+ reflect.ValueOf(errPtr).Type().Name() +"{})`"
	}
	rValue := rValuePtr.Elem()
	rType := rValue.Type()
	// B: 必须包含 Message 属性
	{
		pass := false
		for i:=0;i<rType.NumField();i++ {
			itemType := rType.Field(i)
			if itemType.Name == "Message" {
				if itemType.Type.Kind() != reflect.String {
					return false, "Field `Message` must be string"
				}
				pass = true
			}
		}
		if !pass{
			return false, "Missing `Message string` field"
		}
	}
	// C: 除了 `Message string` 字段值只允许存在 `bool` 和 `struct` 两种类型
	var boolField []reflect.Value
	for i:=0;i<rType.NumField();i++ {
		itemValue := rValue.Field(i)
		itemType := rType.Field(i)
		switch itemType.Type.Kind(){
		case reflect.String:
			if itemType.Name != "Message" {
				return false, "`" + itemType.Name + "`" + " root field can not must be string (except Message string)"
			}
		case reflect.Bool:
			boolField = append(boolField, rValue.Field(i))
			// D: `bool` 类型字段只能用于标明错误种类,并且必须以 `Err` 为前缀
			if !strings.HasPrefix(itemType.Name, "Err") {
				return false, "`" + itemType.Name + "`" + " bool field prefix must be Err"
			}
			itemValue.SetBool(true)
			if !errPtr.Has() {
				return false , "Has() lose " + itemType.Name + "\nplease write\n----------\n" + OrangeErrorGenerateHasMethod(errPtr) + "\n----------\n"
			}
			// undo
			itemValue.SetBool(false)
		case reflect.Struct:
		default:
			return false, "`" + itemType.Name + "`" + " only support bool or struct"
		}
		if errPtr.Has() == true {
			return false, "Maybe Has() code is wrong, when all bool field is false, Has() should return false \nplease write\n----------\n" + OrangeErrorGenerateHasMethod(errPtr) + "\n----------\n"
		}
	}
	// E: `Has() bool` 方法类似于 `if err != nil`
	_, hasHasFunc := rType.MethodByName("Has")
	if !hasHasFunc {
		return false, "can not found method `Has() bool`"
	}
	{
		values := rValue.MethodByName("Has").Call([]reflect.Value{})
		if len(values) != 1 {
			return false, "method Has() must return bool `Has() bool`"
		}
		if values[0].Type().Kind() != reflect.Bool {
			return false, "method Has() must return bool `Has() bool`"
		}
	}
	if len(boolField) == 0 {
		return false, "field need at least one bool"
	}
	return true, ""
}

func OrangeErrorGenerateHasMethod(err OrangeError) (suggest string) {
	rValue := reflect.ValueOf(err)
	if rValue.Kind() == reflect.Ptr {
		rValue = rValue.Elem()
	}
	rType := rValue.Type()
	suggest += "// Generate by ge.OrangeErrorGenerateHasMethod()\n"
	suggest += "func(err "+rType.Name() + ") Has() bool {\n    return "
	boolFields := []string{}
	for i:=0;i<rType.NumField();i++ {
		itemType := rType.Field(i)
		if itemType.Type.Kind() == reflect.Bool {
			if strings.HasPrefix(itemType.Name, "Err") {
				boolFields = append(boolFields, "err." + itemType.Name)
			}
		}
	}
	suggest += strings.Join(boolFields, " || ") + "\n"
	suggest += "}"
	return
}