package l

import (
	"fmt"
	glist "github.com/og/x/list"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func getSpace(count int) string {
	if (count < 0) { return "" }
	return strings.Repeat(" ", count*2)
}
var newLine = string('\n')
func V(v ...interface{}) {
	var sList glist.StringList
	_, file, line, _ := runtime.Caller(1)
	sList.Push(file, ":", strconv.Itoa(line), newLine)
	for vIndex, item := range v {
		switch reflect.ValueOf(item).Kind().String() {
			case "struct":
				s := fmt.Sprintf("%#v", item)
				runeList := []rune(s)
				// spaceCount
				c := 0
				var sliceLevel int
				var whenScanSliceSpaceCountRecord int
				runeListLen := len(runeList)
				var isInsertDoubleQuote bool
				for index, byte := range runeList {
					var lastWord, nextWord string
					if index > 0 {
						lastWord = string(runeList[index-1])
					}
					if index < runeListLen-1 {
						nextWord = string(runeList[index+1])
					}
					word := string(byte)
					switch(word) {
					case `"`:
						isInsertDoubleQuote = !isInsertDoubleQuote
					}
					if isInsertDoubleQuote {
						sList.Push(word)
						continue
					}
					switch word {
					case "[":
						if nextWord == "]" {
							sliceLevel++
							whenScanSliceSpaceCountRecord = c
						}
					case "{":
						c++
					case "}":
						sList.Push(newLine)
						if c != 1 {
							sList.Push(getSpace(c))
						}
					}
					// struct { String string; Number int; Float float64 }{String:"abc", Number:1, Float:1.2422414}
					if lastWord == "{" && word !=" " {
						sList.Push(" ")
					}
					sList.Push(word)
					switch word {
					case "{":
						sList.Push(newLine,getSpace(c))
					case ",":
						sList.Push(newLine, getSpace(c))
					case ";":
						sList.Push(newLine, getSpace(c))
					case "}":
						c--
						if c!= 0 && c == whenScanSliceSpaceCountRecord {
							whenScanSliceSpaceCountRecord = 0
							sliceLevel--
						}
					}
				}
		case "string":
			sList.Push(item.(string))
		default:
			sList.Push(fmt.Sprintf("%#v", item))
		}
		if vIndex != len(v)-1 {
			sList.Push(newLine, "- - - - - - - - - - - - - - - - " + "(" +  strconv.Itoa(vIndex+2) +")", newLine)
		} else {
			sList.Push(newLine, "------------------------------- (End)", newLine)
		}
	}
	log.Print(sList.Join(""))
}