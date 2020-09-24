package ge

import (
	ogjson "github.com/og/x/json"
)

type reject struct {
	Response interface{}
	ShouldRecord bool
}
func (reject reject) Error() string {
	return ogjson.String(reject.Response)
}
func ErrorToReject(err error) (rejectValue *reject, isReject bool) {
	switch err.(type) {
	case *reject:
		return err.(*reject), true
	default:
		return &reject{}, false
	}
}
func NewReject(response interface{}, shouldRecord bool) error {
	return &reject{
		Response: response,
		ShouldRecord: shouldRecord,
	}
}