# gtest

> Base on https://github.com/matryer/is
> better api design



```go
package some_test

import (
	"encoding/json"
	"errors"
	"github.com/og/x/test"
	"testing"
)

func TestTest_Demo(t *testing.T) {
    is := gis.New(t)
    is.Eql("a", "a")
    is.True(true)
    is.False(false)

    is.Err(errors.New("err"))
    var err error
    is.NoErr(err)
    
    // 对 附带 err 的多返回值提供更便携的写法
    is.EqlAndNoErr(json.Marshal(map[string]int{})).
        Expect([]byte(`{}`))

    ch := make(chan int)
    is.EqlAndErr(json.Marshal(ch)).
        Expect(nil)
}
```
