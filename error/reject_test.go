package ge_test

import (
	ge "github.com/og/x/error"
	gtest "github.com/og/x/test"
	"log"
	"testing"
)

func Some() ge.Reject {
	return ge.Reject{
		Response: NewFail("用户不存在"),
	}
	// return ge.NilReject()
}
func TestReject(t *testing.T) {
	as := gtest.NewAS(t)
	_ =as
	reject := Some()
	if reject.Fail() {
		if reject.ShouldRecord {
			// 记录日志
		}
		as.Equal(reject.Response, Response{
			Type: "fail",
			Msg:  "用户不存在",
			Data: nil,
		})
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