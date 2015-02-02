package sss

import (
	"fmt"
	"log"
	"math"
	. "math/big"
	"math/rand"
	"os"
	"time"
)

var (
	logger *log.Logger
)

func MakeKeys(data, k, n *Int) []Key {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)

	zero := NewInt(0)
	if k.Cmp(zero) <= 0 || n.Cmp(zero) <= 0 || k.Cmp(n) >= 0 {
		logger.Fatal("Incorrect inputs")
	}

	min := new(Int)
	if data.Cmp(n) == 1 {
		min.Set(data)
	} else {
		min.Set(n)
	}

	p := prime(min)

	// get k-1 random ints between [0,p)
	coeffs := kRand(k, p)
	coeffs[0] = data

	// calculate Ds
	// TODO randomly choose indices
	ds := make([]*Int, n.Int64()+1)
	tmpN := int(n.Int64())
	for i := 0; i < tmpN+1; i++ {
		ds[i] = poly(coeffs, i, p)
	}

	keys := make([]Key, len(ds)-1)
	for i := 1; i < len(ds); i++ {
		keys[i-1].Xi = NewInt(int64(i))
		keys[i-1].Yi = ds[i]
		keys[i-1].fillRats()
		fmt.Printf("%d:%d\n", i, ds[i])
	}

	return keys
}

// TODO choose a random prime greater than min
// find the next prime greater than min using the Miller-Rabin primality test
func prime(min *Int) *Int {
	// find next prime larger than min
	one := NewInt(1)
	i := new(Int)

	for i.Add(min, one); true; i.Set(i.Add(i, one)) {
		//		bi := NewInt(i)

		// TODO: see if decoding the message is faster than manually verifying primality
		if i.ProbablyPrime(10) { // && verifyPrimality(i) { // primality verification for large numbers is SLOOOOOWWWWW
			//			fmt.Printf("Prime: %v\n", i)
			return i
		}
	}

	return NewInt(-1)
}

func verifyPrimality(num int64) bool {
	for j := num / 2; j > 1; j-- {
		if num%j == 0 {
			return false
		}
	}

	return true
}

func kRand(k, p *Int) (coeffs []*Int) {
	coeffs = make([]*Int, k.Int64())

	source := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(source)

	tmpK := int(k.Int64())
	for i := 1; i < tmpK; i++ {
		rando := new(Int)
		rando.Rand(r, p)
		coeffs[i] = rando // % p
	}

	return
}

func poly(coeffs []*Int, x int, prime *Int) *Int {
	d := NewInt(0)
	max := len(coeffs)
	for i := 0; i < max; i++ {
		term := new(Int)
		term.Mul(coeffs[i], NewInt(int64(math.Pow(float64(x), float64(i)))))
		d.Set(d.Add(d, term))
	}

	return d
}
