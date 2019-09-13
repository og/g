package ge_test

import (
	ge "github.com/og/x/error"
	"log"
	"testing"
)

type Error struct {
	ge.Err
}
type ErrCodeMap struct {
	AddLastName struct{
		NameCanNotBeAnEmptyString string
		NameLengthMinimumTwoLetters string
	}
}
var errCode ErrCodeMap
func init () { ge.FillErrCode(&errCode) }
func ErrCode() ErrCodeMap {
	return errCode
}
func AddLastname (name string) (fullname string, err Error) {
	if name == "" {
		err.SetCode(ErrCode().AddLastName.NameCanNotBeAnEmptyString)
		return
	}
	if len(name) <2 {
		err.SetCode(ErrCode().AddLastName.NameLengthMinimumTwoLetters)
		return
	}
	fullname = name + " Chu"
	return
}
func ExampleErr() {
	name := ""
	fullname, gerr := AddLastname(name)
	if gerr.Fail() {
		switch gerr.Code() {
		case ErrCode().AddLastName.NameCanNotBeAnEmptyString: // name := ""
			log.Print(gerr.Msg())
			log.Print("do some empty name")
		case ErrCode().AddLastName.NameLengthMinimumTwoLetters: // name := "n"
			log.Print("do some few name")
		default:
			panic(gerr)
		}
	}
	log.Print("fullname: ", fullname)
}

func TestErr(t *testing.T) {
	ExampleErr()
}
