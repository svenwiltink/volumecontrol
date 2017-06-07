package volumecontrol

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

func SetVolume(volume int) (err error) {
	if volume < 0 || volume > 100 {
		err = errors.New(fmt.Sprintf("invalid volume %n", volume))
		return
	}

	volumeFloat := float64(volume) / 100 * 7

	cmd := exec.Command("osascript", "-e", "set Volume "+strconv.FormatFloat(volumeFloat, 'f', 1, 64))
	err = cmd.Run()
	return
}

func IncreaseVolume(volume int) (err error) {
	err = errors.New("not implemented")
	return
}

func DecreaseVolume(volume int) (err error) {
	err = errors.New("not implemented")
	return
}
