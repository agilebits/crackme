package crackme

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

// TestVector is for challenges with Expected values
type TestVector struct {
	Challenge
	Expected string
}

var set5 = TestVector{
	Challenge: Challenge{
		Rounds:  4096,
		KeyLen:  40,
		Salt:    []byte("saltSALTsaltSALTsaltSALTsaltSALTsalt"),
		PRF:     "HMAC-SHA256",
		Pwd:     "passwordPASSWORDpassword",
		Hint:    "Set 5 https://github.com/ircmaxell/quality-checker/blob/master/tmp/gh_18/PHP-PasswordLib-master/test/Data/Vectors/pbkdf2-draft-josefsson-sha256.test-vectors",
		prfHash: sha256.New,
		Dk:      nil,
	},
	// 348c89dbcbd32b2f32d814b8116e84cf2b17347ebc1800181c4e2a1fb8dd53e1c635518c7dac47e9
	// From http://stackoverflow.com/a/5136918/1304076
	Expected: "348c89dbcbd32b2f32d814b8116e84cf2b17347ebc1800181c4e2a1fb8dd53e1c635518c7dac47e9",
}

// Pass checks of Dk is Expected
func TestPBKDF2(t *testing.T) {
	if set5.Dk == nil {
		set5.DeriveKey()
	}
	e, _ := hex.DecodeString(set5.Expected)
	if bytes.Compare(set5.Dk, e) != 0 {
		t.Error("Didn't derive expected key")
	}
}

// String for test vector challenge
func (tvec TestVector) String() string {
	if tvec.Dk == nil {
		tvec.DeriveKey()
	}
	r := fmt.Sprintf("Passwd:\t\"%s\"\n", tvec.Pwd)
	r += tvec.Challenge.String()
	r += fmt.Sprintf("Expect:\t%s\n", tvec.Expected)

	return r
}
