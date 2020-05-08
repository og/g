package greject

type Service struct {
	// 字段写这么长是为了防止粗心复制导致与  ValidationResponse 混淆
	ServiceResponse interface{}
	ShouldRecord bool
}
type FormatValidation struct {
	// 字段写这么长是为了防止粗心复制导致与  ServiceResponse 混淆
	ValidationResponse interface{}
}