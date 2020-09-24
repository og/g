package greflect

import (
	"errors"
	"reflect"
)

type DeepEachInfo struct {
	IsRoot bool
}
func DeepEach(v interface{}, callback func(rValue reflect.Value, rType reflect.Type, field reflect.StructField)) {
	if callback == nil {
		panic(errors.New("greject.DeepEach(&v, callback) callback can not be nil"))
	}
	rootValuePtr := reflect.ValueOf(v)
	if rootValuePtr.Type().Kind() != reflect.Ptr {
		panic(errors.New("greject.DeepEach(&v, callback) v must be pointer"))
	}
	rootValue := rootValuePtr.Elem()
	rootType := rootValue.Type()
	info := DeepEachInfo{
		IsRoot: true,
	}
	coreEach(coreEachProps{
		parentValue: rootValue,
		parentType: rootType,
		field: reflect.StructField{},
		callback: callback,
		info: info,
	})
}
type coreEachProps struct {
	parentValue reflect.Value
	parentType reflect.Type
	field reflect.StructField
	callback func(rValue reflect.Value, rType reflect.Type, field reflect.StructField)
	info DeepEachInfo
}
func coreEach(props coreEachProps) {
	if !props.info.IsRoot {
		props.callback(props.parentValue, props.parentType, props.field)
	}
	if props.info.IsRoot {
		props.info.IsRoot = false
	}
	switch props.parentType.Kind() {
	case reflect.Struct:
		for i:=0;i< props.parentType.NumField();i++ {
			rValue := props.parentValue.Field(i)
			rType := rValue.Type()
			field := props.parentType.Field(i)
			coreEach(coreEachProps{
				parentValue: rValue,
				parentType: rType,
				field: field,
				callback: props.callback,
				info: props.info,
			})
		}
	case reflect.Slice:
		for i:=0;i<props.parentValue.Len();i++ {
			elementValue := props.parentValue.Index(i)
			elementType := elementValue.Type()
			coreEach(coreEachProps{
				parentValue: elementValue,
				parentType: elementType,
				field: reflect.StructField{},
				callback: props.callback,
				info: props.info,
			})
		}
	default:

	}
}
