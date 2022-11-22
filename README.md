# go-tricks

#### SliceTricks
See https://github.com/golang/go/wiki/SliceTricks

#### Goroutine ID
see https://blog.sgmansfield.com/2015/12/goroutine-ids/

**Never use this in runtime!!!**
``` go
func getGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}
```
