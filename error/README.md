# orange error

> 提供一种更友好的 go 错误处理解决方案

可自由选择章节阅读

1. [ErrorSay示例](#ErrorSay)
2. [实现规则](#实现规则)
3. [为什么需要 orange error](#为什么需要 orange error)


## ErrorSay示例

例如需要调用 一个 `ge.Say(message string) ErrorSay` 函数，使用体验是：

当需要 panic 任何错误时候调用 `err.Check()`

```go
package main
import (
 "github.og/x/error"
)
err := ge.Say("orange")
// Check() 类似 if err != nil { panic(err) }
err.Check()
```

如果没有处理错误，则会出现 warning: 

```go
Say("orange")  
// ⚠️ Unhandled error
```

因为 Say 函数的返回值 `ErrorSay` 实现了 `Error()` 接口符合 `go error`，所以当不处理错误时候回出现 `Unhandled error`。

当需要细致处理时候可使用 `Has() bool` 和 `switch` 处理

先粗略看下 `ErrorSay` 结构体：

```go
type ErrorSay struct {
	Message string
	ErrSensitiveWord bool
	SensitiveWord SensitiveWordError

	ErrNoMessage bool
}
type SensitiveWordError struct{
	SensitiveWords []string
}
```

```go
package main
import (
 "github.og/x/error"
)
err := ge.Say("fuck")
if err.Has() {
    switch {
    case err.ErrNoMessage:
        log.Print(err.Message)
    case err.ErrSensitiveWord:
        log.Print(" SensitiveWords: " , err.SensitiveWord.SensitiveWords)
    default:
        panic(err)
    }
}
```

orange error 的优点是：

1. `err.Has()` 比 `err != nil` 可读性更高
2. 在IDE中通过输入 `err.Err` 能联想出所有错误种类，不用担心难以找到错误种类。
3.  `err.SensitiveWord.SensitiveWords` 包含了敏感词列表，比 go error 的 `Error()` 更详细。

> 实现源码: https://github.com/og/x/blob/master/error/say.go


## 实现规则

orange error结构体要求

满足 `OrangeError interface`

```go
type OrangeError interface {
	Has() bool
	Error() string
	Check()
}
// 当一个函数会有2种错误， 1 消息不能为空 2 消息包含敏感信息 。请参考 ErrorSay 的实现
// Orange error: https://github.com/og/x/tree/master/error
type ErrorSay struct {
	Message string
	ErrSensitiveWord bool
	SensitiveWord SensitiveWordError

	ErrNoMessage bool
}
type SensitiveWordError struct{
	SensitiveWords []string
}
func (err ErrorSay) Error() string {
	return err.Message
}
func (err ErrorSay) Has() bool {
	return err.ErrSensitiveWord || err.ErrNoMessage
}
func (err ErrorSay) Check() {
	if err.Has() {
		panic(err)
	}
}
```
参考 `ge.ErrorSay{}`

A: 定义结构体上一行必须附带注释以便以使用者了解用法 `// Orange error: https://github.com/og/x/tree/master/error`  
B: 有 `Message string` 字段  
C: 根节点字段除了 `Message string` 字段值只允许存在 `bool` 和 `struct` 两种类型
D: 表示错误种类的字段只能是`bool` 类型，并且必须以 `Err` 为前缀        
E: `Has()` 的实现是所有 `bool` 类型字段进行逻辑或的判断  
F: `Check()`的实现必须是 `if err.Has() { panic(err) }`
G: 实现 `Error() string` 方法  
H: 必须在单元测试中使用 `ge.IsOrangeError()` 验证实现
I: 结构体名称必须以 `Error` 作为前置，例如 `ErrorSay` `ErrorOpen` 


> 为什么有这些要求？  
> A: go 的一项优点是可以通过IDE快速跳转查看某个库的实现，在结构体上加上 orange error 的文档地址能让包的使用者了解orange error.
>    
> B: 错误应当包含可读的消息，统一使用 `Message string` 存放错误消息。
>
> C: `bool` 用于指明是什么样的错误，`struct` 用于包含详细的错误信息.一些 `string int` 等类型可以包含在 `struct` 中。
>
> D:  为了在IDE中输入 `err.Err` 时通过智能提示快速索引所有错误
>
> E: `Has() bool` 方法类似于 `if err != nil`
>
> F: `Check()` 方法就是一个便捷发放，目的是让代码更漂亮
>
> G: 实现 `Error() string` 是为了符合 go error interface,这样能在忘记处理错误时候出现 ⚠️ Unhandled error
>
> H: `ge.IsOrangeError()` 能够确保所有 orange error 的方法一致，并对容易出现错误的 `Has() bool` 进行快速单元测试。
>
> I: 结构体使用 `Error` 作为前置是为了统一命名 


为了快速检查是否实现与 orange error 规范一致，并防止 `Has` 发放实现时遗漏判断条件，
可使用 `ge.IsOrangeError(err interface)  (matched bool, suggest string)` 验证。

```go
func TestIsOrangeError(t testing.T) {
    // err := ErrorSome{}
    matched, suggest := ge.IsOrangeError(err)
    if !matched {
        t.Fatal(suggest)
    }
}
```

## 为什么需要 orange error

Go 错误处理核心是速错，即函数调用时可能会出现错误返回参数附带 error 。
error 有一个 `Error() string` 方法返回错误信息字符串。

速错比 `try catch` 优秀，因为速错能将快速定位错误，但是速错附带的条件就是处理错误繁琐。

例如其他语言解析JSON只需要 

```js
const data = JSON.parse(`{"name":"nimoc"}`)
```

而 go 解析JSON则是：

```go
// 强类型语言定义结构体
data := struct {
    Name string `json:"name"`
}{}
// 定义bytes 因为语言层面使用 []byte 性能更高
jsonBytes := []byte(`{"name":"nimoc"}`)
// &data 通过指针绑定json 
err := json.Unmarshal(jsonBytes, &data)
// 没有错误时 err 等于 nil
if err != nil {
    panic(err)
    /*  or
    log.Fatal("json 解析错误：", err.Error())
    */
} 
```

带着其他语言的使用习惯去看go代码会觉得麻烦，深入接触go之后会体会到go要求开发人员编写非常严谨的代码。
每个可能出错的函数的返回参数都附带 error,让开发人员根据业务场景选择中断程序提示或者通过panic交由上层处理。
> `panic` 类似其他语言的 `throw`

无论开发人员喜不喜欢，go要求开发人员必须采取速错的方式处理错误。

这带来了一个问题： 代码中充斥着大量的 `if err != nil {panic(err)}`

按道理应该判断错误的类型，根据不同的错误进行不同的处理。但是基本上所有的包都没有提供有好的错误判断。

无法知道到底是哪种错误，通过 `err.Error()` 返回的字符串进行判断不稳定不准确。

有些库会提供错误变量变让开发者判断，例如

```go
err := foo.Do("")
if err !=nil {
 switch err {
    case foo.ErrNoMessage:
        log.Print("message can not empty")
    case foo.ErrMessageTooLong:
        log.Print("message is too long")
    default:
        panic(err)   
 }
}
```

由于一个库可能提供多个函数，使用者无法快速知道每个函数对应可能出现的错误，并且很多库没有提供 `{package}.Err{Kind}` 这样的变量。
这就导致虽然使用速错来处理错误了，实际上却跟`try catch` 方案没有分别。

这个锅要甩给官方，官方建议了速错，提供了一个 `error`

```go
type error interface {
	Error() string
}
```

在语言层面上没有对错误处理做足够多的支持，也没有做出更详细的指导。故此 使用orange error 提供更友好的错误处理方式。