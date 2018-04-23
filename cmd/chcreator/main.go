package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/jpgoldberg/cryptopg/crackme"

	"os"
)

var withPwds = flag.Bool("p", false, "Output should contain all passwords")
var testKeys = flag.Bool("t", false, "Test whether input derived keys match calculated")

func main() {
	flag.Parse()

	badDkCnt := 0
	dkCnt := 0
	// wFilePath := "Resources/AgileWords.txt"

	challenges := getChallenges()

	for _, c := range challenges {
		c.FleshOut()

		if *testKeys {
			if c.Dk != nil {
				dkCnt++
				k := c.Dk
				c.DeriveKey()
				if bytes.Compare(k, c.Dk) != 0 {
					badDkCnt++
					fmt.Printf("Derived key mismatch for %s\n", c.ID)
				}
			}
		} else {
			_, err := c.DeriveKey()
			if err != nil {
				log.Printf("Couldn't derive key: %v\n", err)
			}
		}
	}

	if *testKeys {
		fmt.Printf("%d bad derived keys out of %d tested\n", badDkCnt, dkCnt)
		if badDkCnt > 0 {
			os.Exit(2)
		} else {
			os.Exit(0)
		}
	}
	// else if not just testing keys, we prepare output

	for _, c := range challenges {
		if !(*withPwds || c.IsSample) {
			c.Pwd = ""
		}
	}
	s, err := json.MarshalIndent(challenges, "", "\t")
	if err != nil {
		log.Printf("Can't marshal: %s\n", err)
	}

	fmt.Println(string(s))
	os.Exit(0)
}

func getChallenges() []*crackme.Challenge {
	scanner := bufio.NewScanner(os.Stdin)
	var fText []byte
	for scanner.Scan() {
		fText = append(fText, scanner.Bytes()...)
	}

	result := new([]*crackme.Challenge)
	if err := json.Unmarshal(fText, result); err != nil {
		log.Fatalf("Couldn't unmarshal: %s\n", err)
	}
	return *result
}
