package sss

import (
	"log"
	. "math/big"
	"os"
)

func Decode(keys []Key) string {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	code := lagrange(keys, 0)

	return code.String()
}

// Multi-precision Lagrange Interpolation
func lagrange(keys []Key, n int64) *Int {
	if len(keys) < 1 {
		logger.Fatal("not enough keys")
	}

	k := keys[0].K
	x := NewRat(n, 1)
	sum := NewRat(0, 1)
	for j := 0; j < k; j++ {
		product := new(Rat)
		product.SetInt64(1)

		for m := 0; m < k; m++ {
			if m != j {
				// TODO might be able to reuse some of these temps
				temp := new(Rat)
				temp.Sub(x, keys[m].Xr)

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
		temp6.Mul(product, keys[j].Yr)
		product.Set(temp6)

		sum.Add(sum, product)
	}

	if sum.Denom().Cmp(NewInt(1)) != 0 {
		logger.Fatalf("Something has gone terribly wrong! Denominator does not equal one, it is: %d\n", sum.Denom())
	}
	return sum.Num()
}
