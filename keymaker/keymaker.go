package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	data, _ := strconv.ParseInt(args[1], 10, 64)
	k, _ := strconv.ParseInt(args[2], 10, 64)
	n, _ := strconv.ParseInt(args[3], 10, 64)

	if k <= 0 || n <= 0 || k >= n {
		fmt.Println("Incorrect inputs")
	}

	var min int64 = 0
	if data > n {
		min = data
	} else {
		min = n
	}

	p := prime(min)

	// get k-1 random ints between [0,p)
	coeffs := kRand(k, p)
	coeffs[0] = data

	// calculate Ds
	// TODO randomly choose indices
	ds := make([]int64, n+1)
	tmpN := int(n)
	for i := 0; i < tmpN+1; i++ {
		ds[i] = poly(coeffs, i, p) // % p //WHY DOESNT THIS WORK!
	}

	for i := 1; i < len(ds); i++ {
		fmt.Printf("%d:%d\n", i, ds[i])
	}
}

// TODO choose a random prime greater than min
// find the next prime greater than min using the Miller-Rabin primality test
func prime(min int64) int64 {
	// find next prime larger than min
	for i := min + 1; true; i++ {
		bi := big.NewInt(i)

		// TODO: see if decoding the message is faster than manually verifying primality
		if bi.ProbablyPrime(10) { // && verifyPrimality(i) { // primality verification for large numbers is SLOOOOOWWWWW
			fmt.Printf("Prime: %d\n", i)
			return i
		}
	}

	return -1
}

func verifyPrimality(num int64) bool {
	for j := num / 2; j > 1; j-- {
		if num%j == 0 {
			return false
		}
	}

	return true
}

func kRand(k, p int64) (coeffs []int64) {
	coeffs = make([]int64, k)
	//	fmt.Printf("time now: %d", time.Now().Nanosecond())
	source := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(source)
	tmpK := int(k)
	for i := 1; i < tmpK; i++ {
		rando := r.Int63()
		coeffs[i] = rando % p
		//		fmt.Printf("rand[%d]: %d\n", i, rando)
	}

	return
}

func poly(coeffs []int64, x int, prime int64) (d int64) {
	for i := 0; i < len(coeffs); i++ {
		d += coeffs[i] * (int64(math.Pow(float64(x), float64(i))) % prime)
	}

	return
}
