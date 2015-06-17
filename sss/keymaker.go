package sss

import (
	"bytes"
	"log"
	"math"
	. "math/big"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func MakeKeys(data string, k, n int) (keys []string) {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)

	substrs, _ := splitter(data, 8)

	keys = make([]string, n)
	keyGrid := make([][]string, n)
	for i := 0; i < n; i++ {
		keyGrid[i] = make([]string, len(substrs))
	}

	for i, subSecret := range substrs {
		kSet := keySet(subSecret, k, n)
		for j, subKey := range kSet {
			keyGrid[j][i] = subKey
		}
	}

	return combiner(keyGrid, k)
}

func splitter(dStr string, sliceSize int) (substrs []string, pieces int) {
	size := len(dStr) / sliceSize
	if len(dStr)%sliceSize != 0 {
		size += 1
	}
	substrs = make([]string, size)
	pieces = len(substrs)
	begin := 0
	for i := range substrs {
		var end int
		if len(dStr)-begin < sliceSize {
			end = len(dStr)
		} else {
			end = begin + sliceSize
		}

		substrs[i] = dStr[begin:end]

		begin += sliceSize
	}

	return
}

func keySet(dStr string, kIn, nIn int) (column []string) {
	data := ParseBigInt(dStr)
	k := NewInt(int64(kIn))
	n := NewInt(int64(nIn))

	zero := NewInt(0)
	if k.Cmp(zero) <= 0 || n.Cmp(zero) <= 0 || k.Cmp(n) >= 0 {
		logger.Fatalf("Incorrect inputs:\nk=%d\nn=%d\n", k, n)
	}

	min := new(Int)
	if data.Cmp(n) == 1 {
		min.Set(data)
	} else {
		min.Set(n)
	}

	p := new(Int)
	p.Set(min)
	for {
		// need to reset p to avoid infinite loop for primality false positives
		p.Set(prime(p))

		// get k-1 random ints between [0,p)
		coeffs := kRand(k, p)
		coeffs[0] = data

		// calculate Ds
		column = make([]string, nIn)
		for i := 1; i <= nIn; i++ {
			column[i-1] = poly(coeffs, i, p).String()
		}

		// Verifying primality by making sure keys work
		//if dStr == DecodeColumn(kIn, column) {
		return column
		//}

	}
}

// TODO choose a random prime greater than min
// find the next prime greater than min using the Miller-Rabin primality test
func prime(min *Int) *Int {
	// find next prime larger than min
	one := NewInt(1)
	i := new(Int)

	for i.Add(min, one); true; i.Set(i.Add(i, one)) {
		if i.ProbablyPrime(10) {
			return i
		}
	}

	return NewInt(-1)
}

func kRand(k, p *Int) (coeffs []*Int) {
	coeffs = make([]*Int, k.Int64())

	source := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(source)

	tmpK := int(k.Int64())
	for i := 1; i < tmpK; i++ {
		rando := new(Int)
		rando.Rand(r, p)
		coeffs[i] = rando
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

func combiner(keyGrid [][]string, k int) (keys []string) {
	keys = make([]string, len(keyGrid))
	for i, row := range keyGrid {
		var buffer bytes.Buffer
		buffer.WriteString(strconv.Itoa(k))
		buffer.WriteString(":")
		buffer.WriteString(strconv.Itoa(i + 1))
		for _, cell := range row {
			buffer.WriteString(":")
			buffer.WriteString(cell)
		}
		keys[i] = buffer.String()
	}

	return
}
