package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/jpgoldberg/cryptopg/crackme"

	"os"
)

var withPwds = flag.Bool("p", false, "Output should contain all passwords")

func main() {
	flag.Parse()

	// wFilePath := "Resources/AgileWords.txt"

	challenges := getChallenges()
	for _, c := range challenges {
		c.FleshOut()
		_, err := c.DeriveKey()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't derive key: %v\n", err)
			os.Exit(1)
		}
	}

	for _, c := range challenges {
		if !*withPwds {
			if !c.IsSample {
				c.Pwd = ""
			}
		}
	}
	s, err := json.MarshalIndent(challenges, "", "\t")
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't marshal: %s\n", err)
	}

	fmt.Println(string(s))
}

func getChallenges() []*crackme.Challenge {
	/*
		file, err := os.Open("./data/secret/2017-11-01-secrets.json")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
	*/

	scanner := bufio.NewScanner(os.Stdin)
	var fText []byte
	for scanner.Scan() {
		fText = append(fText, scanner.Bytes()...)
	}

	result := new([]*crackme.Challenge)
	if err := json.Unmarshal(fText, result); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't unmarshal: %s\n", err)
	}
	return *result
}
