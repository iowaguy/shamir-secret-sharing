package main

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

func main() {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	args := os.Args
	data := ParseInt(args[1])

	k := ParseInt(args[2])
	n := ParseInt(args[3])

	// TODO fix this for math/big
	// if k <= 0 || n <= 0 || k >= n {
	// 	fmt.Println("Incorrect inputs")
	// }

	// TODO fix this for math/big
	// var min int64 = 0
	// if data > n {
	// 	min = data
	// } else {
	// 	min = n
	// }

	p := prime(data)

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

	for i := 1; i < len(ds); i++ {
		fmt.Printf("%d:%d\n", i, ds[i])
	}
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
			fmt.Printf("Prime: %v\n", i)
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
		//term.Set(term.Mod(term, prime))
		d.Set(d.Add(d, term))
	}

	return d
}

func ParseInt(s string) *Int {
	i := new(Int)
	_, success := i.SetString(s, 10)
	if !success {
		logger.Fatal("SetString failed in: ParseInt")
	}

	return i
}
