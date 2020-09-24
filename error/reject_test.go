package ge_test

import (
	ge "github.com/og/x/error"
	gtest "github.com/og/x/test"
	"log"
	"testing"
)

func TestReject_Error(t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(ge.NewReject(map[string]string{"type":"pass"}, true).Error(), `{"type":"pass"}`)

	testInterface := func(err error) {/* 编译期不报错即可 */}
	testInterface(ge.NewReject(nil, false))
}
func TestErrorToReject(t *testing.T) {
	as := gtest.NewAS(t)
	{
		var err error
		err = nil
		reject, isReject := ge.ErrorToReject(err)
		as.Equal(reject, ge.NewReject(nil, false))
		as.Equal(isReject, false)
	}
	{
		err := func () error {
			return ge.NewReject("abc", false)
		}()
		reject, isReject := ge.ErrorToReject(err)
		as.Equal(reject, ge.NewReject("abc", false))
		as.Equal(isReject, true)
	}
	{
		err := func () error {
			return ge.NewReject("abc", true)
		}()
		reject, isReject := ge.ErrorToReject(err)
		as.Equal(reject, ge.NewReject("abc", true))
		as.Equal(isReject, true)
	}

}
func Some() error {
	return ge.NewReject(NewFail("用户不存在"), false)
	// return nil
}
func TestReject(t *testing.T) {
	as := gtest.NewAS(t)
	_ =as
	err := Some()
	if err != nil {
		reject, isReject := ge.ErrorToReject(err)
		if isReject {
			log.Print(reject.Response)
		} else {
			log.Print(err)
		}
	} else {
		log.Print("pass")
	}
}

type ResponseType string
func (v ResponseType) Enum() (enum struct{
	Pass ResponseType
	Fail ResponseType
}) {
	enum.Pass = "pass"
	enum.Fail = "fail"
	return
}
type Response struct {
	Type ResponseType `json:"type"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
func NewPass(data interface{}) Response {
	return Response{
		Type: Response{}.Type.Enum().Pass,
		Data: data,
	}
}
func NewFail(msg string) Response {
	return Response{
		Type: Response{}.Type.Enum().Fail,
		Msg: msg,
	}
}