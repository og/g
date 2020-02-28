package gmap

import (
	"reflect"
	"sort"
)


type keyList struct {
	value []reflect.Value
}
func (kList keyList) String() (keys []string) {
	for _, key := range kList.value {
		keys = append(keys, key.String())
	}
	sort.Strings(keys)
	return
}
func (kList keyList) Int() (keys []int) {
	for _, key := range kList.value {
		keys = append(keys, int(key.Int()))
	}
	sort.Ints(keys)
	return
}
func (kList keyList) Float32() (keys []float32) {
	var tempKeys  []float64
	for _, key := range kList.value {
		tempKeys = append(tempKeys, key.Float())
	}
	sort.Float64s(tempKeys)
	for _, v := range tempKeys {
		keys = append(keys, float32(v))
	}
	return
}
func (kList keyList) Float64() (keys []float64) {
	for _, key := range kList.value {
		keys = append(keys, float64(key.Float()))
	}
	sort.Float64s(keys)
	return
}
/*
通过 UnsafeKeys(map[keyType]valueType).keyType()语法可以获取 map 的key列表
例如:
	data := map[string]string{"a":"1","b":"2"}
	UnsafeKeys(data).String() //  []strings{"a","b"}

	data := map[string]int{"a":1,"b":2}
	UnsafeKeys(data).String() //  []strings{"a","b"}

	data := map[int]int{1:1,2:2}
	UnsafeKeys(data).Int() //  []int{1,2}

因为 UnsafeKeys(data interface{}) 接收的参数是 interface{} 所以这个函数是类型不安全的，会留下隐患。
比如上面例子中 map[string]string 改成了 map[int]string 必须要改 UnsafeKeys(data).Int()
否则会导致报错，且在编译期无法发现。所以不要轻易使用 UnsafeKeys，而是使用 gmap 基于 UnsafeKeys 封装的
StringStringKeys StringIntKeys IntStringKeys 等方法来确保类型安全。

之所以将 UnsafeKey 公开而不是命名为 unsafeKeys 是因为可以让其他人对特点的map类型进行封装
例如：


type StringTimeMap map[string]time.Time

func (self StringTimeMap) Keys() (keys []string) {
	keys = UnsafeKeys(self).String()
	_= self[keys[0]] // 此代码是为了当 keyType 变化时编译失败
	return
}

type IntTimeMap map[int]time.Time

func (self IntTimeMap) Keys() (keys []int) {
	keys = UnsafeKeys(self).Int()
	_= self[keys[0]] // 此代码是为了当 keyType 变化时编译失败
	return
}

*/



func UnsafeKeys(data interface{}) (keys keyList){
	mapKeys := reflect.ValueOf(data).MapKeys()
	for _, key := range mapKeys {
		keys.value = append(keys.value, key)
	}
	return
}
// Any equal `map[string]interface{}`
type Any map[string]interface{}

func StringStringKeys(data map[string]string) (keys []string) {
	keys = UnsafeKeys(data).String() ; _=data[keys[0]] ; return
}
func StringIntKeys(data map[string]int) (keys []string) {
	keys = UnsafeKeys(data).String() ; _=data[keys[0]] ; return
}
func StringBoolKeys(data map[string]bool) (keys []string) {
	keys = UnsafeKeys(data).String() ; _=data[keys[0]] ; return
}
func StringInt32Keys(data map[string]int32) (keys []string) {
	keys = UnsafeKeys(data).String() ; _=data[keys[0]] ; return
}
func StringInt64Keys(data map[string]int64) (keys []string) {
	keys = UnsafeKeys(data).String() ; _=data[keys[0]] ; return
}
func StringFloat32Keys(data map[string]float32) (keys []string) {
	keys = UnsafeKeys(data).String() ; _=data[keys[0]] ; return
}
func StringFloat64Keys(data map[string]float64) (keys []string) {
	keys = UnsafeKeys(data).String() ; _=data[keys[0]] ; return
}

func IntStringKeys(data map[int]string) (keys []int) {
	keys = UnsafeKeys(data).Int() ; _=data[keys[0]] ; return
}
func IntIntKeys(data map[int]int) (keys []int) {
	keys = UnsafeKeys(data).Int() ; _=data[keys[0]] ; return
}
func IntBoolKeys(data map[int]bool) (keys []int) {
	keys = UnsafeKeys(data).Int() ; _=data[keys[0]] ; return
}
func IntInt32Keys(data map[int]int32) (keys []int) {
	keys = UnsafeKeys(data).Int() ; _=data[keys[0]] ; return
}
func IntInt64Keys(data map[int]int64) (keys []int) {
	keys = UnsafeKeys(data).Int() ; _=data[keys[0]] ; return
}
func IntFloat32Keys(data map[int]float32) (keys []int) {
	keys = UnsafeKeys(data).Int() ; _=data[keys[0]] ; return
}
func IntFloat64Keys(data map[int]float64) (keys []int) {
	keys = UnsafeKeys(data).Int() ; _=data[keys[0]] ; return
}

/*
// sort.Ints 的接口是 func Ints(a []int) { Sort(IntSlice(a)) } 所以key到只支持 int
func Int32StringKeys(data map[int32]string) (keys []int32) {
	keys = UnsafeKeys(data).Int32() ; _=data[keys[0]] ; return
}
func Int32IntKeys(data map[int32]int) (keys []int32) {
	keys = UnsafeKeys(data).Int32() ; _=data[keys[0]] ; return
}
func Int32BoolKeys(data map[int32]bool) (keys []int32) {
	keys = UnsafeKeys(data).Int32() ; _=data[keys[0]] ; return
}
func Int32Int32Keys(data map[int32]int32) (keys []int32) {
	keys = UnsafeKeys(data).Int32() ; _=data[keys[0]] ; return
}
func Int32Int64Keys(data map[int32]int64) (keys []int32) {
	keys = UnsafeKeys(data).Int32() ; _=data[keys[0]] ; return
}
func Int32Float32Keys(data map[int32]float32) (keys []int32) {
	keys = UnsafeKeys(data).Int32() ; _=data[keys[0]] ; return
}
func Int32Float64Keys(data map[int32]float64) (keys []int32) {
	keys = UnsafeKeys(data).Int32() ; _=data[keys[0]] ; return
}

func Int64StringKeys(data map[int64]string) (keys []int64) {
	keys = UnsafeKeys(data).Int64() ; _=data[keys[0]] ; return
}
func Int64IntKeys(data map[int64]int) (keys []int64) {
	keys = UnsafeKeys(data).Int64() ; _=data[keys[0]] ; return
}
func Int64BoolKeys(data map[int64]bool) (keys []int64) {
	keys = UnsafeKeys(data).Int64() ; _=data[keys[0]] ; return
}
func Int64Int32Keys(data map[int64]int32) (keys []int64) {
	keys = UnsafeKeys(data).Int64() ; _=data[keys[0]] ; return
}
func Int64Int64Keys(data map[int64]int64) (keys []int64) {
	keys = UnsafeKeys(data).Int64() ; _=data[keys[0]] ; return
}
func Int64Float32Keys(data map[int64]float32) (keys []int64) {
	keys = UnsafeKeys(data).Int64() ; _=data[keys[0]] ; return
}
func Int64Float64Keys(data map[int64]float64) (keys []int64) {
	keys = UnsafeKeys(data).Int64() ; _=data[keys[0]] ; return
}
*/
func Float32StringKeys(data map[float32]string) (keys []float32) {
	keys = UnsafeKeys(data).Float32() ; _=data[keys[0]] ; return
}
func Float32IntKeys(data map[float32]int) (keys []float32) {
	keys = UnsafeKeys(data).Float32() ; _=data[keys[0]] ; return
}
func Float32BoolKeys(data map[float32]bool) (keys []float32) {
	keys = UnsafeKeys(data).Float32() ; _=data[keys[0]] ; return
}
func Float32Int32Keys(data map[float32]int32) (keys []float32) {
	keys = UnsafeKeys(data).Float32() ; _=data[keys[0]] ; return
}
func Float32Int64Keys(data map[float32]int64) (keys []float32) {
	keys = UnsafeKeys(data).Float32() ; _=data[keys[0]] ; return
}
func Float32Float32Keys(data map[float32]float32) (keys []float32) {
	keys = UnsafeKeys(data).Float32() ; _=data[keys[0]] ; return
}
func Float32Float64Keys(data map[float32]float64) (keys []float32) {
	keys = UnsafeKeys(data).Float32() ; _=data[keys[0]] ; return
}


func Float64StringKeys(data map[float64]string) (keys []float64) {
	keys = UnsafeKeys(data).Float64() ; _=data[keys[0]] ; return
}
func Float64IntKeys(data map[float64]int) (keys []float64) {
	keys = UnsafeKeys(data).Float64() ; _=data[keys[0]] ; return
}
func Float64BoolKeys(data map[float64]bool) (keys []float64) {
	keys = UnsafeKeys(data).Float64() ; _=data[keys[0]] ; return
}
func Float64Int32Keys(data map[float64]int32) (keys []float64) {
	keys = UnsafeKeys(data).Float64() ; _=data[keys[0]] ; return
}
func Float64Int64Keys(data map[float64]int64) (keys []float64) {
	keys = UnsafeKeys(data).Float64() ; _=data[keys[0]] ; return
}
func Float64Float32Keys(data map[float64]float32) (keys []float64) {
	keys = UnsafeKeys(data).Float64() ; _=data[keys[0]] ; return
}
func Float64Float64Keys(data map[float64]float64) (keys []float64) {
	keys = UnsafeKeys(data).Float64() ; _=data[keys[0]] ; return
}
