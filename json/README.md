# go-json

> go-json 是对 "encoding/json" 包的接口友好化封装。

> 绝大部分情况下 `json.Unmarshal(data, &v)` 和 `json.Marshal(v)` 接收的 json 和结构体都是格式正确的。重复处理 `err` 很繁琐。



## Parse

```go
type User struct {
    Name string
    Age int
}
var user User
// 必须是传入 user 的指针，否则会 panic 错误 
Parse(`{"Name":"nimo","Age":27}`, &user)
```

如果出错则会 `panic(err)`，如果你想自己处理错误请使用`ParseWithErr(jsonString string, v interface{}) error`

## String

```go
type User struct {
    Name string
    Age int
}
user := User {
    Name: "nimo",
    Age: 27,
}
String(user)
```

如果出错则会 `panic(err)`，如果你想自己处理错误请使用`StringWithErr(v interface{}) error`

需要便于人类阅读的格式可以使用

- `StringSpace(v interface{}, space int) string` 
- `StringSpaceWithErr (v interface{}, sapce int) ( string,  error)`

## byte

如果要处理 byte 请使用

- [Byte()](https://godoc.org/github.com/og/x/json#Byte)
- [ByteWithErr()](https://godoc.org/github.com/og/x/json#ByteWithErr)
- [ParseByte()](https://godoc.org/github.com/og/x/json#ParseByte)
- [ParseByteWithErr()](https://godoc.org/github.com/og/x/json#ParseByteWithErr)

