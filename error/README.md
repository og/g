# 错误处理

go 提供的 error 过于简陋,会导致90%的开发人员是不处理错误的，而是直接使用 `panic` 当做异常处理了


比如实现了一个 `UpdatePhoto(filename string) (url string, err error)` 函数

```go

var errTooLarge = errors.New("photo too large")
var errNotPhoto = errors.New("is not photo")
var errNetFail = errors.New("network connection failed")

func UpdatePhoto(filename string) (url string, err error) {
	log.Print("--------------" + filename)
	// mock
	switch filename {
	case "bigPhoto":
		return "", errTooLarge
	case "textFile":
		return "", errNotPhoto
	case "netFail":
		return "", errNetFail
	default:
		return "http://nimoc.io/" + filename, nil
	}
}
```

使用者基本上都会直接将错误当做异常抛出

```go
url, err := UpdatePhoto("bigPhoto")
if err != nil {
    panic(err)
}
log.Print(url)
```

或者在另外一个函数内再次返回错误

```go
type User struct {
	Name string
	Avatar string
}
func (user *User) Update(name string, avatar string) error {
	avatarUrl, err := UpdatePhoto(avatar)
	if err !=nil {
		return err
	}
	user.Name = name
	user.Avatar = avatarUrl
	return nil
}

func TestDoc_returnErr(t *testing.T) {
	user := User{}
	err := user.Update("nimoc", "a.jpg")
	if err !=nil {
		log.Print(err)
	}
}
```

正确的处理方式是判断每一种错误情况

```go
func TestDoc_errHandle(t *testing.T) {
	{
		url, err := UpdatePhoto("bigPhoto")
        if err!=nil {
            switch true {
            case errors.Is(errTooLarge, err):
                // 编写当图片太大时的处理代码
                // 比如编写面向用户的错误消息，更已读的本地语言，而不是代码中的英文
                log.Print("文件太大了")
            case errors.Is(errNotPhoto, err):
                // 编写当不是图片时的处理代码
                // 比如增加监控通知，因为客户端会控制图片文件类型，如果客户端上传了非图片文件类型，可能是有人恶意攻击
                log.Print("上传的不是图片")
            case errors.Is(errNotPhoto, err):
                // 编写当网络抖动时的处理代码
                // 比如重试三次
                log.Print("请稍后重试")
            default:
                panic(err)
            }
        }
        log.Print(url)
	}
}
```

上面的细致化处理有个问题：使用者必须知道 `UpdatePhoto()` 返回的错误可能是 `errTooLarge` `errNetFail` 或 `errNotPhoto`，
实际上大多数包并没有友好的告知使用者错误有哪些，需要使用者查看源码或者在文档中查找。

众所周知 godoc 会自动生成文档。
有利有弊，弊就是有些开发人员会不主动写文档。且自动生成的文档并不能自定义每个函数和注释在文档页面出现的顺序。
所以go自动生成的文档其实是接口文档，并不一定是使用文档。

go文档这种做法是正确的，维护文档会让开发者投入很大的精力，愿意写文档的人即时go提供了接口文档还是会写使用文档。
不愿意写文档的人反正不会写文档。go能基于代码生成文档起码能让不写文档的人慢慢愿意在代码中加注释来生成文档。

而最好的包设计和实现是让使用者尽量不需要看文档。nimoc.io 给出的解决方案是:

>通过自定义错误类型强制让使用者处理错误,并且明确区分异常和错误，在调用方无法处理的情况下进行 panic 操作

先看如何调用

```go
// 细致处理
{
    filename := "bigPhoto"
    url, err := docUpdate.Photo(filename)
    errorMessage := ""
    // 当错误时候， err.Fail() 会返回 true ,并会执行对应的处理函数
    if err.Fail(
        func(_netFail bool) {
            errorMessage = "网络错误，请重试"
        },
        func(_notPhoto bool) {
            errorMessage = "上传的不是图片"
        },
        func(_tooLarge bool) {
            // 错误不只是包含类型和消息，还能包含的更多有助于处理的数据比如上传限制是多少MB
            // 必要时还可以包裹另外一个包的错误类型，因为包可能依赖另外一个包
            errorMessage = "图片不能大于" + gconv.IntString(err.MaxPhotoMBLimit) + "MB"
        },
    ) {
            log.Print(errorMessage)
    }
    log.Print(url)
}
// 不处理直接抛出
{
    filename := "bigPhoto"
    url, err := docUpdate.Photo(filename)
    if err.IError() == nil {
        panic(err.IError())
    }
    log.Print(url)
}
```

当调用 err.Fail() 时候必须传入 netFail notPhoto tooLarge 三个函数，进行细致化的错误处理。
至于是函数内写什么代码就需要包使用者去决定了。如果少传函数或者多传都会报错。
函数第一个参数都会是 bool，比如 `func(_netFail bool)` 这个参数只是用于标记函数名字的可以理解为是注释，布尔值没有任何使用意义。

看完使用后再看实现：


```go
package docUpdate

import (
	"errors"
)

const maxPhotoMBLimit = 32
type ErrPhoto struct {
	isFail bool
	kind string
	message string
	MaxPhotoMBLimit int
}
// 为了满足 error interface
func (err ErrPhoto) Error() string {
	return err.message
}
// 将dict小写作为私有方法的目的是为了强制让外部只能通过 err.Fail() 处理错误
// ErrPhoto 的字典就是玩给包作者用的，而不是给包的调用方使用
func (err ErrPhoto) dict() (dict struct {
	Kind struct {
		TooLarge string
		NotPhoto string
		NetFail string
	}
}) {
	dict.Kind.TooLarge = "tooLarge"
	dict.Kind.NotPhoto = "notPhoto"
	dict.Kind.NetFail = "netFail"
	return
}
func newErrPhoto(kind string) (err ErrPhoto) {
	dict := ErrPhoto{}.dict()
	switch kind {
	default:
		// 这种情况下是包开发者代码写漏了kind判断，应该 panic 。
		panic(errors.New("newErrorPhoto(kind)  kind not found"))
	case dict.Kind.NetFail:
		err.message = "network connection failed"
	case dict.Kind.NotPhoto:
		err.message = "file is not photo"
	case dict.Kind.TooLarge:
		err.message = "photo size too large"
	}
	err.kind = kind
	err.isFail = true
	err.MaxPhotoMBLimit = maxPhotoMBLimit
	return
}
func (err *ErrPhoto) IError() error {
	if err.isFail {
		return err 
	} else {
		return nil
	}
}
func (err ErrPhoto) Fail(
	netFail func(_netFail bool),
	notPhoto func (_notPhoto bool),
	tooLarge func (_tooLarge bool),
	) (fail bool) {
	if err.isFail == false {
		return false
	}
	dict := err.dict()
	switch err.kind {
	default:
		// 这种情况下是包开发者代码写漏了kind判断，应该 panic 。
		panic(errors.New("ErrPhoto.Fail()  err.kind not found"))
	case dict.Kind.NetFail:
		netFail(true)
	case dict.Kind.NotPhoto:
		notPhoto(true)
	case dict.Kind.TooLarge:
		tooLarge(true)
	}
	return true
}
func Photo(filename string) (url string, err ErrPhoto) {
	// mock
	switch filename {
	case "bigPhoto":
		return "", newErrPhoto(err.dict().Kind.TooLarge)
	case "textFile":
		return "", newErrPhoto(err.dict().Kind.NotPhoto)
	case "netFail":
		return "", newErrPhoto(err.dict().Kind.NetFail)
	}
	url = "//:nimoc.io/" + filename
	return
}
```

实现代码非常长，而实际上逻辑非常简单。长的原因是使用了 dict 字典来弥补 go 中缺少枚举的问题。

需要说服自己：如果我是包开发者，我应该提供更细致的错误处理方案，即使需要我写很多代码，只要复杂度不高就应该让包使用者能更方便的处理错误。


> err.Fail() 是基于 nimoc.io/switch 规范，所以要记得对 fail编写单元测试，防止包开发者修改 fail 回调函数参数顺序导致的bug.