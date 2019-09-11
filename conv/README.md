#  go-conv

[![GoDoc](https://godoc.org/github.com/og/x/conv?status.svg)](https://godoc.org/github.com/og/x/conv)

golang 便捷的类型转换函数


> golang 中各种 itoa  atoi iota (C语言风格)名字缩写易混淆并使得代码可读性差，各种其他类型转换函数命名不统一。通过goconv 可以编写统一命名可读性强的代码。  

# bool

## BoolInt

```go
BoolInt(true)       // int(1)
BoolInt(false)      // int(0)

BoolInt32(true)     // int32(1)
BoolInt32(false)    // int32(0)

BoolInt64(true)     // int64(1)
BoolInt64(false)    // int64(0)

BoolUint(true)      // uint(1) 
BoolUint(false)     // uint(0) 

BoolUint8(true)     // uint8(1)
BoolUint8(false)    // uint8(0)

BoolUint16(true)    // uint16(1)
BoolUint16(false)   // uint16(0)

BoolUint32(true)    // uint32(1)
BoolUint32(false)   // uint32(0)

BoolUint64(true)    // uint64(1)
BoolUint64(false)   // uint64(0)
```

## BoolString

```go
BoolString(true) // "true"
BoolString(true) // "false"
```

## int

### IntBool

```go
IntBool(1) // true 
IntBool(0) // false
```

### IntString

一般情况下使用 `IntString` `Int32String` `Int64String` 即可，这三个函数默认都是以10进制进行转换。

如需控制进制则使用 `Int64StringWithBase`， 第二个参数为进制（2~36）。


```go
IntString(123456) // "123456"
Int32String(123456) // "123456"
Int64String(123456) // "123456"
Int64StringWithBase(123456, 2) // "11110001001000000"
```

## string

字符串转换为 int float64 可能会出错，在任何情况下务必处理错误。 比如 `StringInt("1k)` 就出现错误

```go
i , err := StringInt("123") ; if err!=nil{panic(err)}
```

```go
i64 , err := StringInt64("123") ; if err!=nil{panic(err)}
```

```go
f64 , err := StringFloat64("123.1") ; if err!=nil{panic(err)}
```

```go
f32 , err := StringFloat32("123.1") ; if err!=nil{panic(err)}
```