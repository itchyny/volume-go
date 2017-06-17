package volume

import (
	"errors"
	"os/exec"
)

func GetVolume() (int, error) {
	cmds := getVolumeCmd()
	out, err := exec.Command(cmds[0], cmds[1:]...).Output()
	if err != nil {
		return 0, err
	}
	return parseVolume(string(out))
}

func SetVolume(volume int) error {
	if volume < 0 || 100 < volume {
		return errors.New("out of valid volume range")
	}
	cmds := setVolumeCmd(volume)
	_, err := exec.Command(cmds[0], cmds[1:]...).Output()
	return err
}
