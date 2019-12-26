package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSales(numberOfSales int) []Sale {
	saleA := Sale{"A", 48.87573090445785, 2.3293433555529797}
	saleB := Sale{"B", 48.8742208345030355, 2.332755125419923}
	saleC := Sale{"C", 48.873529291923894, 2.327412165062257}
	saleD := Sale{"D", 48.871567517076905, 2.3280773528979504}
	saleE := Sale{"E", 48.87172276824897, 2.3308453926013195}
	saleF := Sale{"F", 48.872047382779314, 2.324987448112794}
	saleG := Sale{"G", 48.87474301338154, 2.320245302574464}
	saleH := Sale{"H", 48.87433373850616, 2.3230991729663097}
	saleI := Sale{"I", 48.87409381719493, 2.3328624137805187}
	saleJ := Sale{"J", 48.87413615633347, 2.327691114799806}
	saleK := Sale{"K", 48.87413615634347, 2.327691114798806}
	availableSales := []Sale{saleA, saleB, saleC, saleD, saleE, saleF, saleG, saleH, saleI, saleJ, saleK}

	sales := make([]Sale, numberOfSales)
	for i := 0; i < numberOfSales; i++ {
		sales[i] = availableSales[i]
	}
	return sales
}

func TestGetBestRoute10(t *testing.T) {
	sales := getSales(10)
	route := GetBestRoute(sales, 100)

	expectedSalesOrder := []Sale{sales[2], sales[9], sales[0], sales[1], sales[8], sales[4], sales[3], sales[5], sales[7], sales[6]}

	for index, expectedSales := range expectedSalesOrder {
		assert.Equal(t, (*route.Steps[index].Sale).Name, expectedSales.Name, fmt.Sprintf("The sale %s position should be %d", expectedSales.Name, index))
	}

}

func benchmarkGetBestRoute(numberOfSales int, b *testing.B, multithread bool) {
	sales := getSales(numberOfSales)

	maxGoroutines := 100
	if !multithread {
		maxGoroutines = 0
	}
	for n := 0; n < b.N; n++ {
		GetBestRoute(sales, maxGoroutines)
	}
}

func BenchmarkGetBestRoute5(b *testing.B)             { benchmarkGetBestRoute(5, b, false) }
func BenchmarkGetBestRoute6(b *testing.B)             { benchmarkGetBestRoute(6, b, false) }
func BenchmarkGetBestRoute7(b *testing.B)             { benchmarkGetBestRoute(7, b, false) }
func BenchmarkGetBestRoute8(b *testing.B)             { benchmarkGetBestRoute(8, b, false) }
func BenchmarkGetBestRoute9(b *testing.B)             { benchmarkGetBestRoute(9, b, false) }
func BenchmarkGetBestRoute10(b *testing.B)            { benchmarkGetBestRoute(10, b, false) }
func BenchmarkGetBestRoute11(b *testing.B)            { benchmarkGetBestRoute(11, b, false) }
func BenchmarkGetBestRouteMultithread5(b *testing.B)  { benchmarkGetBestRoute(5, b, true) }
func BenchmarkGetBestRoutMultithreade6(b *testing.B)  { benchmarkGetBestRoute(6, b, true) }
func BenchmarkGetBestRouteMultithread7(b *testing.B)  { benchmarkGetBestRoute(7, b, true) }
func BenchmarkGetBestRouteMultithread8(b *testing.B)  { benchmarkGetBestRoute(8, b, true) }
func BenchmarkGetBestRouteMultithread9(b *testing.B)  { benchmarkGetBestRoute(9, b, true) }
func BenchmarkGetBestRouteMultithread10(b *testing.B) { benchmarkGetBestRoute(10, b, true) }
func BenchmarkGetBestRouteMultithread11(b *testing.B) { benchmarkGetBestRoute(11, b, true) }
