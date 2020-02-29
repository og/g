package gis

import (
	is "github.com/og/x/test/lib"
	"testing"
)

type Test struct {
	core *is.I
}
func New(t *testing.T) Test {
	return Test{core: is.New(t)}
}
func (gt Test)Eql(a interface{}, b interface{}) {
	gt.core.Equal(a, b)
}
func (gt Test)Fail(a interface{}, b interface{}) {
	gt.core.Fail()
}
func (gt Test) NoErr(err error) {
	gt.core.NoErr(err)
}
func (gt Test) Err(err error) {
	if err == nil {
		gt.core.Logf("err: should not be nil")
	}
}
func (gt Test) True(expression bool) {
	gt.core.True(expression)
}
func (gt Test) False(expression bool) {
	gt.core.True(!expression)
}
type Flow struct {
	Actual interface{}
	Test Test
}
func (self Flow) Expect(v interface{}) {
	self.Test.Eql(self.Actual, v)

}
func (gt Test) EqlAndNoErr(actual interface{}, err error) Flow {
	gt.NoErr(err)
	return Flow{
		Actual: actual,
		Test: gt,
	}
}
func (gt Test) EqlAndErr(actual interface{}, err error) Flow {
	gt.Err(err)
	return Flow{
		Actual: actual,
		Test: gt,
	}
}