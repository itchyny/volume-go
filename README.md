# volume-go [![CI Status][ci-badge]][ci-url] [![Go Report Card][goreportcard-badge]][goreportcard-url] [![MIT License][license-badge]][license-url] [![PkgGoDev][pkggodev-badge]][pkggodev-url]
### Volume control in Go
This is a Go language package for controlling audio volume.

## volume command
### Homebrew
```sh
brew install itchyny/tap/volume
```

### Build from source
```sh
go get github.com/itchyny/volume-go/cmd/volume
```

### Basic usage
```sh
 $ # Get the current audio volume.
 $ volume get 
20

 $ # Set the audio volume.
 $ volume set 40
 $ volume status
volume: 40
muted: false

 $ # Increase/decrease the audio volume.
 $ volume get
40
 $ volume up
 $ volume get
46

 $ volume down
 $ volume get
44

 $ # Get current audio volume status.
 $ volume status
volume: 20
muted: false

 $ # Mute and unmute.
 $ volume mute
 $ volume status
volume: 20
muted: true

 $ volume unmute
 $ volume status
volume: 20
muted: false
```

## Package usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/itchyny/volume-go"
)

func main() {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	fmt.Printf("current volume: %d\n", vol)

	err = volume.SetVolume(10)
	if err != nil {
		log.Fatalf("set volume failed: %+v", err)
	}
	fmt.Printf("set volume success\n")

	err = volume.Mute()
	if err != nil {
		log.Fatalf("mute failed: %+v", err)
	}

	err = volume.Unmute()
	if err != nil {
		log.Fatalf("unmute failed: %+v", err)
	}
}
```

## Bug Tracker
Report bug at [Issues・itchyny/volume-go - GitHub](https://github.com/itchyny/volume-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.

[ci-url]: https://github.com/itchyny/volume-go/actions
[ci-badge]: https://github.com/itchyny/volume-go/workflows/CI/badge.svg
[goreportcard-url]: https://goreportcard.com/report/github.com/itchyny/volume-go
[goreportcard-badge]: https://goreportcard.com/badge/github.com/itchyny/volume-go
[license-url]: https://github.com/itchyny/volume-go/blob/master/LICENSE
[license-badge]: http://img.shields.io/badge/license-MIT-blue.svg
[pkggodev-url]: https://pkg.go.dev/github.com/itchyny/volume-go
[pkggodev-badge]: https://pkg.go.dev/badge/github.com/itchyny/volume-go
