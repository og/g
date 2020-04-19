package glist_test

import (
	glist "github.com/og/x/list"
	"log"
	"testing"
)

func TestRun(t *testing.T) {
	data := []int{}
	glist.Run(10, func(i int) (_break glist.OP) {
		log.Print(i)
		data = append(data, i)
		if i==5 {
			return glist.Break()
		}
		return
	})
}
