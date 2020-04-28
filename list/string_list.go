package glist

import "strings"

type StringList struct {
	Value []string
}
func (sList *StringList) Concat(v []string) StringList {
	sList.Value  = append(sList.Value, v...)
	return *sList
}
func (sList *StringList) Push(v ...string) StringList {
	sList.Value  = append(sList.Value, v...)
	return *sList
}

func (sList *StringList) Pop() StringList {
	return sList.PopBind(&StringListBindValue{})
}
type StringListBindValue struct {
	Value string
	Has bool
}
func (sList *StringList) PopBind(last *StringListBindValue) StringList {
	listLen := len(sList.Value)
	if listLen == 0 {
		/*
		Clear StringListBindValue Because in this case
			```
			list.PopBind(&last)
			// do Something..
			list.PopBind(&last)
			```
			last test same var
		*/
		last.Value = StringListBindValue{}.Value
		last.Has = false
		return *sList
	}
	last.Value = sList.Value[listLen-1]
	last.Has = true
	sList.Value = sList.Value[:listLen-1]
	return *sList
}

func (sList *StringList) Unshift(v string) StringList {
	sList.Value = append([]string{v}, sList.Value...)
	return *sList
}

func (sList *StringList) Shift() StringList {
	return sList.ShiftBind(&StringListBindValue{})
}

func (sList *StringList) ShiftBind(first *StringListBindValue) StringList {
	listLen := len(sList.Value)
	if listLen == 0 {
		/*
			Clear StringListBindValue Because in this case
				```
				list.ShiftBind(&first)
				// do Something..
				list.ShiftBind(&first)
				```
				first test same var
		*/
		first.Value = StringListBindValue{}.Value
		first.Has = false
		return *sList
	}
	first.Value = sList.Value[0]
	first.Has = true
	sList.Value = sList.Value[1:]
	return *sList
}
func (sList StringList) Find(callback func(index int, item string) (find bool)) (foundIndex int, found bool) {
	foundIndex = -1
	for index, item := range sList.Value {
		if callback(index, item) {
			foundIndex = index
			found = true
			return
		}
	}
	return
}



func (sList StringList) CheckAll(callback func(index int, item string) (pass bool)) (allPass bool) {
	for index, item := range sList.Value {
		if !callback(index, item) {
			return false
		}
	}
	return true
}
func (sList StringList) Join(sep string) string {
	return strings.Join(sList.Value, sep)
}
func (sList StringList) In (valueToFind string) bool {
	_, find := sList.Find(func(_ int, item string) bool {
		if item == valueToFind {
			return true
		}
		return false
	})
	return find
}