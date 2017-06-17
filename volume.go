package volume

import (
	"errors"
	"os"
	"os/exec"
)

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
