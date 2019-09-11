package gconv

import "strconv"

func IntBool (i int) bool{
	if i == 1 { return true }
	return false
}
// int 类型转换为字符串（10进制）
func IntString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
// int32 类型转换为字符串（10进制）
func Int32String(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}
// int64 类型转换为字符串（10进制）
func Int64String(i int64) string {
	return strconv.FormatInt(i, 10)
}
// int64 类型转换为字符串（自定义进制）2~36
func Int64StringWithBase(i int64, base int) string {
	return strconv.FormatInt(i, base)
}