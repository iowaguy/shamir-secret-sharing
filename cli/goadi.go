package main

import (
	"fmt"
	"log"
	"os"
	"github.com/iowaguy/shamir-secret-sharing/sss"
	"strconv"
)

var (
	logger *log.Logger
)

func printUsage(name string) {
	fmt.Printf("Usage: %s [-k <data> <subset_size> <# of keys>] [-d <keys>... ]\n", name)
}

func main() {
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
	args := os.Args

	if len(args) < 3 || (args[1] != "-k" && args[1] != "-d") {
		printUsage(args[0])
		logger.Fatal("incorrect args")
	}

	// clean up using "flag" package
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

		keys := sss.MakeKeys(args[2], k, n)
		for _, key := range keys {
			fmt.Println(key)
		}
	} else if args[1] == "-d" {
		// parse keys in input
		inKeys := args[2:]
		fmt.Printf("Message is: %s\n", sss.Decode(inKeys))
	}

}
