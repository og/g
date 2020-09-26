package greflect

import (
	"errors"
	ogjson "github.com/og/json"
	gtest "github.com/og/x/test"
	"log"
	"reflect"
	"testing"
)
func TestDeepEachPnaicAndError(t *testing.T) {
	as := gtest.NewAS(t)
	as. PanicError("greject.DeepEach(&v, callback) callback can not be nil", func() {
		err := DeepEach(struct {}{}, nil)
		as.NoError(err)
	})
}

func TestDeepEach(t *testing.T) {
	as := gtest.NewAS(t)
	_=as
	type Info struct {
		FieldName string
		TypeName string
		TypeKind reflect.Kind
		AnonymousField bool
		JSONString string
		JSONTag string
		CanNotSet bool
	}
	var infos []Info
	type ID string
	type AnonymousCombination struct {
		Title string
	}
	type News struct {
		Content string
	}
	type IntList []int
	type Demo struct {
		Name string `json:"name"`
		Age uint
		UserID ID
		AnonymousCombination
		News News
		Hobby []string
		Array [1]string
		Numbers IntList
		NewsList []News
		StringPtr *string
		Map map[string]string
		NewsPtr1 *News
		NewsPtr2 *News
		NewsListPtr *[]News
		NewsList2Ptr []*News
	}
	testStr := "orange"
	demo := Demo{
		Name:"nimoc", Age: 27,
		UserID: "a",
		AnonymousCombination: AnonymousCombination {
			Title: "t",
		},
		News: News{Content:"c"},
		Hobby: []string{"read"},
		Array: [1]string{"a"},
		Numbers: IntList{1},
		NewsList: []News{News{}},
		StringPtr: &testStr,
		Map: map[string]string{
			"type": "pass",
		},
		NewsPtr1: nil,
		NewsPtr2: &News{Content:""},
		NewsListPtr: &[]News{{Content:"a"}},
		NewsList2Ptr: []*News{{Content:"b"}},
	}
	actualInfos := []Info{
		{
			FieldName: "Name",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"nimoc"`,
			JSONTag: "name",
		},
		{
			FieldName: "Age",
			TypeName: "uint",
			TypeKind: reflect.Uint,
			JSONString: `27`,
		},
		{
			FieldName: "UserID",
			TypeName: "ID",
			TypeKind: reflect.String,
			JSONString: `"a"`,
		},
		{
			FieldName: "AnonymousCombination",
			TypeName: "AnonymousCombination",
			TypeKind: reflect.Struct,
			AnonymousField: true,
			JSONString: `{"Title":"t"}`,
		},
		{
			FieldName: "Title",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"t"`,
		},
		{
			FieldName: "News",
			TypeName: "News",
			TypeKind: reflect.Struct,
			AnonymousField: false,
			JSONString: `{"Content":"c"}`,
		},
		{
			FieldName: "Content",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"c"`,
		},
		{
			FieldName: "Hobby",
			TypeName: "",
			TypeKind: reflect.Slice,
			JSONString: `["read"]`,
		},
		{
			FieldName: "",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"read"`,
		},
		{
			FieldName: "Array",
			TypeName: "",
			TypeKind: reflect.Array,
			JSONString: `["a"]`,
		},
		{
			FieldName: "Numbers",
			TypeName: "IntList",
			TypeKind: reflect.Slice,
			JSONString: `[1]`,
		},
		{
			FieldName: "",
			TypeName: "int",
			TypeKind: reflect.Int,
			JSONString: `1`,
		},
		{
			FieldName: "NewsList",
			TypeName: "",
			TypeKind: reflect.Slice,
			JSONString: `[{"Content":""}]`,
		},
		{
			FieldName: "",
			TypeName: "News",
			TypeKind: reflect.Struct,
			JSONString: `{"Content":""}`,
		},
		{
			FieldName: "Content",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `""`,
		},
		{
			FieldName: "StringPtr",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"orange"`,
		},
		{
			FieldName: "Map",
			TypeName: "",
			TypeKind: reflect.Map,
			JSONString: `{"type":"pass"}`,
		},
		{
			FieldName: "",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"pass"`,
			CanNotSet: true,
		},
		{
			FieldName: "NewsPtr2",
			TypeName: "News",
			TypeKind: reflect.Struct,
			JSONString: `{"Content":""}`,
		},
		{
			FieldName: "Content",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `""`,
		},
		{
			FieldName: "NewsListPtr",
			TypeName: "",
			TypeKind: reflect.Slice,
			JSONString: `[{"Content":"a"}]`,
		},
		{
			FieldName: "",
			TypeName: "News",
			TypeKind: reflect.Struct,
			JSONString: `{"Content":"a"}`,
		},
		{
			FieldName: "Content",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"a"`,
		},
		{
			FieldName: "NewsList2Ptr",
			TypeName: "",
			TypeKind: reflect.Slice,
			JSONString: `[{"Content":"b"}]`,
		},
		{
			FieldName: "",
			TypeName: "News",
			TypeKind: reflect.Struct,
			JSONString: `{"Content":"b"}`,
		},
		{
			FieldName: "Content",
			TypeName: "string",
			TypeKind: reflect.String,
			JSONString: `"b"`,
		},
	}
	err := DeepEach(&demo, func(rValue reflect.Value, rType reflect.Type, field reflect.StructField) (op EachOperator) {
		infos = append(infos, Info{
			FieldName: field.Name,
			TypeName: rType.Name(),
			TypeKind: rType.Kind(),
			AnonymousField: field.Anonymous,
			JSONString: func() string {
				if rValue.CanInterface() {
					return ogjson.String(rValue.Interface())
				} else {
					return "nil"
				}
			}(),
			JSONTag: field.Tag.Get("json"),
			CanNotSet: !rValue.CanSet(),
		})
		return
	})
	as.NoError(err)
	as.Equal(infos, actualInfos)
	if t.Failed() {
		log.Print(ogjson.StringUnfold(infos))
	}
}

func TestEachOperator(t *testing.T) {
	as := gtest.NewAS(t)
	_=as
	list := []string{"a","b","c"}
	msg := ""
	err := DeepEach(&list, func(rValue reflect.Value, rType reflect.Type, field reflect.StructField) (op EachOperator) {
		msg += rValue.String()
		if rValue.String() == "b" {
			return op.Break()
		}
		return
	})
	as.NoError(err)
	as.Equal(msg, "ab")
}

// func TestDeepEachMap(t *testing.T) {
// 	as := gtest.NewAS(t)
// 	type Item struct {
// 		Value string
// 	}
// 	type Data map[string]Item
//
// 	data := Data{
// 		"name": Item{"nimoc"},
// 		"title": Item{"abc"},
// 	}
// 	err := DeepEach(data, func(rValue reflect.Value, rType reflect.Type, field reflect.StructField) (op EachOperator) {
// 		if rType.Kind() == reflect.String {
// 			rValue.SetString(rValue.String() + "!")
// 		}
// 		return
// 	})
// 	as.NoError(err)
// 	log.Print(data)
// }
func TestDeepEachError(t *testing.T) {
	as := gtest.NewAS(t)
	err := DeepEach(map[string]int{"a":1,"b":2}, func(rValue reflect.Value, rType reflect.Type, field reflect.StructField) (op EachOperator) {
		if rValue.Int() == 2 {
			return op.Error(errors.New("value can not be 2"))
		}
		return
	})
	as.ErrorString(err, "value can not be 2")
}

