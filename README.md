# volume-go
### Volume control in Go
This is a Go language package for controlling audio volume.

## Installation
```sh
 $ go get -u github.com/itchyny/volume-go/cmd/volume
 $ volume get
25
 $ volume set 20
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
		log.Fatal("get volume failed: %+v", err)
	}
	fmt.Printf("current volume: %d\n", vol)

	err = volume.SetVolume(10)
	if err != nil {
		log.Fatal("set volume failed: %+v", err)
	}
	fmt.Printf("set volume success\n")

	err = volume.Mute()
	if err != nil {
		log.Fatal("mute failed: %+v", err)
	}

	err = volume.Unmute()
	if err != nil {
		log.Fatal("unmute failed: %+v", err)
	}
}
```

## Bug Tracker
Report bug at [Issues・itchyny/volume-go - GitHub](https://github.com/itchyny/volume-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.
