package volumecontrol

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

func ChangeVolume(volume int) (err error) {
	if volume < 0 || volume > 100 {
		err = errors.New(fmt.Sprintf("invalid volume %n", volume))
		return
	}

	volumeFloat := float64(volume) / 10

	cmd := exec.Command("osascript", "-e", "set Volume " + strconv.FormatFloat(volumeFloat, 'f', 1, 64))
	err = cmd.Run()
	return
}