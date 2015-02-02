package main

import (
	"kn-threshold/sss"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	logger *log.Logger
)

func main() {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	args := os.Args

	if len(args) < 3 {
		logger.Fatal("incorrect number of args")
	}

	if args[1] == "-k" || args[1] == "k" {
		if len(args) != 5 {
			logger.Fatal("incorrect number of args")
		}

		data, err := strconv.Atoi(args[2])
		if err != nil {
			logger.Fatal("error in data Atoi")
		}

		k, err := strconv.Atoi(args[3])
		if err != nil {
			logger.Fatal("error in k Atoi")
		}

		n, err := strconv.Atoi(args[4])
		if err != nil {
			logger.Fatal("error in n Atoi")
		}

		sss.MakeKeys(data, k, n)
	} else if args[1] == "-d" || args[1] == "d" {
		// parse prime in input
		//	prime := parseRat(args[2])

		// parse keys in input
		inKeys := args[2:]
		keys := make([]sss.Key, len(inKeys))
		for i := 0; i < len(keys); i++ {
			pair := strings.Split(inKeys[i], ":")

			index := utils.ParseRat(pair[0])
			d := utils.ParseRat(pair[1])

			keys[i].Xr = index
			keys[i].Yr = d
		}

		sss.Decode(keys)
	}

}
