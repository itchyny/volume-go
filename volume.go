package volume

import (
	"errors"
	"os"
	"os/exec"
)

// GetVolume returns the current volume (0 to 100).
func GetVolume() (int, error) {
	cmdArgs := getVolumeCmd()
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Env = append(os.Environ(), cmdEnv()...)
	out, err := cmd.Output()
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
	cmdArgs := setVolumeCmd(volume)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Env = append(os.Environ(), cmdEnv()...)
	_, err := cmd.Output()
	return err
}

// Mute mutes the audio.
func Mute() error {
	cmdArgs := muteCmd()
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Env = append(os.Environ(), cmdEnv()...)
	_, err := cmd.Output()
	return err
}

// Unmute unmutes the audio.
func Unmute() error {
	cmdArgs := unmuteCmd()
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Env = append(os.Environ(), cmdEnv()...)
	_, err := cmd.Output()
	return err
}
