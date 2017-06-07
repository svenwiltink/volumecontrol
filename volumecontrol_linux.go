package volumecontrol

import (
	"fmt"
	"os/exec"
	"strconv"
)

func SetVolume(volume int) (err error) {
	if volume < 0 || volume > 100 {
		err = fmt.Errorf("invalid volume %n", volume)
		return
	}

	cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", strconv.Itoa(volume)+"%")
	err = cmd.Run()
	return
}

func IncreaseVolume(volume int) (err error) {
	cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", strconv.Itoa(volume)+"%+")
	err = cmd.Run()
	return
}

func DecreaseVolume(volume int) (err error) {
	cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", strconv.Itoa(volume)+"%-")
	err = cmd.Run()
	return
}
