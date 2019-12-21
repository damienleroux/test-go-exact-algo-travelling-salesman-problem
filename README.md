# test exact algorithm performance with go

## Motivation

Just for fun to play with go and to test performance using exact algorithm [Travelling salesman problem](https://en.wikipedia.org/wiki/Travelling_salesman_problem#Exact_algorithms)

**disclaimer**: this repository is not following best practices concerning `go` project structure. Maybe I have time, I'll give it more toughts ;)

## How this algo works?

1) Computes every possible tours using a list of sales from a file `sale.json` mocking random data points.

During this step, the distance from one sale to the next one is based on the fly-bird distance. See `rawDistance` func.

2) Selects the best tour based on the `TotalCoveredDistance` of every tours

3) Returns a tour formated as an array of steps

## Benchmark

### Locally

On my local computer (Intel Core i7-4790K (4.0 GHz)) :
- computing the perfect tour with more that 10 sales is two slow (2/3 minutes for 11 sales and 39 916 800 possible solutions).
- computing the perfect tour with less than 11 sales is maximum a 5s task (10 sales represents 3 628 800 solutions).

## test it

- download
- launch `go run .`