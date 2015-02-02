package sss

import (
	"log"
	"math/big"
)

var (
	logger *log.Logger
)

type Key struct {
	Xr, Yr *big.Rat
	Xi, Yi *big.Int
	K      int
}

func (k *Key) fillInts() {
	k.Xi = new(big.Int)
	k.Xi = k.Xr.Num()

	k.Yi = new(big.Int)
	k.Yi = k.Yr.Num()
}

func (k *Key) fillRats() {
	k.Xr = new(big.Rat)
	k.Xr.SetInt(k.Xi)

	k.Yr = new(big.Rat)
	k.Yr.SetInt(k.Yi)
}

func parseBigInt(s string) *big.Int {
	i := new(big.Int)
	_, success := i.SetString(s, 10)
	if !success {
		logger.Fatal("SetString failed in: ParseInt")
	}

	return i
}

func ParseRat(s string) *big.Rat {
	i := new(big.Rat)
	_, success := i.SetString(s)
	if !success {
		logger.Fatal("SetString failed in: parseRat")
	}

	return i
}
