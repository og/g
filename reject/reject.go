package greject

import (
	gjson "github.com/og/json"
)

type Service struct {
	// 字段写这么长是为了防止粗心复制导致与  ValidationResponse 混淆
	ServiceResponse interface{}
	ShouldRecord bool
}
func(s Service) Error() string {
	return gjson.String(s.ServiceResponse)
}
type FormatValidation struct {
	// 字段写这么长是为了防止粗心复制导致与  ServiceResponse 混淆
	ValidationResponse interface{}
}
func(s FormatValidation) Error() string {
	return gjson.String(s.ValidationResponse)
}