package helpers

import (
	"sync"

	"go.zenithar.org/butcher"
	"go.zenithar.org/butcher/hasher"
)

var (
	b      *butcher.Butcher
	once   sync.Once
	pepper = []byte(")[bP6W9U=uai:'&Wqu2d%SOy>3'?jed>")
)

func init() {
	once.Do(func() {
		var err error
		b, err = butcher.New(
			butcher.WithAlgorithm(hasher.Argon2id),
			butcher.WithPepper(pepper),
			butcher.WithSaltFunc(butcher.RandomNonce(32)),
		)
		if err != nil {
			panic(err)
		}
	})
}

// DerivePasswordFunc is used to derive password from secret
var DerivePasswordFunc = func(password string) (string, error) {
	// Derive password
	return b.Hash([]byte(password))
}

// CheckPasswordFunc is usedto check password hash
var CheckPasswordFunc = func(encoded, password string) (bool, error) {
	return b.Verify([]byte(encoded), []byte(password))
}

// SetPasswordPepperKey used to set the key of password encoding function
func SetPasswordPepperKey(key []byte) {
	if len(key) != 32 {
		panic("Password pepper hash key length must be 32bytes long.")
	}

	pepper = key

	// Reload hasher
	var err error
	b, err = butcher.New(
		butcher.WithAlgorithm(hasher.Argon2id),
		butcher.WithPepper(pepper),
		butcher.WithSaltFunc(butcher.RandomNonce(32)),
	)
	if err != nil {
		panic(err)
	}
}
