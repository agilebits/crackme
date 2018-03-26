// gentest password creates wordlist passwords challenges.
package main

import (
	"bufio"
	rand "crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/jpgoldberg/cryptopg/crackme"
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

	_ = numPerKind
	_ = shortest
	_ = longest

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

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	gen := newGeneratorFromList(words)

	var challenges []crackme.Challenge
	for length := shortest; length <= longest; length++ {
		for i := 1; i <= numPerKind; i++ {
			pwd := gen.generate(length)
			c := new(crackme.Challenge)
			c.Pwd = pwd
			c.Hint = fmt.Sprintf("%d words", length)
			c.FleshOut()
			c.DeriveKey()
			challenges = append(challenges, *c)
		}
	}
	s, err := json.MarshalIndent(challenges, "", "\t")
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't marshal: %s\n", err)
		os.Exit(3)
	}

	fmt.Println(string(s))
	os.Exit(0)
}

type generator struct {
	Wordlist  []string
	rng       io.Reader
	Separator string
	nw        int64
	bigNw     *big.Int
}

func newGeneratorFromList(list []string) *generator {
	g := new(generator)
	g.Wordlist = list
	g.nw = int64(len(g.Wordlist))
	if g.nw == 0 {
		fmt.Fprintln(os.Stderr, "empty word list")
		return nil
	}
	g.bigNw = big.NewInt(g.nw)
	g.rng = rand.Reader
	g.Separator = " "

	return g
}

func (g *generator) generate(n int) string {
	if n < 1 {
		log.Panic("number of words requested should be > 0")
	}

	pp := ""
	for i := 1; i <= n; i++ {
		if i > 1 {
			pp += g.Separator
		}
		bigIndex, _ := rand.Int(g.rng, g.bigNw)
		index := bigIndex.Uint64()
		pp += g.Wordlist[index]
	}
	return pp
}
