package main

import (
	"github.com/damienleroux/test-go-exact-algo-travelling-salesman-problem/internal/mathUtils"
)

func duplicateWithoutOneSale(sales []Sale, index int) []Sale {
	salesLen := len(sales)
	duplicatedSales := make([]Sale, salesLen)
	copy(duplicatedSales, sales)
	duplicatedSales[index] = duplicatedSales[salesLen-1] //replace sale to delete with the last sale
	return duplicatedSales[:salesLen-1]                  //return a duplicated array without the last element
}

func createStep(lastStep Step, sale Sale) Step {
	distanceFromLastStep := 0.0
	totalCoveredDistance := 0.0
	if lastStep.Sale != nil {
		distanceFromLastStep = mathUtils.RawDistance(lastStep.Sale.Lat, lastStep.Sale.Long, sale.Lat, sale.Long)
		totalCoveredDistance = lastStep.TotalCoveredDistance + distanceFromLastStep
	}

	step := Step{&sale, &lastStep, distanceFromLastStep, totalCoveredDistance}
	return step
}

//provide to channel every possible values
func createRoutes(routesChan chan<- Step, sales []Sale, lastStep Step, safeGuard chan int, maxGoroutines int, onEnd CallbackEnd) {
	if len(sales) == 1 {
		routesChan <- createStep(lastStep, sales[0])
	} else {
		for index, sale := range sales {

			// transform the sale as a step
			step := createStep(lastStep, sale)

			// remove the sale from the current list of sales
			remainingSales := duplicateWithoutOneSale(sales, index)

			if len(safeGuard) < maxGoroutines {
				safeGuard <- 1
				releaseSafeGuard := func() { <-safeGuard }
				go createRoutes(routesChan, remainingSales, step, safeGuard, maxGoroutines, releaseSafeGuard)
			} else {
				createRoutes(routesChan, remainingSales, step, safeGuard, maxGoroutines, nil)
			}
		}
	}

	if onEnd != nil {
		onEnd()
	}
}

// Compute the best route to link a number of sales, each sale being represented by geo coordinates.
func GetBestRoute(sales []Sale, maxGoroutines int) Route {

	safeGuard := make(chan int, maxGoroutines)
	routesChan := make(chan Step)

	// Create routes
	go createRoutes(routesChan, sales, Step{}, safeGuard, maxGoroutines, nil)

	// Find best route
	bestStep := Step{nil, nil, 0, 0}
	numberOfSolutions := mathUtils.Factorial(len(sales))
	stepIndex := 0
	for step := range routesChan {

		if bestStep.Sale == nil || step.TotalCoveredDistance < bestStep.TotalCoveredDistance {
			bestStep = step
		}
		stepIndex++
		if stepIndex == numberOfSolutions {
			close(routesChan)
		}
	}

	//format route as an array of step
	steps := []Step{}
	for step := bestStep; step.Sale != nil; step = *step.PreviousStep {
		steps = append([]Step{step}, steps...)
	}
	return Route{steps}
}
