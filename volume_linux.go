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
}

func cmdEnv() []string {
	return []string{"LANG=C", "LC_ALL=C"}
}

func getVolumeCmd() []string {
	if useAmixer {
		return []string{"amixer", "get", "Master"}
	}
	return []string{"pactl", "list", "sinks"}
}

func getPADefaultSink() (string, error) {
	out, err := execCmd([]string{"pactl", "info"})

	if err != nil {
		return "", err
	}

	lines := strings.Split(string(out), "\n")

	defaultSinkStr := "Default Sink: "
	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")
		if strings.HasPrefix(s, defaultSinkStr) {
			return strings.TrimSpace(strings.Replace(s, defaultSinkStr, "", 1)), nil
		}
	}
	return "", errors.New("Could not find PulseAudio Default Sink")
}

var volumePattern = regexp.MustCompile(`\d+%`)

func parseVolume(out string) (int, error) {
	sinkName, sinkNameErr := getPADefaultSink()

	paDefaultSinkHooked := false

	if sinkNameErr != nil {
		paDefaultSinkHooked = true
	}

	lines := strings.Split(out, "\n")

	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")

		if !useAmixer && strings.Contains(s, "Name: "+string(sinkName)) {
			paDefaultSinkHooked = true
		}

		if useAmixer && strings.Contains(s, "Playback") && strings.Contains(s, "%") ||
			!useAmixer && paDefaultSinkHooked && strings.HasPrefix(s, "Volume:") {
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
	sinkName, sinkNameErr := getPADefaultSink()

	paDefaultSinkHooked := false

	if sinkNameErr != nil {
		paDefaultSinkHooked = true
	}

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		s := strings.TrimLeft(line, " \t")

		if !useAmixer && strings.Contains(s, "Name: "+string(sinkName)) {
			paDefaultSinkHooked = true
		}

		if useAmixer && strings.Contains(s, "Playback") && strings.Contains(s, "%") ||
			!useAmixer && paDefaultSinkHooked && strings.HasPrefix(s, "Mute: ") {
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
