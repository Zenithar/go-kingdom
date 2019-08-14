package main

import (
	"math/rand"
	"time"

	"go.zenithar.org/kingdom/cli/kingdom/internal/cmd"
	"go.zenithar.org/pkg/log"

	_ "gocloud.dev/secrets/hashivault"
	_ "gocloud.dev/secrets/localsecrets"
)

// -----------------------------------------------------------------------------

func init() {
	// Set time locale
	time.Local = time.UTC

	// Initialize random seed
	rand.Seed(time.Now().UTC().Unix())
}

// -----------------------------------------------------------------------------

func main() {
	if err := cmd.Execute(); err != nil {
		log.CheckErr("Unable to complete command execution", err)
	}
}
