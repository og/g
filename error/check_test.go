package ge_test

import (
	"errors"
	gconv "github.com/og/x/conv"
	ge "github.com/og/x/error"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"testing"
)


func ExampleC () {
	ageString := "10岁"
	age, err := strconv.ParseInt(ageString, 10, 664); ge.Check(err)
	log.Print(age)
}


func TestGE_Check(t *testing.T) {
	func () {
		defer func () {
			err := recover()
			if err == nil {
				t.Errorf("should panic error")
			}
		}()
		var err error
		err = errors.New("some error")
		ge.Check(err)
	}()
	func () {
		defer func () {
			err := recover()
			assert.Equal(t, err, nil)
		}()
		ge.Check(nil)
	}()
}

func TestGetStringWithError(t *testing.T) {
	defer func() {
		err := recover()
		assert.Equal(t, true, err!=nil)
	}()
	n := ge.Int(gconv.StringInt("nimo"))
	assert.Equal(t, 0, n)
}
func TestGetString(t *testing.T) {
	n := ge.Int(gconv.StringInt("123"))
	assert.Equal(t, 123, n)
}


func TestDefer(t *testing.T) {
	file, err := os.Open("some.txt") ; _= err // 故意忽略 err
	defer ge.Func(file.Close)
}