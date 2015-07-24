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

	k := len(keys)//strings.Count(keys[0], ":")-1
	logger.Printf("k:%d", k)
//	dataSegments, k := keyMetaData(keys[0])
	keyGrid := make([][]int64, k)
	xKeys := make([]int64, k)
	
	// split keys
	logger.Printf("xdim:%d", len(keyGrid))
	logger.Printf("ydim:%d", len(keyGrid[0]))
	for i, key := range keys {
		var index int64
		logger.Printf("i:%d", i)
		logger.Printf("key:%s", key)
		index, keyGrid[i] = parseKey(key, k)
		xKeys[i] = index
	}
	logger.Printf("grid: %v", keyGrid)		
	var buffer bytes.Buffer
	// decode each subKey array
	for i,_ := range keyGrid[0] {
		logger.Printf("i:%d", i)
		column := gridColumn(keyGrid, i)
		messageSegment := lagrange(k, xKeys, column, 0)
		logger.Printf("piece: %d", messageSegment)		
		buffer.WriteString(strconv.FormatInt(messageSegment, 10))
	}

	return buffer.String()
}

// returns a slice of the x value of the keys
// func xKeysSlice(grid [][]int64) (column []int64) {
// 	column = make([]int64, 0)
// 	for _, row := range grid {
// 		column = append(column, row[1])
// 	}
// 	return
// }

// returns the requested column of the keyGrid (columnIndex), corresponds to y-values 
// from a single message segment
func gridColumn(grid [][]int64, columnIndex int) (column []int64) {
	column = make([]int64, 0)
	for _, row := range grid {
		logger.Printf("%d, row[columnIndex]:%d", columnIndex, row[columnIndex])
		column = append(column, row[columnIndex])
	}
	logger.Printf("column: %v", column)
	return
}

// returns a slice of an individuals key values, including key index
func parseKey(fullKey string, k int) (index int64, keyTokens []int64) {
	subKeys := strings.Split(fullKey, ":")
	index, err := strconv.ParseInt(subKeys[1], 10, 64)
	if err != nil {
		logger.Fatal("Error occurred in ParseInt")
	}

	subKeys = subKeys[2:]

	keyTokens = make([]int64, len(subKeys))
	for i, subKey := range subKeys {
		num, err := strconv.ParseInt(subKey, 10, 64)
		if err != nil {
			logger.Fatal("Error occurred in ParseInt")
		}
		keyTokens[i] = num
	}

	return
}


// func DecodeColumn(k int, column []string) string {
// 	numColumn := make([]int64, len(column))
// 	for i, value := range column {
// 		tmp, err := strconv.ParseInt(value, 10, 64)
// 		numColumn[i] = tmp
// 		if err != nil {
// 			logger.Fatal("Error occurred in ParseInt")
// 		}
// 	}

// 	return strconv.FormatInt(lagrange(k, numColumn, 0), 10)
// }

// takes one key and returns the key threshold and the number of key columns
// func keyMetaData(key string) (numSplits int, k int) {
// 	tokens := strings.Split(key, ":")
// 	k, err := strconv.Atoi(tokens[0])
// 	if err != nil {
// 		logger.Fatal("Error occurred in Atoi")
// 	}
// 	// need to subtract 1 because the k is included in the key
// 	numSplits = len(tokens) - 2
// 	return
// }

// Lagrange Interpolation
// x is Lagrange Interpolate input i.e. f(x)
// for this threshold scheme, only zero will be passed in
func lagrange(k int, xKeys []int64, yKeys []int64, x int64) (sum int64) {
	if len(yKeys) < 1 {
		logger.Fatal("not enough keys")
	}

	sum = 0
	for j := 0; j < k; j++ {
		product := float64(1)
		for m := 0; m < k; m++ {
			if m != j {
				product = product * float64(x-xKeys[m]) / float64(xKeys[j]-xKeys[m])
				logger.Printf("product in loop:%f",product)
			}
		}
//		logger.Printf("preproduct:%f",product)
		product = product * float64(yKeys[j])
//		logger.Printf("postproduct:%f",product)
		sum += int64(product)
	}
	return sum
}
