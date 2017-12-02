// Package crackme is for creating cracking challenges
package crackme

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"

	"golang.org/x/crypto/pbkdf2"
	// "github.com/jpgoldberg/cryptopg/crackme"
)

// Challenge has details for each PBKDF2 challenges
type Challenge struct {
	Rounds  int              `json:"rounds"`
	KeyLen  int              `json:"-"`
	Salt    []byte           `json:"-"`
	SaltB64 string           `json:"salt"`
	Dk      []byte           `json:"-"`
	DkB64   string           `json:"derived,omitempty"`
	Method  string           `json:"method"`
	Pwd     string           `json:"pwd,omitempty"`
	Hint    string           `json:"hint"`
	ID      string           `json:"ID"`
	prfHash func() hash.Hash `json:"-"`
}

// String prints the Challenge without the password
func (c Challenge) String() string {
	if c.Method == "" {
		return ""
	}
	r := fmt.Sprintf("PRF:\t%s\n", c.Method)
	r += fmt.Sprintf("Rounds:\t%d\n", c.Rounds)
	if c.Hint != "" {
		r += fmt.Sprintf("Hint:\t%s\n", c.Hint)
	}
	if c.Salt != nil {
		r += fmt.Sprintf("Salt:\t%s\n", hex.EncodeToString(c.Salt))
	}
	if c.Dk != nil {
		r += fmt.Sprintf("DKey:\t%s\n", hex.EncodeToString(c.Dk))
	}
	return r
}

// DeriveKeyWithLength calculates the key of size bytes using PBKDF2
func (c *Challenge) DeriveKeyWithLength(size int) ([]byte, error) {
	c.KeyLen = size

	switch c.Method {
	case "HMAC-SHA256":
		c.prfHash = sha256.New
	default:
		return nil, errors.New("unknown PRF")
	}

	c.Dk = pbkdf2.Key([]byte(c.Pwd), c.Salt, c.Rounds, c.KeyLen, c.prfHash)
	return c.Dk, nil
}

// DeriveKey calculates key of default size using PBKDF2
func (c *Challenge) DeriveKey() ([]byte, error) {
	length := 32
	if c.KeyLen == 0 {
		c.KeyLen = length
	}
	return c.DeriveKeyWithLength(c.KeyLen)
}
