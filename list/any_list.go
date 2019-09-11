package glist

type AnyList struct {
	Value []interface{}
}
func (aList *AnyList) Push(v ...interface{}) *AnyList {
	aList.Value  = append(aList.Value, v...)
	return aList
}