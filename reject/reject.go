package greject

type ServiceLogic struct {
	// 字段写这么长是为了防止粗心复制导致与  FormatValidation 混淆
	ResponseServiceLogic interface{}
}
type FormatValidation struct {
	// 字段写这么长是为了防止粗心复制导致与  ServiceLogic 混淆
	ResponseFormatValidation interface{}
}