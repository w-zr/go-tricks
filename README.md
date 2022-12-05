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

#### Functional interface
```go
// Handler is an interface with only one func.
type Handler interface {
    Do (k, v any)
}

// HandlerFunc is a function type that implements Handler
type HandlerFunc func(k, v any)

func (hf HandlerFunc) Do(k, v any) {
    hf(k, v)  // call itself in interface function. 
}

// Usage
func Handle(h Handler) {
    // call h.Do()
}

func HandleFunc(f func(k, v any)) {
    Handle(HandlerFunc(f))
}
```

#### TestMain
```go
import (
    "testing"
    "os"
)

func TestMain(m *testing.M) {
    log.Println("Do stuff BEFORE the tests!")
    defer log.Println("Do stuff AFTER the tests!")

    os.Exit(m.Run())
}
```

### From https://github.com/cristaloleg/go-advice
#### Check interface implementation during compilation
```go
var _ io.Reader = (*MyFastReader)(nil)
```

#### To prevent struct comparison
```go
type Point struct {
    _ [0]func()	// unexported, zero-width non-comparable field
    X, Y float64
}
```

#### To prevent unkeyed literals
```go
type Point struct {
    X, Y float64
    _    struct{} // to prevent unkeyed literals
}

Point{1,1} // error
Point{X: 1, Y: 1} // no error
```

#### Check if there are mistakes in code formatting
```bash
diff -u <(echo -n) <(gofmt -d .)
```

#### Go test pprof
```bash
# run go test
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

# use go tool
go tool pprof cpu.prof
go tool pprof mem.prof
```
