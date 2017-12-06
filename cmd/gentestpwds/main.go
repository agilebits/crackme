package main

import (
	"bufio"
	rand "crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	numPerKindPtr := flag.Int("n", 3, "number to generate for each length")
	shortestPtr := flag.Int("s", 2, "number of words in shortest password")
	longestPtr := flag.Int("l", 4, "number of words in longest password")

	flag.Parse()

	// Let's let the compiler do a little more work so that
	// I don't have to have ugly stars or variable names hanging
	// around.
	numPerKind := *numPerKindPtr
	shortest := *shortestPtr
	longest := *longestPtr

	var file *os.File
	var err error

	switch flag.NArg() {
	case 0:
		file = os.Stdin
	case 1:
		file, err = os.Open(flag.Arg(0))
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't open \"%s\": %s", flag.Arg(0), err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Specify one file name or read standard input\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	// Let's just see that I'm parsing these right
	fmt.Printf("n == %d; s == %d; l == %d\n", numPerKind, shortest, longest)
	if flag.NArg() == 1 {
		fmt.Printf("opened <%s> for reading\n", flag.Arg(0))
	} else {
		fmt.Println("reading from standard input")

	}

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	// just some testing
	fmt.Printf("The %d-th word is \"%s\"\n", 123+1, words[123])
}

func genPassphrase(n int, list []string) string {
	nw := int64(len(list))
	if nw == 0 {
		log.Panic("empty word list")
	}
	if n < 1 {
		log.Panic("number of words requested should be > 0")
	}

	bigNw := big.NewInt(nw)
	rng := rand.Reader

	const sep = " "
	pp := ""
	for i := 1; i <= n; i++ {
		if i > 1 {
			pp += sep
		}
		bigIndex, _ := rand.Int(rng, bigNw)
		index := bigIndex.Uint64()
		pp += list[index]
	}
	return pp
}
