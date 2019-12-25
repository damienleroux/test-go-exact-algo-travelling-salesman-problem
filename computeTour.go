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

func createRoutes(sales []Sale, lastStep Step) []Step {
	// create the number of possible solutions considering the total numerb of sales
	steps := make([]Step, mathUtils.Factorial(len(sales)))

	stepIndex := 0
	for index, sale := range sales {

		// transform the sale as a step
		step := createStep(lastStep, sale)

		// remove the sale from the current list of sales
		remainingSales := duplicateWithoutOneSale(sales, index)

		if len(remainingSales) > 0 {
			for _, step := range createRoutes(remainingSales, step) {
				steps[stepIndex] = step
				stepIndex++
			}
		} else {
			// no more remaining sales, only one solution possible
			steps[stepIndex] = step
			stepIndex++
		}
	}
	return steps
}

func getBestRoute(sales []Sale) Route {
	// Create routes
	routes := createRoutes(sales, Step{})

	bestStep := routes[0]

	// Find best route

	for i := 1; i < len(routes); i++ {
		step := routes[i]
		if step.TotalCoveredDistance < bestStep.TotalCoveredDistance {
			bestStep = step
		}
	}

	//format route as an array of step
	steps := []Step{}
	for step := bestStep; step.Sale != nil; step = *step.PreviousStep {
		steps = append([]Step{step}, steps...)
	}
	return Route{steps}
}
