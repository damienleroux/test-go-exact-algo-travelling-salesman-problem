# Resolving the Travelling salesman problem using exact algorithme & go

## Motivation

Just for fun to play with go and to benchmark performance using exact algorithm [Travelling salesman problem](https://en.wikipedia.org/wiki/Travelling_salesman_problem#Exact_algorithms)

## How this algo works?

All begins with the function `GetBestRoute(sales , maxGoroutines)` with `sales` a list of geolocalised sales to deliver and `maxGoroutines` the number of possible parralel computation (used by the algorithm and depening on your CPU core number)

1) Computes every possible tours using a list of sales. Computation may or not use multiple thread (go routine): this option is a parameter when calling `GetBestRoute`

During this step, the distance from one sale to the next one is based on computing the "fly-bird distance" between two coordinates (See `RawDistance` func).

2) Selects the best tour based on the `TotalCoveredDistance` of every tours

3) Returns a tour formated as an array of steps

## Run benchmark

launch `go test -bench="." -run="^a" -v ./...`

Last running benchmarks on my personal computer for `X` number of sales (`BenchmarkGetBestRouteX`):

```config
goos: windows
goarch: amd64
CPU Intel Core i7-4790K (4.0 GHz) <= 4 Cores
```

Below the benchmark launched without go routines (only one thread). 30% of my cpu used over 90% available

```bash
BenchmarkGetBestRoute5                   7922             159934 ns/op
BenchmarkGetBestRoute6                   1198             985810 ns/op
BenchmarkGetBestRoute7                    182            6719754 ns/op
BenchmarkGetBestRoute8                     21           57523724 ns/op
BenchmarkGetBestRoute9                      2          540000000 ns/op  540ms
BenchmarkGetBestRoute10                     1         6446002000 ns/op     6s
BenchmarkGetBestRoute11                     1        59872999000 ns/op    59s
```

Below the benchmark launched with go routines (up to 100 threads). 90% of my cpu used for that over 90% available.

```bash
BenchmarkGetBestRouteMultithread5       12206              93888 ns/op
BenchmarkGetBestRoutMultithreade6        2998             407605 ns/op
BenchmarkGetBestRouteMultithread7         481            2482329 ns/op
BenchmarkGetBestRouteMultithread8          58           18724136 ns/op
BenchmarkGetBestRouteMultithread9           6          175166767 ns/op  175ms
BenchmarkGetBestRouteMultithread10          1         1647998600 ns/op   1,6s
BenchmarkGetBestRouteMultithread11          1        19823001300 ns/op    19s
```


## Run  tests

launch `go test -v ./...`
