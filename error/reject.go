package ge

import (
	"fmt"
)

type Reject struct {
	Response interface{}
	ShouldRecord bool
}
func NotReject() Reject {
	return Reject{}
}
func (reject Reject) Fail() bool {
	return reject.Response != nil
}
func (reject Reject) Error() string {
	return fmt.Sprintf("%+v", reject.Response)
}