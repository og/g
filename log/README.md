# go-log 

[![GoDoc](https://godoc.org/github.com/og/x/log?status.svg)](https://godoc.org/github.com/og/x/log)

```go
example := struct {
    String string
    Number int
    Float float64
}{
    String: "abc",
    Number: 1,
    Float: 1.2422414,
}
l.V(example)
/*
    /Users/nimo/go/src/github.com/og/x/log/golog_test.go:16
    struct {
       String string;
       Number int;
       Float float64 
    }{
       String:"abc",
       Number:1,
       Float:1.2422414
    }
    -------------------------------
*/
 
```

```go
example := struct {
    String string
}{
    String: "string",
}
V(example, "github.com/og/x/log")
/*
/Users/nimo/go/src/github.com/og/x/log/golog_test.go:79
struct {
   String string 
}{
   String:"string"
}
- - - - - - - - - - - - - - - - (2)
"github.com/og/x/log"
-------------------------------
*/
```
