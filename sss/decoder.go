package sss

import (
	"bytes"
	"log"
	//	. "math/big"
	"os"
	"strconv"
	"strings"
)

func Decode(keys []string) string {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	if len(keys) == 0 {
		logger.Fatal("Need atleast one key")
	}

	// each row is a keySet
	dataSegments, k := keyMetaData(keys[0])
	keyGrid := make([][]int64, dataSegments)

	for i := range keys {
		keyGrid[i] = make([]int64, k)
	}

	// split keys
	for _, key := range keys {
		subKeys := strings.Split(key, ":")
		index, err := strconv.ParseInt(subKeys[1], 10, 64)
		if err != nil {
			logger.Fatal("Error occurred in ParseInt")
		}
		for j := 2; j < dataSegments; j++ {
			keyGrid[j-2][index], err = strconv.ParseInt(subKeys[j], 10, 64)
			if err != nil {
				logger.Fatal("Error occurred in ParseInt")
			}
		}
	}

	var buffer bytes.Buffer
	// decode each subKey array
	for _, row := range keyGrid {
		buffer.WriteString(strconv.FormatInt(lagrange(k, row, 0), 10))
	}

	return buffer.String()
}

func DecodeColumn(k int, column []string) string {
	numColumn := make([]int64, len(column))
	for i, value := range column {
		tmp, err := strconv.ParseInt(value, 10, 64)
		numColumn[i] = tmp
		if err != nil {
			logger.Fatal("Error occurred in ParseInt")
		}
	}

	return strconv.FormatInt(lagrange(k, numColumn, 0), 10)
}

func keyMetaData(key string) (numKeys int, k int) {
	tokens := strings.Split(key, ":")
	k, err := strconv.Atoi(tokens[0])
	if err != nil {
		logger.Fatal("Error occurred in Atoi")
	}
	// need to subtract 1 because the k is included in the key
	numKeys = len(tokens) - 2
	return
}

// Multi-precision Lagrange Interpolation
func lagrange(k int, keys []int64, n int) (sum int64) {
	if len(keys) < 1 {
		logger.Fatal("not enough keys")
	}

	//	k := keys[0].K
	//	x := NewRat(n, 1)
	sum = 0 //NewRat(0, 1)
	for j := 0; j < k; j++ {
		//		product := new(Rat)
		//		product.SetInt64(1)
		product := float64(1)
		for m := 0; m < k; m++ {
			if m != j {
				// TODO might be able to reuse some of these temps
				// temp := new(Rat)
				// temp.Sub(x, m)
				product = product * float64(n-m) / float64(j-m)

				// temp2 := new(Rat)
				// temp2.Sub(j, m)

				// temp3 := new(Rat)
				// temp3.Quo(temp, temp2)

				// temp4 := new(Rat)
				// temp4.Mul(product, temp3)

				// temp5 := new(Int)
				// temp5.Mod(product.Num(), prime)

				// product.Set(temp4)
			}
		}

		// temp6 := new(Rat)
		// temp6.Mul(product, keys[j].Yr)
		// product.Set(temp6)
		product = product * float64(keys[j])
		sum += int64(product)
		// sum.Add(sum, product)
	}

	// if sum.Denom().Cmp(NewInt(1)) != 0 {
	// 	logger.Fatalf("Something has gone terribly wrong! Denominator does not equal one, it is: %d\n", sum.Denom())
	// }
	// return sum.Num()

	return sum
}
