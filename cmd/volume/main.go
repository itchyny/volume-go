package main

import (
	"fmt"
	"os"
)

const name = "volume"
const version = "0.0.0"
const author = "itchyny"

var revision = "HEAD"

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
		os.Exit(1)
	}
}
