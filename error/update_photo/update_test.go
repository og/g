package docUpdate_test

import (
	gconv "github.com/og/x/conv"
	docUpdate "github.com/og/x/error/update_photo"
	"log"
	"testing"
)

func TestErrPhoto_Error(t *testing.T) {
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
}

