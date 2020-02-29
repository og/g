package gis_test
import (
	"encoding/json"
	"errors"
	gis "github.com/og/x/test"
	"testing"
)

func TestTest_EqlAndNoErr(t *testing.T) {
	is := gis.New(t)
	is.EqlAndNoErr(json.Marshal(map[string]int{})).
		Expect([]byte(`{}`))

	ch := make(chan int)
	is.EqlAndErr(json.Marshal(ch)).
		Expect(nil)
}

func TestTest_True_False(t *testing.T) {
	is := gis.New(t)
	is.True(true)
	is.False(false)
}
func TestTest_Eql(t *testing.T) {
	is := gis.New(t)
	is.Eql("a", "a")
}
func TestTest_Err(t *testing.T) {
	is := gis.New(t)
	is.Err(errors.New("err"))
	var err error
	is.NoErr(err)
}