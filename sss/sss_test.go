package sss

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMakeKeys(t *testing.T) {
	inMessage := makeMessage()
	k := 6
	n := 11
	keys, prime := MakeKeys(inMessage, k, n)
	outMessage := Decode(keys, prime)

	if inMessage != outMessage {
		fmt.Printf("Input message: %s\n", inMessage)
		fmt.Printf("Output message: %s\n", outMessage)
		t.Fail()
	}
}

func makeMessage() string {
	source := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(source)
	randoCalrissian := r.Int63()
	return strconv.FormatInt(randoCalrissian, 10)
}
