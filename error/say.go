package ge

import (
	"log"
	"strings"
)


func ExampleSay() {
	{
		err := Say("nimo")
		// if err != nil { panic(err) }
		err.Check()
	}
	{
		err := Say("fuck")
		if err.Has() {
			switch {
			case err.MessageIsEmpty:
				log.Print("MessageIsEmpty: ", err.Message)
			case err.SensitiveWord:
				log.Print(err.Message, " SensitiveWords: " , err.SensitiveWordError.SensitiveWords)
			default:
				panic(err)
			}
		}
	}
}

// Orange error: https://github.com/og/x/tree/master/error
type ErrorSay struct {
	Message string

	SensitiveWord bool
	SensitiveWordError SensitiveWordError

	MessageIsEmpty bool
}
type SensitiveWordError struct{
	SensitiveWords []string
}
// Implements go error interface， make sure warning Unhandled error
func (err ErrorSay) Error() string {
	return err.Message
}
// Similar: if err != nil { panic(err) }
func (err ErrorSay) Check() {
	if err.Has() {
		panic(err)
	}
}
// Similar: if err != nil
func (err ErrorSay) Has() bool {
	// has 必须有单元测试确保所有的bool类型判断字段被匹配
	return err.SensitiveWord || err.MessageIsEmpty
}
func Say(message string) ErrorSay {
	if len(strings.TrimSpace(message)) == 0 {
		return ErrorSay{MessageIsEmpty: true, Message:"message is empty"}
	}
	sensitiveWords := []string{"fuck", "bitch"}
	for _, sensitiveWord := range sensitiveWords {
		if strings.Contains(message, sensitiveWord) {
			return ErrorSay{
				SensitiveWord: true,
				Message: "Watch Your Mouth",
				SensitiveWordError: SensitiveWordError{ sensitiveWords, },
			}
		}
	}
	log.Print(message)
	return ErrorSay{}
}
