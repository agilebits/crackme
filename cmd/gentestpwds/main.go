// gentest password creates wordlist passwords challenges.
package main

import (
	"bufio"
	"bytes"
	rand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
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

	gen, err := newGenerator(words)
	if err != nil {
		log.Fatalf("couldn't create generator: %v", err)
	}

	var challenges []crackme.Challenge
	for length := shortest; length <= longest; length++ {
		for i := 1; i <= numPerKind; i++ {
			pwd := gen.generate(length)
			c := new(crackme.Challenge)
			c.Pwd = pwd
			c.Hint = fmt.Sprintf("%d words", length)
			c.ID = crackme.MakeID(nil)

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
	listSize  uint32 // number of words on list
}

// newGenerator creates a password generator from a word list
func newGenerator(list []string) (*generator, error) {
	g := new(generator)
	g.Wordlist = list
	if len(g.Wordlist) > math.MaxUint32 {
		// Seriously. Who is going to feed in a wordlist this lone?
		return nil, fmt.Errorf("too many words (%d)", len(g.Wordlist))
	}
	g.listSize = uint32(len(g.Wordlist))
	if g.listSize == 0 {
		return nil, fmt.Errorf("empty word list")
	}
	g.rng = rand.Reader
	g.Separator = " "

	return g, nil
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
		index := int31n(g.listSize)
		pp += g.Wordlist[index]
	}
	return pp
}

// int31n returns, as an int32, a non-negative random number in [0,n) from a cryptographic appropriate source. It panics if n <= 0 or if
// a security-sensitive random number cannot be created. Care is taken to avoid modulo bias.
//
// Copied from the math/rand package..
func int31n(n uint32) uint32 {
	if n <= 0 {
		panic("invalid argument to int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return randomInt32() & (n - 1)
	}
	max := uint32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := randomInt32()
	for v > max {
		v = randomInt32()
	}
	return v % n
}

// randomInt32 creates a random 32 bit unsigned integer
func randomInt32() uint32 {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic("PRNG gen error:" + err.Error())
	}

	var result int32
	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.LittleEndian, &result)

	if err != nil {
		panic("PRNG conversion error:" + err.Error())
	}

	return uint32(result)
}
