package gdict_test

import (
	"errors"
	"github.com/matryer/is"
	"testing"
	"time"
)
// 宽松模式的字典
type Range struct {
	Type string
	Start time.Time
	End time.Time
}
var RangeTypeError = errors.New("range type not match")
func (self Range) Dict() (dict struct {
	Type struct {
		Month string
		Day string
	}
}) {
	dict.Type.Month ="month"
	dict.Type.Day = "day"
	return
}


func (self Range) Format() (text string, err error){
	switch self.Type {
	case self.Dict().Type.Day:
		text = self.Start.Format("2006-01-02")  + " ~ " + self.End.Format("2006-01-02")
	case self.Dict().Type.Month:
		text = self.Start.Format("2006-01")  + " ~ " + self.End.Format("2006-01")
	default:
		err = RangeTypeError
	}
	return
}


func TestRange(t *testing.T) {
	someDay, err := time.Parse("2006-01-02", "1992-12-19") ; if err != nil {panic(err)}
	r := Range{
		Type:  Range{}.Dict().Type.Day,
		Start: someDay,
		End:   someDay.AddDate(0,0,2),
	}
	is := is.New(t)
	{
		text, err := r.Format()
		is.Equal(text, "1992-12-19 ~ 1992-12-21")
		is.NoErr(err)
	}
}

// 严格模式的字典
type dateTypeDict struct {v string}
type dateStatusDict struct {v string}
type Date struct {
	Type dateTypeDict
	Status dateStatusDict
}
func (self Date) Dict() (dict struct {
	Type struct {
		Month dateTypeDict
		Day dateTypeDict
	}
	Status struct{
		Pass dateStatusDict
		Fail dateStatusDict
	}
}) {
	dict.Type.Month = dateTypeDict{"month"}
	dict.Type.Day = dateTypeDict{"day"}
	dict.Status.Pass = dateStatusDict{"pass"}
	dict.Status.Fail = dateStatusDict{"fail"}
	return
}

func TestDate(t *testing.T) {
	r := Date{
		Type:   Date{}.Dict().Type.Day,
		Status: Date{}.Dict().Status.Fail,
	}
	is := is.New(t)
	is.Equal(r.Type, r.Dict().Type.Day)
	is.Equal(r.Type.v, "day")
	is.Equal(r.Status.v, "fail")
}