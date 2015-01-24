package main

import (
	"fmt"
	//	matrix "github.com/skelterjohn/go.matrix"
	"log"
	//	"math"
	"os"
	"strconv"
	"strings"
)

var (
	logger *log.Logger
)

// type key struct {
// 	index int
// 	d     int
// }

func main() {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	prime, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		logger.Fatal("ParseInt conversion, prime")
	}

	args := os.Args[2:]
	keys := make([]struct{ x, y float64 }, len(args))
	for i := 0; i < len(keys); i++ {
		pair := strings.Split(args[i], ":")

		index, err := strconv.ParseFloat(pair[0], 64) // index
		if err != nil {
			logger.Fatal("ParseFloat conversion, index")
		}

		d, err := strconv.ParseFloat(pair[1], 64) // d
		if err != nil {
			logger.Fatal("ParseFloat conversion, d")
		}
		keys[i].x = index
		keys[i].y = d
		// keys[i] = key{
		// 	index,
		// 	d,
		// }
	}

	answer := int64(lagrange(keys, prime)) //% prime

	//	answer := linearAlgebra(keys)
	fmt.Printf("Prime: %d\n", prime)
	fmt.Printf("Before mod: %d\n", answer)
	fmt.Printf("Message is: %d\n", answer%prime)
}

// Determine value at x=0 by Lagrange Interpolation
func lagrange(keys []struct{ x, y float64 }, prime int64) (sum float64) {
	sum = 0
	k := len(keys)

	for j := 0; j < k; j++ {
		var product int64 = 1
		for m := 0; m < k; m++ {
			if m != j {
				product *= int64((0 - keys[m].x) / (keys[j].x - keys[m].x))
			}
		}
		fmt.Println(product * int64(keys[j].y))
		sum += float64(product * int64(keys[j].y) % prime)
	}

	return
}

// func linearAlgebra(keys []struct{ x, y int }) float64 {
// 	// make A
// 	mat := make([][]float64, len(keys))
// 	for row := range mat {
// 		mat[row] = make([]float64, len(keys))
// 	}

// 	for row := 0; row < len(mat); row++ {
// 		myKey := keys[row]
// 		for col := range mat {
// 			mat[row][col] = math.Pow(float64(myKey.x), float64(col))
// 		}
// 	}
// 	a := matrix.MakeDenseMatrixStacked(mat)

// 	// make B
// 	ds := make([]float64, len(keys))
// 	for i := 0; i < len(keys); i++ {
// 		ds[i] = float64(keys[i].y)
// 	}
// 	b := matrix.MakeDenseMatrix(ds, len(keys), 1)

// 	inverse, err := a.Inverse()
// 	if err != nil {
// 		logger.Fatal("Inverse error")
// 	}

// 	coeffs := matrix.Product(inverse, b).Array()

// 	return coeffs[0]
// }
