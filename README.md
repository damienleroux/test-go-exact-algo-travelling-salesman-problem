# test exact algorithm performance with go

## Motivation

Just for fun to play with go and to benchmark performance using exact algorithm [Travelling salesman problem](https://en.wikipedia.org/wiki/Travelling_salesman_problem#Exact_algorithms)

## How this algo works?

1) Computes every possible tours using a list of sales. Computation is not channeled yet.

During this step, the distance from one sale to the next one is based on computing the "fly-bird distance" between two coordinates (See `RawDistance` func).

2) Selects the best tour based on the `TotalCoveredDistance` of every tours

3) Returns a tour formated as an array of steps

## Run benchmark

launch `go test -bench="." -run="^a"`

Last running benchmarks on my personal computer:

```bash
goos: windows
goarch: amd64
pkg: github.com/damienleroux/test-go-exact-algo-travelling-salesman-problem
BenchmarkGetBestRoute5-8            9765            115362 ns/op
BenchmarkGetBestRoute6-8            1411            820696 ns/op
BenchmarkGetBestRoute7-8             219           5143919 ns/op
BenchmarkGetBestRoute8-8              25          46540916 ns/op
BenchmarkGetBestRoute9-8               2         509702800 ns/op
BenchmarkGetBestRoute10-8              1        5450451300 ns/op
```

## Run  tests

launch `go test -v ./...`
