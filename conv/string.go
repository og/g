package gconv

import (
	"errors"
	"reflect"
	"strconv"
)

func StringInt(s string) (i int, err error) {
	return strconv.Atoi(s)
}
func StringInt64 (s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
func StringFloat64 (s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
func StringFloat32 (s string) (float32, error) {
	f64, err := strconv.ParseFloat(s, 32)
	return float32(f64), err
}
func StringBool(s string) (bool, error) {
	switch s {
	case "true",
	"True",
	"t",
	"T",
	"1":
		return true, nil
	case "false",
	"False",
	"f",
	"F",
	"0":
		return false, nil
	}
	return false, errors.New("og/x/conv: " + s + " can't conv to bool")
}
func StringReflect(s string, rValue reflect.Value)  error {
	rType := rValue.Type()
	if rType.Kind() != reflect.Ptr {
		return errors.New("StringReflect(s, reflect.ValueOf(&v)) v must be pointer")
	}
	return coreStringReflect(s, rValue.Elem(), rType.Elem())
}
func coreStringReflect (s string, rValue reflect.Value, rType reflect.Type) error {
	switch rType.Kind() {
	case reflect.String:
		rValue.SetString(s)
	case reflect.Int:
		i, err := StringInt64(s) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int8:
		i, err := strconv.ParseInt(s, 10, 8) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int16:
		i, err := strconv.ParseInt(s, 10, 16) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int32:
		i, err := strconv.ParseInt(s, 10, 32) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Uint:
		i, err := strconv.ParseUint(s, 10, 64) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint8:
		i, err := strconv.ParseUint(s, 10, 8) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint16:
		i, err := strconv.ParseUint(s, 10, 16) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint32:
		i, err := strconv.ParseUint(s, 10, 32) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint64:
		i, err := strconv.ParseUint(s, 10, 64) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Bool:
		b, err := StringBool(s) ; if err != nil {panic(err)}
		rValue.SetBool(b)
	case reflect.Float32:
		f, err := StringFloat32(s) ; if err != nil {panic(err)}
		rValue.SetFloat(float64(f))
	case reflect.Float64:
		f, err := StringFloat64(s) ; if err != nil {panic(err)}
		rValue.SetFloat(f)
	case reflect.Array, reflect.Slice:
		elemType := rType.Elem()
		switch elemType.Kind() {
		case reflect.Uint8: // []byte
			bytes := []byte(s)
			rValue.SetBytes(bytes)
		default:
			return errors.New("field (" + rType.Name() + ")" + "can not string conv type " + rType.String())
		}
	default:
		return errors.New("field (" + rType.Name() + ")" + "can not string conv type " + rType.Kind().String())
	}
	return nil
}