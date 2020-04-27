package glist_test
//
// import (
// 	glist "github.com/og/x/list"
// 	gtest "github.com/og/x/test"
// 	"log"
// 	"testing"
// )
//
// func TestRun(t *testing.T) {
// 	data := []int{}
// 	glist.Run(10, func(i int) (_break bool) {
// 		log.Print(i)
// 		data = append(data, i)
// 		if i==5 {
// 			return true
// 		}
// 		return
// 	})
// }
//
// func TestFind(t *testing.T) {
// 	as := gtest.NewAS(t)
// 	{
// 		sList := []string{"a","b", "c", "d"}
// 		find := false
// 		{
// 			for _,item := range sList {
// 				if item == "a" {
// 					find = true
// 					break
// 				}
// 			}
// 		}
// 		as.True(find)
// 	}
// 	{
// 		sList := []string{"a","b", "c", "d"}
// 		_, find := glist.Find(len(sList), func(i int) (find bool) {
// 			item := sList[i]
// 			return item == "c"
// 		})
// 		as.True(find)
// 	}
// 	{
// 		sList := []string{"a","b", "c", "d"}
// 		find := false
// 		findIndex := -1
// 		{
// 			for i,item := range sList {
// 				if item == "a" {
// 					find = true
// 					findIndex = i
// 					break
// 				}
// 			}
// 		}
// 		as.True(find)
// 		as.GtOrEql(1, findIndex)
// 	}
// 	{
// 		sList := []string{"a","b", "c", "d"}
// 		firstMatchIndex, find := glist.Find(len(sList), func(i int) (find bool) {
// 			item := sList[i]
// 			return item == "c"
// 		})
// 		as.Equal(firstMatchIndex, 2)
// 		as.True(find)
// 	}
// }
// func TestCheckAll(t *testing.T) {
// 	as := gtest.NewAS(t)
// 	{
// 		allPass := true
// 		sList := []string{"a","b", "c", "d"}
// 		for _,item := range sList {
// 			if item == "c" {
// 				allPass = false
// 				break
// 			}
// 		}
// 		as.False(allPass)
// 	}
// 	{
// 		sList := []string{"a","b", "c", "d"}
// 		allPass := glist.CheckAll(len(sList), func(i int) (pass bool) {
// 			item := sList[i]
// 			return item == "c"
// 		})
// 		as.False(allPass)
// 	}
// }
//
// func TestMap(t *testing.T) {
// 	as := gtest.NewAS(t)
// 	{sList := []string{"a","b", "c", "d"}
// 		// map
// 		newList := []string{}
// 		for i,item := range sList {
// 			item = item + "!"
// 			newList[i] = item
// 		}
// 		as.Equal([]string{"!a","!b", "!c", "!d"}, newList)
// 	}
// }
//
// func TestFilter(t *testing.T) {
// 	as := gtest.NewAS(t)
// 	{
// 		sList := []string{"a","b", "c", "d"}
// 		newList := []string{}
// 		// filter
// 		for _,item := range sList {
// 			if item != "c" {
// 				newList = append(newList, item)
// 			}
// 		}
// 		as.Equal([]string{"a","b", "d"}, newList)
// 	}
// }
