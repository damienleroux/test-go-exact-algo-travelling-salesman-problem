package main

import (
	"fmt"
	"time"
)

func main() {
	sales := loadFile("./mock/sales.json")
	start := time.Now()
	route := getBestRoute(sales)
	elapsed := time.Since(start)
	fmt.Printf("time to compute best route for %d sales is %s\n", len(sales), elapsed)
	fmt.Println("Tour:")
	for _, step := range route.Steps {
		fmt.Println(*step.Sale)
	}
}
