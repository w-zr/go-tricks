```
goos: linux
goarch: amd64
pkg: byte-slice-indexed-maps
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz

BenchmarkOptimizedStringKeyed-8         10055385	       104.9 ns/op
BenchmarkArrayKeyed-8   	        14126439	        77.97 ns/op
BenchmarkUnsafeStringKeyed-8   	        27734868	        40.39 ns/op
```
