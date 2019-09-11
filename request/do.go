package greq

import (
	"bytes"
	"errors"
	gdict "github.com/og/x/dict"
	ge "github.com/og/x/error"
	gjson "github.com/og/x/json"
	gmap "github.com/og/x/map"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)
type dictT struct {
	Request struct{
		Method struct{
			Get string `dict:"GET"`
			Post string `dict:"POST"`
			Head string `dict:"HEAD"`
			Options string `dict:"OPTIONS"`
			Put string `dict:"put"`
			Patch string `dict:"patch"`
		}
		ResponseType struct {
			JSON string `dict:"json"`
			Text string `dict:"text"`
		}
		DataType struct{
			JSON string `dict:"json"`
		}
	}
	Error struct{
		Code struct{
			Error string `dict:"error"`
			InvalidMethod string `dict:"invalidMethod"`
			InvalidURL string `dict:"invalidURL"`
		}
	}
}

var dict dictT

func init () {
	gdict.Fill(&dict)
}
func Dict() dictT {
	return dict
}

type Request struct {
	Method string
	Path string
	Params url.Values
	Headers map[string]string
	DataType string
	FormData interface{} `note:"Available when the request is a POST PUT PATCH"`
	JSONData interface{} `note:"json var"`
	Timeout time.Duration
	Cookie http.Cookie
	ResponseType string
	MaxContentLength int
}
type Response struct {
	Status int
	Headers map[string]string
	Content []byte
}
type Error struct {
	Fail bool
	Code string
	Msg string
}
func (err *Error) ToError () error {
	return errors.New(err.Msg)
}

func do (req Request) (res Response, e Error) {
	client := &http.Client{}
	requestBody := bytes.NewReader([]byte{})
	if req.JSONData != nil {
		requestBody = bytes.NewReader(gjson.Byte(req.JSONData))
	}
	var requestURL string
	{
		requestURL = req.Path
		queryString := req.Params.Encode()
		if queryString != "" {
			requestURL += "?" + req.Params.Encode()
		}
	}
	request, err := http.NewRequest(req.Method, requestURL, requestBody)
	if err != nil {
		errMessage := err.Error()
		switch 0 {
		case strings.LastIndex(errMessage, "net/http: invalid method"):
			e.Fail = true
			e.Code = Dict().Error.Code.InvalidMethod
			e.Msg = err.Error()
			return
		case strings.LastIndex(errMessage, "parse"):
			e.Fail = true
			e.Code = Dict().Error.Code.InvalidURL
			e.Msg = err.Error()
			return
		}
		e.Fail = true
		e.Code = Dict().Error.Code.Error
		e.Msg = err.Error()
		return
	}
	for _, key := range gmap.Keys(req.Headers).String() {
		request.Header.Set(key, req.Headers[key])
	}
	response, err := client.Do(request); ge.Check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body); ge.Check(err)
	res.Content = body
	return
}
