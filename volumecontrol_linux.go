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

	cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", strconv.Itoa(volume)+"%")
	err = cmd.Run()
	return
}
