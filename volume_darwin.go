package volume

import (
	"strconv"
	"strings"
)

func cmdEnv() []string {
	return nil
}

func getVolumeCmd() []string {
	return []string{"osascript", "-e", "output volume of (get volume settings)"}
}

func parseVolume(out string) (int, error) {
	return strconv.Atoi(strings.TrimSuffix(out, "\n"))
}

func setVolumeCmd(volume int) []string {
	return []string{"osascript", "-e", "set volume output volume " + strconv.Itoa(volume)}
}

func muteCmd() []string {
	return []string{"osascript", "-e", "set volume output muted true"}
}

func unmuteCmd() []string {
	return []string{"osascript", "-e", "set volume output muted false"}
}
