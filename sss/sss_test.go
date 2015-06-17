package sss

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// End to end test, including unsorted, non-consecutive keys
func TestEnd2End(t *testing.T) {
	kMax := 8
	nMax := 15

	inMessage, k, n := makeParams(kMax, nMax)
	keys := MakeKeys(inMessage, k, n)
	subset := chooseRandomKeys(keys)
	outMessage := Decode(subset)

	if inMessage != outMessage {
		fmt.Printf("Input message: %s\n", inMessage)
		fmt.Printf("Output message: %s\n", outMessage)
		t.Fail()
	}
}

func TestColumn(t *testing.T) {
	kMax := 8
	nMax := 15

	inMessage, k, n := makeParams(kMax, nMax)
	keys := MakeKeys(inMessage, k, n)
	//subset := chooseRandomKeys(keys)
	outMessage := Decode(subset)

	if inMessage != outMessage {
		fmt.Printf("Input message: %s\n", inMessage)
		fmt.Printf("Output message: %s\n", outMessage)
		t.Fail()
	}

}

func makeParams(kMax, nMax int) (message string, k, n int) {
	source := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(source)
	randoCalrissian := r.Int63()
	message = strconv.FormatInt(randoCalrissian, 10)
	for {
		k = r.Int() % kMax
		n = r.Int() % nMax

		if k < n && k > 0 && n > 0 {
			return
		}
	}
}

func chooseLargestKeys(keys []Key) []Key {
	subset := make([]Key, keys[0].K)
	for i := 0; i < keys[0].K; i++ {
		subset[i] = keys[len(keys)-keys[0].K+i]
	}

	return subset
}

func chooseLargestKeysBackwardsSkipOne(keys []Key) []Key {
	subset := make([]Key, keys[0].K)
	j := 0
	for i := 0; i < keys[0].K; i++ {
		if i == 1 {
			j++
		}
		subset[i] = keys[len(keys)-1-j]
		j++
	}

	return subset
}

func chooseRandomKeys(keys []Key) []Key {
	source := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(source)

	n := len(keys)

	subset := make([]Key, keys[0].K)
	used := make([]bool, n)
	for i := 0; i < keys[0].K; i++ {
		rando := r.Int()
		index := rando % n

		if used[index] {
			i -= 1
		} else {
			used[index] = true
			subset[i] = keys[index]
		}
	}

	return subset
}
