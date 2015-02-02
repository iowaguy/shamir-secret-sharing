package sss

import (
	"math/big"
)

type Key struct {
	Xr, Yr *big.Rat
	Xi, Yi *big.Int
}

// type converter interface {
// 	fillInts()
// 	fillRats()
// }

func (k Key) fillInts() {
	k.Xi = new(big.Int)
	k.Xi = k.Xr.Num()

	k.Yi = new(big.Int)
	k.Yi = k.Yr.Num()
}

func (k Key) fillRats() {
	k.Xr = new(big.Rat)
	k.Xr.SetInt(k.Xi)

	k.Yr = new(big.Rat)
	k.Yr.SetInt(k.Yi)
}
