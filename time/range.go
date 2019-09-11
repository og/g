package gtime

import (
	gdict "github.com/og/x/dict"
	"time"
)

type Range struct {
	Type string
	Start time.Time
	End time.Time
}
type dict struct {
	Range struct{
		Type struct{
			Year string `dict:"year"`
			Month string `dict:"month"`
			Day string `dict:"day"`
		}
	}
}
var protectDict = dict{}
func init () {
	gdict.Fill(&protectDict)
}
func Dict() dict {
	return protectDict
}