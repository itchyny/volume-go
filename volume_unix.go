// +build linux freebsd netbsd

package volume

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func cmdEnv() []string {
	return []string{"LANG=C", "LC_ALL=C"}
}

func getVolumeCmd() []string {
	return []string{"pactl", "list", "sinks"}
}

var volumePattern = regexp.MustCompile(`\d+%`)

func parseVolume(out string) (int, error) {
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")
		if strings.HasPrefix(s, "Volume: 0:") {
			volumeStr := volumePattern.FindString(s)
			return strconv.Atoi(volumeStr[:len(volumeStr)-1])
		}
	}
	return 0, errors.New("no volume found")
}

func setVolumeCmd(volume int) []string {
	return []string{"pactl", "set-sink-volume", "0", strconv.Itoa(volume) + "%"}
}

func muteCmd() []string {
	return []string{"pactl", "set-sink-mute", "0", "1"}
}

func unmuteCmd() []string {
	return []string{"pactl", "set-sink-mute", "0", "0"}
}
