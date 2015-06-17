package sss

import (
	//	"fmt"
	"log"
	"math/big"
)

var (
	logger *log.Logger
)

type Key struct {
	K      int
	Xr, Yr *big.Rat
	Xi, Yi *big.Int
	//	Prime  *big.Int
}

func (k *Key) FillInts() {
	k.Xi = new(big.Int)
	k.Yi = new(big.Int)
	k.Xi, k.Yi = k.Xr.Num(), k.Yr.Num()
}

func (k *Key) FillRats() {
	k.Xr = new(big.Rat)
	k.Xr.SetInt(k.Xi)

	k.Yr = new(big.Rat)
	k.Yr.SetInt(k.Yi)
}

// func (k *Key) String() string {
// 	return fmt.Sprintf("%d:%d:%d:%d", k.K, k.Prime, k.Xi, k.Yi)
// }

///////////// For sorting //////////
type Keys []Key

func (k Keys) Len() int { return len(k) }

func (k Keys) Swap(i, j int) { k[i], k[j] = k[j], k[i] }

func (k Keys) Less(i, j int) bool { return k[i].Xi.Cmp(k[j].Xi) == -1 }

////////////////////////////////////

func ParseBigInt(s string) *big.Int {
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
