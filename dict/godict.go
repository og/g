package gdict

import (
	gconv "github.com/og/x/conv"
	"github.com/pkg/errors"
	"log"
	"reflect"
	"runtime"
	"strings"
)

func setString (value reflect.Value, i int, custom Custom) {
	dictValue := ""
	keyName := value.Type().Field(i).Name
	structValue := value
	if custom.UseOtherStructFill {
		structValue = reflect.ValueOf(custom.OtherStruct)
	}
	defer func() {
		err := recover()
		if err !=nil {
			log.Printf("%#v", value)
			log.Print(i)
			panic(err)
		}
	}()
	tagDictValue := reflect.StructTag(structValue.Type().Field(i).Tag).Get(custom.StructTagName)
	if tagDictValue != "" {
		dictValue = tagDictValue
	} else {
		// 首字母转换小写
		runeList := []rune(keyName)
		runeList[0] = []rune(strings.ToLower(string(runeList[0])))[0]
		dictValue = string(runeList)
		_, file, line, _ := runtime.Caller(3)
		fileLine := file + ":" + gconv.IntString(line)
		log.Printf("%s \r\n(github.com/og/x/dict) Suggest a clear definition, example: %s string `dict:\"%s\"`", fileLine, keyName, dictValue)
	}
	value.Field(i).SetString(dictValue)
}
func fillValue (value reflect.Value, custom Custom) {
	fieldLen := value.NumField()
	for i:=0; i<fieldLen; i++ {
		fieldTypeString := value.Field(i).Type().Kind().String()
		fieldValue := value.Field(i)
		switch fieldTypeString {
		case "struct":
			fillValue(fieldValue, custom)
		case "string":
			setString(value, i, custom)
		default:
			panic(errors.New("gdict.Gen(v) v field type must be struct or string, can't be " + fieldTypeString))
		}
	}
}
// 根据 string 类型的 field name 填充 field value，如果 structTag 定义了 field value 则以 structTag 定义的优先
func Fill(v interface{}) {
	custom := Custom{
		StructTagName: "dict",
	}
	CustomFill(v, custom)
}
type Custom struct {
	StructTagName string
	UseOtherStructFill bool
	OtherStruct interface{}
}
func CustomFill(v interface{}, custom Custom) {
	value := reflect.ValueOf(v)
	fieldLen := value.Elem().NumField()
	for i:=0; i<fieldLen; i++ {
		fieldTypeString := value.Elem().Field(i).Type().Kind().String()
		switch fieldTypeString {
		case "struct":
			fillValue(value.Elem().Field(i), custom)
		case "string":
			setString(value.Elem(), i, custom)
		default:
			panic(errors.New("gdict.Gen(v) v field type must be struct or string, can't be " + fieldTypeString))
		}
	}
}