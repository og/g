package gdict_test

import (
	"errors"
	"github.com/matryer/is"
	"testing"
	"time"
)
// 宽松字典
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
	// 宽松字典的缺点是可能会出现
	query := struct {
		Type string
	}{
		Type: "month",
	}
	r.Type = query.Type
	/*
	写这段代码的人家假定了 query的字典与 Range 的字典一致
	但是这种假定是不严谨的，现在写下来的代码可能是一致的，后续如果维护或者修改可能导致
	query.Type 的值不在 Range 的字典中，从而导致无法预料的错误
	*/
}

// 严格字典
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
	// 严格字典的优点是不会出现宽松字典的缺点
	query := struct {
		Type string
	}{
		Type: "month",
	}
	_=query
	// r.Type = query.Type // 赋值会报错
	r.Type = r.Dict().Type.Day // 只有通过 r.Dict().Type 才能赋值
	// r.Type = r.Dict().Status.Fail // 字典选错了也会报错（因为 Type 使用的是 dateTypeDict 而 Status使用的是 dateStatusDict
}

