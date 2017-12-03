// Package crackme is for creating cracking challenges
package crackme

import (
	rand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"os"

	"golang.org/x/crypto/pbkdf2"
	// "github.com/jpgoldberg/cryptopg/crackme"
)

// Challenge has details for each PBKDF2 challenges
type Challenge struct {
	ID       string `json:"id"`
	Hint     string `json:"hint,omitempty"`
	IsSample bool   `json:"sample,omitempty"`
	PRF      string `json:"prf"`
	Rounds   int    `json:"rounds"`
	SaltHex  string `json:"salt"`
	DkHex    string `json:"derived,omitempty"`
	Pwd      string `json:"pwd,omitempty"`

	KeyLen  int    `json:"-"`
	Salt    []byte `json:"-"`
	Dk      []byte `json:"-"`
	prfHash func() hash.Hash
}

const (
	DefaultSaltSize = 16
	DefaultKeySize  = 32
	DefaultMethod   = "HMAC-SHA256"
	DefaultRounds   = 100000
)

// String prints the Challenge without the password
func (c Challenge) String() string {
	if c.PRF == "" {
		return ""
	}
	r := fmt.Sprintf("PRF:\t%s\n", c.PRF)
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

	switch c.PRF {
	case "HMAC-SHA256":
		c.prfHash = sha256.New
	default:
		return nil, errors.New("unknown PRF")
	}

	if c.KeyLen > c.prfHash().Size() {
		fmt.Fprintf(os.Stderr, "PBKDF2 sucks at stretching: keylen %d > hash size %d", c.KeyLen, c.prfHash().Size())
	}

	c.Dk = pbkdf2.Key([]byte(c.Pwd), c.Salt, c.Rounds, c.KeyLen, c.prfHash)
	c.DkHex = hex.EncodeToString(c.Dk)
	return c.Dk, nil
}

// DeriveKey calculates key of default size using PBKDF2
func (c *Challenge) DeriveKey() ([]byte, error) {
	if c.KeyLen < 16 {
		c.KeyLen = DefaultKeySize
	}
	return c.DeriveKeyWithLength(c.KeyLen)
}

// FleshOut takes what exists within a challenge and fills in defaults and other
// fields based on what is there.
func (c *Challenge) FleshOut() {
	if c.Rounds == 0 {
		c.Rounds = DefaultRounds
	}
	if c.KeyLen == 0 {
		c.KeyLen = DefaultKeySize
	}
	if len(c.PRF) == 0 {
		c.PRF = DefaultMethod
	}

	switch {
	case c.Salt == nil && len(c.SaltHex) == 0:
		c.Salt = make([]byte, DefaultSaltSize)
		rand.Read(c.Salt)
		fallthrough
	case c.Salt != nil && len(c.SaltHex) == 0:
		c.SaltHex = hex.EncodeToString(c.Salt)
	case c.Salt == nil && len(c.SaltHex) > 0:
		c.Salt, _ = hex.DecodeString(c.SaltHex)
	}

	// unlike a missing salt, we do not derive a key until explicitly told to
	switch {
	case c.Dk != nil && len(c.DkHex) == 0:
		c.SaltHex = hex.EncodeToString(c.Dk)
	case c.Dk == nil && len(c.DkHex) > 0:
		c.Dk, _ = hex.DecodeString(c.DkHex)
	}

	if len(c.ID) == 0 {
		// create an ID from the salt. (Let's hope all salts are unique)
		// By taking a multiple of three bytes, we get base64 encoding without padding
		c.ID = base64.StdEncoding.EncodeToString(c.Salt[0:9])
	}
}
