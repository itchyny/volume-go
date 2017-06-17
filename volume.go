package volume

import (
	"errors"
	"os"
	"os/exec"
)

func execCmd(cmdArgs []string) ([]byte, error) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Env = append(os.Environ(), cmdEnv()...)
	return cmd.Output()
}

// GetVolume returns the current volume (0 to 100).
func GetVolume() (int, error) {
	out, err := execCmd(getVolumeCmd())
	if err != nil {
		return 0, err
	}
	return parseVolume(string(out))
}

// SetVolume sets the sound volume to the specified value.
func SetVolume(volume int) error {
	if volume < 0 || 100 < volume {
		return errors.New("out of valid volume range")
	}
	_, err := execCmd(setVolumeCmd(volume))
	return err
}

// Mute mutes the audio.
func Mute() error {
	_, err := execCmd(muteCmd())
	return err
}

// Unmute unmutes the audio.
func Unmute() error {
	_, err := execCmd(unmuteCmd())
	return err
}
