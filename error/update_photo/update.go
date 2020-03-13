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