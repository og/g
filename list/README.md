# go-list

> go-list 下的所有方法都支持链式调用 `sList.Push("1").UnShift("a")`

### StringList

*Push*

`[* * add]`

```go
sList := glist.StringList{}
// sList.Value equal []string{}
sList.Push("a")
// sList.Value equal []string{"a"}
sList.Push("b", "c")
// sList.Value equal []string{"a", "b", "c"}
```

*Pop*

`[* * delete]`

```go
sList := glist.StringList{"a","b","c","d"}
sList.Pop()
// sList.Value equal []string{"a", "b", "c"}
sList.Pop()
// sList.Value equal []string{"a", "b"}
sList.Pop().Pop()
// sList.Value equal []string{}
sList.Pop()
// When len(sList.Value) == 0 , will not panic error
```

*PopBind*

`[* * remove]`

```go
sList := glist.StringList{"a", "b"} 
var firstString glist.StringListBindValue
sList.PopBind(&firstString)
// firstString equal {Value: "b", Has: true}
sList.PopBind(&firstString)
// firstString equal {Value: "a", Has: true}
sList.PopBind(&firstString)
// firstString equal {Value: "", Has: false}
```

*Unshift*

`[add * *]`

```go
sList := glist.StringList{}
sList.Unshift("a")
// sList.Value equal []stirng{"a"}
sList.Unshift("b")
// sList.Value equal []stirng{"a", "b"}
```

*Shift*

`[remove * *]`

```go
sList := glist.StringList{"a","b"}
sList.Shift()
// sList.Value equal []stirng{"b"}
sList.Shift()
// sList.Value equal []stirng{}
```

*ShiftBind*

`[remove * *]`

```go
sList := glist.StringList{"a", "b"} 
var lastString glist.StringListBindValue
sList.ShiftBind(&lastString)
// lastString equal {Value: "a", Has: true}
sList.ShiftBind(&lastString)
// lastString equal {Value: "b", Has: true}
sList.ShiftBind(&lastString)
// lastString equal {Value: "", Has: false}
```

*glistme*

```go
sList := glist.StringList{[]string{"a","bb","c"}}
hasTwoWords := sList.glistme(func(item, index) bool {
	return len(item) == 2 
})
// hasTwoWords equal true

hasThreeWords := sList.glistme(func(item, index) bool {
    return len(item) == 3
}) 
// hasThreeWords equal false
```

*Every*
```go
sList := glist.StringList{[]string{"a","b","c"}}
allIglistneWord := sList.Every(func (index int, item string) bool {
	return len(item) == 1
})
// allIglistneWord equal true

allIsB := sList.Every(func (index int, item string) bool {
	return item == "b"
})
// allIsB equal false
```