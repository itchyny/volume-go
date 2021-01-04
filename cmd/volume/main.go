package main

import (
	"fmt"
	"os"
)

const (
	name    = "volume"
	version = "0.2.1"
)

var revision = "HEAD"

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
		os.Exit(1)
	}
}
