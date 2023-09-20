package main

import (
	"os"

	"github.com/katallaxie/acl/runtime"
)

func main() {
	rules, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	interp, err := runtime.FromString(string(rules))
	if err != nil {
		panic(err)
	}

	interp.HasAccess(runtime.Context{}, "/foo/bar")
}
