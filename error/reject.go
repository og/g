package ge

import (
	ogjson "github.com/og/x/json"
)

type Reject struct {
	Response interface{}
	ShouldRecord bool
}
func (reject Reject) Error() string {
	return ogjson.String(reject.Response)
}
func ErrorToReject(err error) (reject Reject, isReject bool) {
	switch err.(type) {
	case Reject:
		return err.(Reject), true
	default:
		return Reject{}, false
	}
}