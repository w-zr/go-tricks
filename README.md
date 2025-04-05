# go-tricks

#### SliceTricks
See https://go.dev/wiki/SliceTricks

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

#### get caller name and measure time elapsed
See https://stackoverflow.com/questions/45766572/is-there-an-efficient-way-to-calculate-execution-time-in-golang
```go
// callerName returns the name of the function skip frames up the call stack.
func callerName(skip int) string {
	const unknown = "unknown"
	pcs := make([]uintptr, 1)
	n := runtime.Callers(skip+2, pcs)
	if n < 1 {
		return unknown
	}
	frame, _ := runtime.CallersFrames(pcs).Next()
	if frame.Function == "" {
		return unknown
	}
	return frame.Function
}

// timer returns a function that prints the name of the calling
// function and the elapsed time between the call to timer and
// the call to the returned function. The returned function is
// intended to be used in a defer statement:
//
//	defer timer()()
func timer() func() {
	name := callerName(1)
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
    defer timer()()
    time.Sleep(time.Second * 2)
}   // prints: main.main took 2s
```

#### byte slice to string in unsafe way
```go
func unsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// After go 1.20
// See https://go.dev/src/strings/builder.go#L48
func unsafeBytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
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
