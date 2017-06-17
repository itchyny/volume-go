package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/itchyny/volume-go"
)

func run() error {
	if len(os.Args) <= 1 {
		return errors.New("no arg")
	}
	switch os.Args[1] {
	case "-v", "version", "-version", "--version":
		return printVersion()
	case "-h", "help", "-help", "--help":
		return printHelp()
	case "status":
		if len(os.Args) == 2 {
			return printStatus()
		}
	case "get":
		if len(os.Args) == 2 {
			return getVolume()
		}
	case "set":
		if len(os.Args) == 3 {
			return setVolume(os.Args[2])
		}
	case "mute":
		if len(os.Args) == 2 {
			return volume.Mute()
		}
	case "unmute":
		if len(os.Args) == 2 {
			return volume.Unmute()
		}
	}
	return fmt.Errorf("invalid argument for volume: %+v", os.Args[1:])
}

func printStatus() error {
	vol, err := volume.GetVolume()
	if err != nil {
		return err
	}
	muted, err := volume.GetMuted()
	if err != nil {
		return err
	}
	fmt.Printf("volume: %d\n", vol)
	fmt.Printf("muted: %t\n", muted)
	return nil
}

func getVolume() error {
	vol, err := volume.GetVolume()
	if err != nil {
		return err
	}
	fmt.Println(vol)
	return nil
}

func setVolume(volStr string) error {
	vol, err := strconv.Atoi(volStr)
	if err != nil {
		return err
	}
	return volume.SetVolume(vol)
}

func printVersion() error {
	fmt.Printf("%s version %s\n", name, version)
	return nil
}

func printHelp() error {
	fmt.Printf(strings.Replace(`NAME:
   $NAME - %s

USAGE:
   $NAME command [argument...]

COMMANDS:
   status     prints the volume status
   get        prints the current volume
   set [vol]  sets the audio volume
   mute       mutes the audio
   unmute     unmutes the audio

VERSION:
   %s

AUTHOR:
   %s
`, "$NAME", name, -1), description, version, author)
	return nil
}
