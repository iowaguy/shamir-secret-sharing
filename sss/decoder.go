package sss

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

func Decode(keys []string) string {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	if len(keys) == 0 {
		logger.Fatal("Need atleast one key")
	}

	k := len(keys) //strings.Count(keys[0], ":")-1
	keyGrid := make([][]int64, k)
	xKeys := make([]int64, k)

	// split keys
	for i, key := range keys {
		var index int64
		index, keyGrid[i] = parseKey(key, k)
		xKeys[i] = index
	}

	var buffer bytes.Buffer
	// decode each subKey array
	for i, _ := range keyGrid[0] {
		column := gridColumn(keyGrid, i)
		messageSegment := lagrange(k, xKeys, column, 0)
		buffer.WriteString(strconv.FormatInt(messageSegment, 10))
	}

	return buffer.String()
}

// returns the requested column of the keyGrid (columnIndex), corresponds to y-values
// from a single message segment
func gridColumn(grid [][]int64, columnIndex int) (column []int64) {
	column = make([]int64, 0)
	for _, row := range grid {
		column = append(column, row[columnIndex])
	}

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
			}
		}

		product = product * float64(yKeys[j])
		sum += int64(product)
	}
	return sum
}
