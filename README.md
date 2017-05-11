# signedint
Benchmark of divison of signed ints vs. bit shifting

```
on go version go1.8.1 darwin/amd64

➤  go test -bench=.
BenchmarkSearch/Div-8         	50000000	        32.7 ns/op
BenchmarkSearch/Shift-8       	100000000	        22.8 ns/op
PASS
ok  	github.com/michaelbernstein/signedint	3.986s
```
