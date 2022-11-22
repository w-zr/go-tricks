# go-tricks

#### SliceTricks
See https://github.com/golang/go/wiki/SliceTricks

#### Goroutine ID
**Never use this in runtime!!!**
``` go
// see https://blog.sgmansfield.com/2015/12/goroutine-ids/
func getGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}
```

#### Merge slices without duplicates
``` go
func merge[T comparable](slices ...[]T) []T {
	m := make(map[T]struct{})

	for _, slice := range slices {
		for _, number := range slice {
			m[number] = struct{}{}
		}
	}

	result := make([]T, 0, len(m))

	for k := range m {
		result = append(result, k)
	}
	return result
}
```
