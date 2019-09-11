package greq

import (
	qs "github.com/google/go-querystring/query"
	ge "github.com/og/x/error"
	gjson "github.com/og/x/json"
	"net/url"
)


type PostJSONProps struct {
	Path string
	Query interface{}
	Data interface{}
	Receiver interface{}
}
func PostJSON (props PostJSONProps) (res Response, e Error) {
	var paramsValue url.Values
	var err error
	if props.Query != nil {
		paramsValue, err = qs.Values(props.Query); ge.Check(err)
	}
	res, e = do(Request{
		Method: "POST",
		Path: props.Path,
		Params: paramsValue,
		Headers: map[string]string {
			"Content-Type": "application/json",
		},
		DataType: Dict().Request.DataType.JSON,
		JSONData: props.Data,
	})
	if !e.Fail {
		gjson.ParseByte(res.Content, props.Receiver)
	}
	return
}

type GetJSONProps struct {
	Path string
	Query interface{}
	Receiver interface{}
}
func GetJSON (props GetJSONProps) (res Response, e Error) {
	var paramsValue url.Values
	var err error
	if props.Query != nil {
		paramsValue, err = qs.Values(props.Query); ge.Check(err)
	}
	res, e = do(Request{
		Method: "GET",
		Path: props.Path,
		Params: paramsValue,
		Headers: map[string]string {
			"Content-Type": "application/json",
		},
	})
	if !e.Fail {
		gjson.ParseByte(res.Content, props.Receiver)
	}
	return
}
