package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/itchyny/volume-go"
)

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("no arg")
	}
	switch args[0] {
	case "-v", "version", "-version", "--version":
		return printVersion()
	case "-h", "help", "-help", "--help":
		return printHelp()
	case "status":
		if len(args) == 1 {
			return printStatus()
		}
	case "get":
		if len(args) == 1 {
			return getVolume()
		}
	case "set":
		if len(args) == 2 {
			return setVolume(args[1])
		}
	case "up":
		if len(args) == 1 {
			return upVolume("6")
		} else if len(args) == 2 {
			return upVolume(args[1])
		}
	case "down":
		if len(args) == 1 {
			return downVolume("6")
		} else if len(args) == 2 {
			return downVolume(args[1])
		}
	case "mute":
		if len(args) == 1 {
			return volume.Mute()
		}
	case "unmute":
		if len(args) == 1 {
			return volume.Unmute()
		}
	}
	return fmt.Errorf("invalid argument for volume: %+v", args)
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

func upVolume(diffStr string) error {
	diff, err := strconv.Atoi(diffStr)
	if err != nil {
		return err
	}
	return volume.IncreaseVolume(diff)
}

func downVolume(diffStr string) error {
	diff, err := strconv.Atoi(diffStr)
	if err != nil {
		return err
	}
	return volume.IncreaseVolume(-diff)
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
