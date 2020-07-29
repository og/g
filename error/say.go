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
			case err.ErrNoMessage:
				log.Print(err.Message)
			case err.ErrSensitiveWord:
				log.Print(" SensitiveWords: " , err.SensitiveWord.SensitiveWords)
			default:
				panic(err)
			}
		}
	}
}

// Orange error: https://github.com/og/x/tree/master/error
type ErrorSay struct {
	Message string

	ErrNoMessage bool

	ErrSensitiveWord bool
	SensitiveWord SensitiveWordError
}
type SensitiveWordError struct{
	SensitiveWords []string
}
// Implements go error interface， make sure warning Unhandled error
func (err ErrorSay) Error() string {
	return err.Message
}
// Similar: if err != nil
func (err ErrorSay) Has() bool {
	return err.ErrSensitiveWord || err.ErrNoMessage
}
// Similar: if err != nil { panic(err) }
func (err ErrorSay) Check() {
	if err.Has() {
		panic(err)
	}
}

func Say(message string) ErrorSay {
	if len(strings.TrimSpace(message)) == 0 {
		return ErrorSay{ErrNoMessage: true, Message:"message is empty"}
	}
	sensitiveWords := []string{"fuck", "bitch"}
	for _, sensitiveWord := range sensitiveWords {
		if strings.Contains(message, sensitiveWord) {
			return ErrorSay{
				ErrNoMessage: true,
				Message: "Watch Your Mouth",
				SensitiveWord: SensitiveWordError{ sensitiveWords, },
			}
		}
	}
	log.Print(message)
	return ErrorSay{}
}
