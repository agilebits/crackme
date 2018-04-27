// gentest password creates wordlist passwords challenges.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/agilebits/spg"

	"github.com/agilebits/crackme"
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

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	wl, err := spg.NewWordList(words)
	if err != nil {
		log.Fatalf("Failed to create wordlist: %v", err)
	}

	var challenges []crackme.Challenge
	for length := shortest; length <= longest; length++ {
		r := spg.NewWLRecipe(length, wl)
		r.SeparatorChar = " "
		for i := 1; i <= numPerKind; i++ {
			pw, err := r.Generate()
			if err != nil {
				log.Fatalf("Failed to generate password: %v", err)
			}
			c := new(crackme.Challenge)
			c.Pwd = pw.String()
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
