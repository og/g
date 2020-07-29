package ge

import (
	gtest "github.com/og/x/test"
	"testing"
)

type TestIsOrangeErrorWithoutMessage struct {}
func (TestIsOrangeErrorWithoutMessage) Has () bool {return false}
func (TestIsOrangeErrorWithoutMessage) Error () string {return ""}
func (TestIsOrangeErrorWithoutMessage) Check () {}

type TestIsOrangeErrorMessageTypeMustByString struct {Message int}
func (TestIsOrangeErrorMessageTypeMustByString) Has () bool {return false}
func (TestIsOrangeErrorMessageTypeMustByString) Error () string {return ""}
func (TestIsOrangeErrorMessageTypeMustByString) Check () {}

type TestIsOrangeErrorRootFieldString struct { Message string; Name string}
func (TestIsOrangeErrorRootFieldString) Has () bool {return false}
func (TestIsOrangeErrorRootFieldString) Error () string {return ""}
func (TestIsOrangeErrorRootFieldString) Check () {}

type TestIsOrangeErrorRootFieldBoolErrorPrefix struct {Message string; NoMessage bool}
func (TestIsOrangeErrorRootFieldBoolErrorPrefix) Has () bool {return false}
func (TestIsOrangeErrorRootFieldBoolErrorPrefix) Error () string {return ""}
func (TestIsOrangeErrorRootFieldBoolErrorPrefix) Check () {}

type TestIsOrangeErrorRootFieldInt struct {Message string; Age int}
func (TestIsOrangeErrorRootFieldInt) Has () bool {return false}
func (TestIsOrangeErrorRootFieldInt) Error () string {return ""}
func (TestIsOrangeErrorRootFieldInt) Check () {}

type TestFailError struct { Message string ;ErrNoMessage bool ;ErrMessageTooLong bool }
func (err TestFailError) Has() bool { return false}
func (TestFailError) Error () string {return ""}
func (TestFailError) Check () {}

type TestFailError2 struct { Message string ;ErrNoMessage bool ;ErrMessageTooLong bool }
func (err TestFailError2) Has() bool { return err.ErrNoMessage}
func (TestFailError2) Error () string {return ""}
func (TestFailError2) Check () {}

type TestFailError3 struct { Message string ;ErrNoMessage bool ;ErrMessageTooLong bool }
func (err TestFailError3) Has() bool { return err.ErrMessageTooLong}
func (TestFailError3) Error () string {return ""}
func (TestFailError3) Check () {}

type TestFailError4 struct { Message string ;ErrNoMessage bool ;ErrMessageTooLong bool }
func (err TestFailError4) Has() bool { return err.ErrNoMessage || err.ErrMessageTooLong}
func (TestFailError4) Error () string {return ""}
func (TestFailError4) Check () {}

type TestFailError5 struct { Message string ;ErrNoMessage bool ;ErrMessageTooLong bool }
func (err TestFailError5) Has() bool { return true}
func (TestFailError5) Error () string {return ""}
func (TestFailError5) Check () {}

func UtilTestIsOrangeError(as *gtest.AS, v OrangeError, matched bool, suggest string) {
	m, s := IsOrangeError(v)
	as.Equal(matched, m)
	as.Equal(suggest, s)
}
func TestIsOrangeError(t *testing.T) {
	as := gtest.NewAS(t)
	UtilTestIsOrangeError(as, TestIsOrangeErrorWithoutMessage{}, false, "IsOrangeError(errPtr) errPtr must be pointer ,like `ge.IsOrangeError(&TestIsOrangeErrorWithoutMessage{})`")
	UtilTestIsOrangeError(as, &TestIsOrangeErrorWithoutMessage{}, false, "Missing `Message string` field")
	UtilTestIsOrangeError(as, &TestIsOrangeErrorMessageTypeMustByString{}, false, "Field `Message` must be string")
	UtilTestIsOrangeError(as, &TestIsOrangeErrorRootFieldString{}, false, "`Name` root field can not must be string (except Message string)")
	UtilTestIsOrangeError(as, &TestIsOrangeErrorRootFieldBoolErrorPrefix{}, false, "`NoMessage` bool field prefix must be Err")
	UtilTestIsOrangeError(as, &TestIsOrangeErrorRootFieldInt{}, false, "`Age` only support bool or struct")
	UtilTestIsOrangeError(as, &TestFailError{}, false,
`Has() lose ErrNoMessage
please write
----------
// Generate by ge.OrangeErrorGenerateHasMethod()
func(err TestFailError) Has() bool {
    return err.ErrNoMessage || err.ErrMessageTooLong
}
----------
`)
	UtilTestIsOrangeError(as, &TestFailError2{}, false,
`Has() lose ErrMessageTooLong
please write
----------
// Generate by ge.OrangeErrorGenerateHasMethod()
func(err TestFailError2) Has() bool {
    return err.ErrNoMessage || err.ErrMessageTooLong
}
----------
`)
	UtilTestIsOrangeError(as, &TestFailError3{}, false,
`Has() lose ErrNoMessage
please write
----------
// Generate by ge.OrangeErrorGenerateHasMethod()
func(err TestFailError3) Has() bool {
    return err.ErrNoMessage || err.ErrMessageTooLong
}
----------
`)
	UtilTestIsOrangeError(as, &TestFailError4{}, true, "")
	UtilTestIsOrangeError(as, &TestFailError5{}, false,
`Maybe Has() code is wrong, when all bool field is false, Has() should return false 
please write
----------
// Generate by ge.OrangeErrorGenerateHasMethod()
func(err TestFailError5) Has() bool {
    return err.ErrNoMessage || err.ErrMessageTooLong
}
----------
`)
}

func TestOrangeErrorGenerateHasMethod(t *testing.T) {
	as := gtest.NewAS(t)
	code :=
`// Generate by ge.OrangeErrorGenerateHasMethod()
func(err ErrorSay) Has() bool {
    return err.ErrSensitiveWord || err.ErrNoMessage
}`
	as.Equal(OrangeErrorGenerateHasMethod(ErrorSay{}),code)
}