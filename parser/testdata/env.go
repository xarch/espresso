package testdata

import (
	"go/build"
)

// Env represents decoded "env.espresso" file.
var Env = build.Default.GOPATH
