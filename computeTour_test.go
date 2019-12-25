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
	availableSales := []Sale{saleA, saleB, saleC, saleD, saleE, saleF, saleG, saleH, saleI, saleJ}

	sales := make([]Sale, numberOfSales)
	for i := 0; i < numberOfSales; i++ {
		sales[i] = availableSales[i]
	}
	return sales
}

func TestGetBestRoute10(t *testing.T) {
	sales := getSales(10)
	route := getBestRoute(sales)

	expectedSalesOrder := []Sale{sales[2], sales[9], sales[0], sales[1], sales[8], sales[4], sales[3], sales[5], sales[7], sales[6]}

	for index, expectedSales := range expectedSalesOrder {
		assert.Equal(t, (*route.Steps[index].Sale).Name, expectedSales.Name, fmt.Sprintf("The sale %s position should be %d", expectedSales.Name, index))
	}

}

func benchmarkGetBestRoute(numberOfSales int, b *testing.B) {
	sales := getSales(numberOfSales)

	for n := 0; n < b.N; n++ {
		getBestRoute(sales)
	}
}

func BenchmarkGetBestRoute5(b *testing.B)  { benchmarkGetBestRoute(5, b) }
func BenchmarkGetBestRoute6(b *testing.B)  { benchmarkGetBestRoute(6, b) }
func BenchmarkGetBestRoute7(b *testing.B)  { benchmarkGetBestRoute(7, b) }
func BenchmarkGetBestRoute8(b *testing.B)  { benchmarkGetBestRoute(8, b) }
func BenchmarkGetBestRoute9(b *testing.B)  { benchmarkGetBestRoute(9, b) }
func BenchmarkGetBestRoute10(b *testing.B) { benchmarkGetBestRoute(10, b) }
