package sss

import (
	"fmt"
	"log"
	. "math/big"
	"os"
)

func Decode(keys []Key) {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	answer := lagrange(keys, 0)
	fmt.Printf("Message is: %v\n", answer)
}

// Multi-precision Lagrange Interpolation
func lagrange(keys []Key, n int64) *Int {
	sum := NewInt(0)
	k := len(keys)

	for j := 0; j < k; j++ {
		product := new(Rat)
		product.SetInt64(1)

		for m := 0; m < k; m++ {
			if m != j {

				temp := new(Rat)
				temp.Sub(NewRat(n, 1), keys[m].Xr)

				temp2 := new(Rat)
				temp2.Sub(keys[j].Xr, keys[m].Xr)

				temp3 := new(Rat)
				temp3.Quo(temp, temp2)

				temp4 := new(Rat)
				product.Set(temp4.Mul(product, temp3))
			}
		}

		temp5 := new(Rat)
		product.Set(temp5.Mul(product, keys[j].Yr))

		addMe := product.Num()
		sum.Add(sum, addMe)
	}

	return sum
}
