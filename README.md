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

## Run  tests

launch `go test -v ./...`
