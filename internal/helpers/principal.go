package helpers

import (
	"encoding/base64"

	"golang.org/x/crypto/blake2b"
)

var principalHashKey = []byte(`,|g{w\vYzqyJFl0<R|T|:'(Zo."cyo]88w9@A]U*M(4>+W!rBXGJ&xRs8/mu/\O`)

// PrincipalHashFunc return the principal hashed using Blake2b keyed algorithm
var PrincipalHashFunc = func(principal string) string {
	// Prepare hasher
	hasher, err := blake2b.New512(principalHashKey)
	if err != nil {
		panic(err)
	}

	// Append principal
	_, err = hasher.Write([]byte(principal))
	if err != nil {
		panic(err)
	}

	// Return base64 hash value of the principal hash
	return base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
}

// SetPrincipalHashKey used to set the key of hash function
func SetPrincipalHashKey(key []byte) {
	if len(key) != 64 {
		panic("Principal hash key length must be 64bytes long.")
	}

	principalHashKey = key
}
