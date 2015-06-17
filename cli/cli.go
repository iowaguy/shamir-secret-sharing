package main

import (
	"fmt"
	"log"
	"os"
	"shamir-secret-sharing/sss"
	"strconv"
	//	"strings"
)

var (
	logger *log.Logger
)

func printUsage(name string) {
	fmt.Printf("Usage: %s [-k <data> <subset_size> <# of keys>] [-d <prime> <keys>... ]\n", name)
}

func main() {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	args := os.Args

	if len(args) < 3 || (args[1] != "-k" && args[1] != "-d") {
		printUsage(args[0])
		logger.Fatal("incorrect args")
	}

	if args[1] == "-k" {

		if len(args) != 5 {
			printUsage(args[0])
			logger.Fatal("incorrect number of args")
		}

		k, err := strconv.Atoi(args[3])
		if err != nil {
			printUsage(args[0])
			logger.Fatal("error in k Atoi")
		}

		n, err := strconv.Atoi(args[4])
		if err != nil {
			printUsage(args[0])
			logger.Fatal("error in n Atoi")
		}

		keys := sss.MakeKeys(args[2], true, k, n)
		for _, key := range keys {
			fmt.Println(key)
		}
	} else if args[1] == "-d" {
		// parse keys in input
		// inKeys := args[2:]
		// keys := make([]sss.Key, len(inKeys))
		// for i := 0; i < len(keys); i++ {
		// 	info := strings.Split(inKeys[i], ":")

		// 	k, err := strconv.Atoi(info[0])
		// 	if err != nil {
		// 		printUsage(args[0])
		// 		logger.Fatalf("error in k Atoi, decoder. k=%s\n", info[0])
		// 	}

		// 	prime := sss.ParseBigInt(info[1])
		// 	index := sss.ParseRat(info[2])
		// 	d := sss.ParseRat(info[3])

		// 	keys[i].Xr = index
		// 	keys[i].Yr = d
		// 	//			keys[i].Prime = prime
		// 	keys[i].K = k
		// 	keys[i].FillInts()
		// }

		// fmt.Printf("Message is: %s\n", sss.Decode(keys))

	}

}
