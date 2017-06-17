package main

import "log"

var name = "volume"
var version = "v0.0.0"
var description = "control audio volume"
var author = "itchyny"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
