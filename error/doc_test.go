package ge_test

import (
	"errors"
	"log"
	"testing"
)

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

// 直接抛出
func TestDoc_panicError(t *testing.T) {
	{
		url, err := UpdatePhoto("bigPhoto")
		if err != nil {
			panic(err)
		}
		log.Print(url)
	}
}

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
// 细致的错误处理
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




