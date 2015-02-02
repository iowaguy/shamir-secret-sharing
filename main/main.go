package main

import (
	"kn-threshold/sss"
	"log"
	"math/big"
	"os"
	"strings"
)

var (
	logger *log.Logger
)

func main() {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	args := os.Args

	if args[1] == "-k" || args[1] == "k" {
		// TODO error check for correct # of params
		data := parseInt(args[2])
		k := parseInt(args[3])
		n := parseInt(args[4])
		sss.MakeKeys(data, k, n)
	} else if args[1] == "-d" || args[1] == "d" {
		// TODO error check for correct # of params

		// parse prime in input
		//		prime := parseRat(args[2])

		// parse keys in input
		inKeys := args[2:]
		keys := make([]sss.Key, len(inKeys))
		for i := 0; i < len(keys); i++ {
			pair := strings.Split(inKeys[i], ":")

			index := parseRat(pair[0])
			d := parseRat(pair[1])

			keys[i].Xr = index
			keys[i].Yr = d
		}

		sss.Decode(keys)
	}

}

func parseInt(s string) *big.Int {
	i := new(big.Int)
	_, success := i.SetString(s, 10)
	if !success {
		logger.Fatal("SetString failed in: ParseInt")
	}

	return i
}

func parseRat(s string) *big.Rat {
	i := new(big.Rat)
	_, success := i.SetString(s)
	if !success {
		logger.Fatal("SetString failed in: parseRat")
	}

	return i
}
