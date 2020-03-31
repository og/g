package gjson

import (
	"github.com/og/x/json/lib"
	"strings"
)

// encode (without error)
func String(v interface{}) (jsonString string) {
	jsonString, err := StringWithErr(v); if err != nil { panic(err) }
	return
}
// encode (with error)
func StringWithErr (v interface{}) (jsonString string, err error) {
	bjson, err := json.Marshal(v)
	if err != nil {
		return
	}
	jsonString = string(bjson)
	return
}
// encode  unfold (space 2)
func StringUnfold(v interface{}) (jsonString string) {
	jsonString, err := StringSpaceWithErr(v ,2)
	if err != nil { panic(err) }
	return
}
// encode  unfold (space 2) (with error)
func StringUnfoldWithErr(v interface{}) (jsonString string, err error) {
	jsonString, err = StringSpaceWithErr(v ,2)
	if err != nil {
		return
	}
	return
}
// encode pretty-print
func StringSpace(v interface{}, space int) (jsonString string) {
	jsonString, err := StringSpaceWithErr(v ,space)
	if err != nil { panic(err) }
	return
}
// encode pretty-print (with error)
func StringSpaceWithErr(v interface{}, space int) (jsonString string, err error) {
	bjson, err := json.MarshalIndent(v, "", strings.Repeat(" ", space))
	if err != nil {
		return
	}
	jsonString = string(bjson)
	return
}
func Bytes(v interface{}) []byte {
	bjson, err := json.Marshal(v)
	if err != nil { panic(err) }
	return bjson
}
// encode to []byte (with error)
func BytesWithErr(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}


// decode format string
// Parse(`{"name":"nimo"}`, &user)
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func Parse(jsonString string, v interface{}) {
	err := ParseWithErr(jsonString, &v)
	if err != nil { panic(err) }
}
// decode format string (with error)
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func ParseWithErr(jsonString string,  v interface{}) (err error) {
	err = json.Unmarshal([]byte(jsonString), &v)
	return
}

// decode by []byte
// Parse([]byte(`{"name":"nimo"}`), &user)
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func ParseBytes (data []byte, v interface{}) {
	err := ParseBytesWithErr(data, &v); if err != nil { panic(err) }
}
// decode by []byte (with error)
// equal json.Unmarshal(data []byte, v interface{}) error
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func ParseBytesWithErr (data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, &v)
	return
}
