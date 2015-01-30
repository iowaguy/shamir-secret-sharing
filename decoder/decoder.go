package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

var (
	logger *log.Logger
)

func main() {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)

	// parse prime in input
	prime := ParseRat(os.Args[1])

	// parse keys in input
	args := os.Args[2:]
	keys := make([]struct{ x, y *big.Rat }, len(args))
	for i := 0; i < len(keys); i++ {
		pair := strings.Split(args[i], ":")

		index := ParseRat(pair[0])
		d := ParseRat(pair[1])

		keys[i].x = index
		keys[i].y = d
	}

	answer := lagrange(keys, prime)

	fmt.Printf("Message is: %d\n", answer.Int64())
}

// Determine value at x=0 by Lagrange Interpolation
func lagrange(keys []struct{ x, y *big.Rat }, prime *big.Rat) *big.Int {
	sum := big.NewInt(0)
	k := len(keys)

	for j := 0; j < k; j++ {
		product := new(big.Rat)
		product.SetInt64(1)

		for m := 0; m < k; m++ {
			if m != j {

				temp := new(big.Rat)
				temp.Sub(big.NewRat(0, 1), keys[m].x)

				temp2 := new(big.Rat)
				temp2.Sub(keys[j].x, keys[m].x)

				temp3 := new(big.Rat)
				temp3.Quo(temp, temp2)

				temp4 := new(big.Rat)
				product.Set(temp4.Mul(product, temp3))
			}
		}

		fmt.Printf("key[j]: %v\n", keys[j].y)
		temp5 := new(big.Rat)
		product.Set(temp5.Mul(product, keys[j].y))

		fmt.Printf("product: %v\n", product)

		addMe := product.Num() //new(big.Int)

		sum.Add(sum, addMe)
		fmt.Printf("sum: %v\n", sum)
	}

	return sum
}

func ParseRat(s string) *big.Rat {
	i := new(big.Rat)
	_, success := i.SetString(s)
	if !success {
		logger.Fatal("SetString failed in: ParseRat")
	}

	return i
}
