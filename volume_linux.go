// +build !windows,!darwin

package volume

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var useAmixer bool

func init() {
	if _, err := exec.LookPath("pactl"); err != nil {
		useAmixer = true
	}
	if _, err := exec.LookPath("pacmd"); err != nil {
		useAmixer = true
	}
}

func cmdEnv() []string {
	return []string{"LANG=C", "LC_ALL=C"}
}

func getVolumeCmd() []string {
	if useAmixer {
		return []string{"amixer", "get", "Master"}
	}
	return []string{"pacmd", "list-sinks"}
}

var volumePattern = regexp.MustCompile(`\d+%`)

func parseVolume(out string) (int, error) {
	lines := strings.Split(out, "\n")
	pa_default_sink_hooked := false

	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")

		if !useAmixer && strings.HasPrefix(s, "* index") {
			pa_default_sink_hooked = true
		}

		if useAmixer && strings.Contains(s, "Playback") && strings.Contains(s, "%") ||
			!useAmixer && pa_default_sink_hooked && strings.HasPrefix(s, "volume:") {
			volumeStr := volumePattern.FindString(s)
			return strconv.Atoi(volumeStr[:len(volumeStr)-1])
		}
	}
	return 0, errors.New("no volume found")
}

func setVolumeCmd(volume int) []string {
	if useAmixer {
		return []string{"amixer", "set", "Master", strconv.Itoa(volume) + "%"}
	}
	return []string{"pactl", "set-sink-volume", "@DEFAULT_SINK@", strconv.Itoa(volume) + "%"}
}

func increaseVolumeCmd(diff int) []string {
	var sign string
	if diff >= 0 {
		sign = "+"
	} else if useAmixer {
		diff = -diff
		sign = "-"
	}
	if useAmixer {
		return []string{"amixer", "set", "Master", strconv.Itoa(diff) + "%" + sign}
	}
	return []string{"pactl", "--", "set-sink-volume", "@DEFAULT_SINK@", sign + strconv.Itoa(diff) + "%"}
}

func getMutedCmd() []string {
	if useAmixer {
		return []string{"amixer", "get", "Master"}
	}
	return []string{"pactl", "list", "sinks"}
}

func parseMuted(out string) (bool, error) {
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")
		if useAmixer && strings.Contains(s, "Playback") && strings.Contains(s, "%") ||
			!useAmixer && strings.HasPrefix(s, "Mute: ") {
			if strings.Contains(s, "[off]") || strings.Contains(s, "yes") {
				return true, nil
			} else if strings.Contains(s, "[on]") || strings.Contains(s, "no") {
				return false, nil
			}
		}
	}
	return false, errors.New("no muted information found")
}

func muteCmd() []string {
	if useAmixer {
		return []string{"amixer", "-D", "pulse", "set", "Master", "mute"}
	}
	return []string{"pactl", "set-sink-mute", "@DEFAULT_SINK@", "1"}
}

func unmuteCmd() []string {
	if useAmixer {
		return []string{"amixer", "-D", "pulse", "set", "Master", "unmute"}
	}
	return []string{"pactl", "set-sink-mute", "@DEFAULT_SINK@", "0"}
}
