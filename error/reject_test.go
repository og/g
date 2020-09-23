package ge_test

import (
	ge "github.com/og/x/error"
	gtest "github.com/og/x/test"
	"log"
	"testing"
)

func Some() error {
	return ge.Reject{
		Response: NewFail("用户不存在"),
	}
	// return nil
}
func TestReject_Error(t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(ge.Reject{Response: map[string]string{"type":"pass"}, ShouldRecord: true}.Error(), `{"type":"pass"}`)

	testInterface := func(err error) {/* 编译期不报错即可 */}
	testInterface(ge.Reject{})
}
func TestErrorToReject(t *testing.T) {
	as := gtest.NewAS(t)
	{
		var err error
		err = nil
		reject, isReject := ge.ErrorToReject(err)
		as.Equal(reject, ge.Reject{})
		as.Equal(isReject, false)
	}
	{
		err := func () error {
			return ge.Reject{"abc", true}
		}()
		reject, isReject := ge.ErrorToReject(err)
		as.Equal(reject, ge.Reject{})
		as.Equal(isReject, false)
	}

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