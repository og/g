package ge

import (
	"log"
	"strings"
	"testing"
)
// Spec: https://github.com/og/x/tree/master/error#ErrorStruct
type ErrorSay struct {
	ErrorMessage string

	SensitiveWord bool
	SensitiveWordError SensitiveWordError

	MessageIsEmpty bool
}
func (err ErrorSay) Error() string {
	return err.ErrorMessage
}
type SensitiveWordError struct{
	SensitiveWords []string
}
func (err ErrorSay) Has() bool {
	// has 必须有单元测试确保所有的bool类型判断字段被匹配
	return err.SensitiveWord || err.MessageIsEmpty
}
func Say(message string) ErrorSay {
	if len(strings.TrimSpace(message)) == 0 {
		return ErrorSay{MessageIsEmpty: true, ErrorMessage:"message is empty"}
	}
	sensitiveWords := []string{"fuck", "bitch"}
	for _, sensitiveWord := range sensitiveWords {
		if strings.Contains(message, sensitiveWord) {
			return ErrorSay{
				SensitiveWord: true,
				ErrorMessage: "Watch Your Mouth",
				SensitiveWordError: SensitiveWordError{ sensitiveWords, },
			}
		}
	}
	return ErrorSay{}
}
func TestDoc(t *testing.T) {
	{
		err := Say("fuck")
		if err.Has() {
			switch true {
			case err.MessageIsEmpty:
				log.Print("MessageIsEmpty: ", err.ErrorMessage)
			case err.SensitiveWord:
				log.Print(err.ErrorMessage, " SensitiveWords: " , err.SensitiveWordError.SensitiveWords)
			default:
				panic(err)
			}
		}
	}
}