package sss

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMakeKeys(t *testing.T) {
	kMax := 8
	nMax := 15
	inMessage, k, n := makeParams(kMax, nMax)
	keys, prime := MakeKeys(inMessage, k, n)
	outMessage := Decode(keys, prime)

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

		if k < n && k != 0 && n != 0 {
			return
		}
	}
}
