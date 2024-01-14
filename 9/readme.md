# Remember to BET
- Benchmark
- Example
- Test

```
BenchmarkFuncSum(b *testing.B)
ExampleFuncSum()
TestFuncSum(t *testing.T)
```

# Commands

```
go test
go test -bench .
go test -cover .
go test -coverprofile c.out
go tool cover -html=c.out

```