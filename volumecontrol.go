package volumecontrol

import (
	"syscall"
	"fmt"
	"time"
)

const (
	VK_VOLUME_MUTE uintptr = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
)

var dll = syscall.NewLazyDLL("user32.dll")
var procKeyBd = dll.NewProc("keybd_event")

func SetVolume(volume int) (err error) {
	if volume < 0 || volume > 100 {
		err = fmt.Errorf("invalid volume %n", volume)
		return
	}

	for i := 0; i < 50; i++ {
		volumeDown()
		time.Sleep(1 * time.Microsecond)
	}
	volume = volume/2

	for i := 0; i < volume; i++ {
		volumeUp()
		time.Sleep(1 * time.Microsecond)
	}

	return
}

func mute() (err error) {
	vkey := VK_VOLUME_MUTE + 0x80
	_, _, err = procKeyBd.Call(uintptr(VK_VOLUME_MUTE), uintptr(vkey), 0, 0)
	return
}

func volumeDown() (err error) {
	vkey := VK_VOLUME_DOWN + 0x80
	_, _, err = procKeyBd.Call(uintptr(VK_VOLUME_DOWN), uintptr(vkey), 0, 0)
	return
}

func volumeUp() (err error) {
	vkey := VK_VOLUME_UP + 0x80
	_, _, err = procKeyBd.Call(uintptr(VK_VOLUME_UP), uintptr(vkey), 0, 0)
	return
}