package greject_test

import (
	gjson "github.com/og/x/json"
	greject "github.com/og/x/reject"
	"log"
	"net/http"
	"testing"
)
// 每个团队的错误响应不一致,所以在项目中自定义
type ServiceLogicError struct {
	Message string
}
// 封装 Reject 是为了消除直接使用 greject.ServiceLogic 的 interface{}
func Reject(response ServiceLogicError) greject.ServiceLogic {
	return greject.ServiceLogic{ResponseServiceLogic: response}
}
// 每个团队的错误响应不一致,所以在项目中自定义
type FormatError struct {
	Field string
	Message string
}
// 封装 Reject 是为了消除直接使用 greject.FormatValidation 的 interface{}
func RejectFormatValidation(response FormatError) greject.FormatValidation {
	return greject.FormatValidation{ResponseFormatValidation: response}
}
func HTTPWriteJSON(w http.ResponseWriter, v interface{}) {
	_, err:= w.Write(gjson.Byte(v)) ; if err != nil { panic(err) }
}
func service(id string) (message string) {
	if id == "nimo" {
		panic(Reject(ServiceLogicError{Message: "id不能为 nimo"}))
	}
	return "你好,你的id是" + id
}
func ctrl (w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		panic(RejectFormatValidation(FormatError{
			Field: "id",
			Message: "id不能为空",
		}))
	}
	message := service(id)
	HTTPWriteJSON(w, message)
}
// use ctrl 是用来代理 ctrl 处理 reject
func useCtrl(ctrl func (w http.ResponseWriter, r *http.Request)) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}
			switch r.(type) {
			case greject.ServiceLogic:
				HTTPWriteJSON(w, r.(greject.ServiceLogic).ResponseServiceLogic)
				return
			case greject.FormatValidation:
				HTTPWriteJSON(w, r.(greject.FormatValidation).ResponseFormatValidation)
				return
			default:
				// 友情提醒,不要将 greject 这种明确带 response 的错误显示给请求方
				// 因为可能会导致信息泄露,如果要调试,应该将错误发生给 sentry 这样的平台
				// 而未知的 recover 一律 500
				w.WriteHeader(500)
				HTTPWriteJSON(w, "server error")
				log.Print(r)
			}
		}()
		ctrl(w, r)
	}
}
func TestHTTP (t *testing.T) {
	http.HandleFunc("/", useCtrl(ctrl))
	addr := ":4151"
	log.Print("http://localhost" + addr)
	http.ListenAndServe(addr, nil)
}
