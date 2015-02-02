package sss

import (
	"fmt"
	"log"
	. "math/big"
	"os"
)

func Decode(keys []Key, prime *Int) string {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	code := lagrange(keys, 0, prime)
	fmt.Printf("Message is: %v\n", code)

	return code.String()
}

// Multi-precision Lagrange Interpolation
func lagrange(keys []Key, n int64, prime *Int) *Int {

	if len(keys) < 1 {
		logger.Fatal("not enough keys")
	}
	k := keys[0].K

	sum := NewInt(0)
	for j := 0; j < k; j++ {
		product := new(Rat)
		product.SetInt64(1)

		for m := 0; m < k; m++ {
			if m != j {
				// TODO might be able to reuse some of these temps
				temp := new(Rat)
				temp.Sub(NewRat(n, 1), keys[m].Xr)

				temp2 := new(Rat)
				temp2.Sub(keys[j].Xr, keys[m].Xr)

				temp3 := new(Rat)
				temp3.Quo(temp, temp2)

				temp4 := new(Rat)
				temp4.Mul(product, temp3)

				// temp5 := new(Int)
				// temp5.Mod(product.Num(), prime)

				product.Set(temp4)
			}
		}

		temp6 := new(Rat)
		product.Set(temp6.Mul(product, keys[j].Yr))

		addMe := product.Num()
		sum.Add(sum, addMe)
	}

	return sum
}
