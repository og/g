package greflect

import (
	"errors"
	"reflect"
)

type DeepEachInfo struct {
	IsRoot bool
}
// 前序遍历，遇到指针跳过指针回调值
type EachCallback func(rValue reflect.Value, rType reflect.Type, field reflect.StructField) EachOperator
type EachOperator string
func (v EachOperator) String() string {
	return string(v)
}
func (v EachOperator) Switch(
	ContinueHandle func(_Continue int),
	BreakHandle func(_Break bool),
	) {
	switch v {
	default:
		panic(errors.New("EachOperator can not be (" + v.String() + ")"))
	case Continue:
		ContinueHandle(0)
	case Break:
		BreakHandle(false)
	}
}
const (
	Continue EachOperator = "continue"
	Break EachOperator = "break"
)
func DeepEach(v interface{}, callback EachCallback) {
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
	callback EachCallback
	info DeepEachInfo
}
func coreEach(props coreEachProps) EachOperator {
	switch {
	case props.info.IsRoot:
	case props.parentType.Kind() == reflect.Ptr:
	default:
		operator := props.callback(props.parentValue, props.parentType, props.field)
		shouldBreak := false
		operator.Switch(func(_Continue int) {
			shouldBreak = false
		}, func(_Break bool) {
			shouldBreak = true
		})
		if shouldBreak {
			return Break
		}
	}
	if props.info.IsRoot {
		props.info.IsRoot = false
	}
	switch props.parentType.Kind() {
	case reflect.Ptr:
		if !props.parentValue.IsNil() {
			elementValue := props.parentValue.Elem()
			elementType := props.parentType.Elem()
			op := coreEach(coreEachProps{
				parentValue: elementValue,
				parentType:  elementType,
				field:       props.field,
				callback:    props.callback,
				info:        props.info,
			})
			if op == Break {
				return Break
			}
		}
	case reflect.Map:
		for _, key := range props.parentValue.MapKeys() {
			op := coreEach(coreEachProps{
				parentValue: props.parentValue.MapIndex(key),
				parentType:  key.Type(),
				field:       reflect.StructField{},
				callback:    props.callback,
				info:        props.info,
			})
			if op == Break {
				return Break
			}
		}
	case reflect.Struct:
		for i:=0;i< props.parentType.NumField();i++ {
			rValue := props.parentValue.Field(i)
			rType := rValue.Type()
			field := props.parentType.Field(i)
			op := coreEach(coreEachProps{
				parentValue: rValue,
				parentType: rType,
				field: field,
				callback: props.callback,
				info: props.info,
			})
			if op == Break {
				return Break
			}
		}
	case reflect.Slice:
		for i:=0;i<props.parentValue.Len();i++ {
			elementValue := props.parentValue.Index(i)
			elementType := elementValue.Type()
			op := coreEach(coreEachProps{
				parentValue: elementValue,
				parentType: elementType,
				field: reflect.StructField{},
				callback: props.callback,
				info: props.info,
			})
			if op == Break {
				return Break
			}
		}
	default:

	}
	return Continue
}
