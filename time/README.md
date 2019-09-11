# go-time

## GetDayRange

使用 `GetDayRange` 前需要了解 gotime 中公开的 `Range` 类型和 `RangeDict` 字典结构 

```go
type Range struct {
	Type string
	Start string
	End string
}
```

`Range.Type` 是字符串，可以使用 `RangeDict.Type.{kind}` 赋值给 `Range.Type`

```go
var RangeDict = struct {
	Type struct{
		Year string
		Month string
		Day string
	}
}{
	Type: struct{
		Year string
		Month string
		Day string
	} {
		Year: "year",
		Month: "month",
		Day: "day",
	},
}
```

搜索最近七天
```go
sevenDayList := GetDayRange(Range{
    Type: RangeDict.Type.Day,
    Start: Now().SubtractDay(7-1).FormatBaseDay(),
    End: Now().FormatBaseDay(),
})
log.Print("sevenDays", sevenDayList)
// [2019-08-14 2019-08-15 2019-08-16 2019-08-17 2019-08-18 2019-08-19 2019-08-20]
```