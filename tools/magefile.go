// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Go mg.Namespace

var deps = []string{
	"github.com/izumin5210/gex/cmd/gex",
}

// Vendor create tools vendors
func (Go) Vendor() error {
	fmt.Println("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// Tools updates tools from package
func (Go) Tools() error {
	mg.SerialDeps(Go.Vendor)
	fmt.Println("## Intalling tools")
	return sh.RunV("go", "run", "github.com/izumin5210/gex/cmd/gex", "--build")
}
