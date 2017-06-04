package volumecontrol

import (
	"fmt"
	"syscall"
	"time"
)

const (
	VK_VOLUME_MUTE uintptr = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
)

const (
	WINDOWS_VOLUME_STEPSIZE = 2
)

var dll = syscall.NewLazyDLL("user32.dll")
var procKeyBd = dll.NewProc("keybd_event")

var volumeState int = -1

func SetVolume(volume int) (err error) {
	if volume < 0 || volume > 100 {
		err = fmt.Errorf("invalid volume %n", volume)
		return
	}

	if volumeState < 0 || !WindowsKeepVolumeState {
		// Reset volume to zero, so we know what the current state is.
		maxSteps := 100 / WINDOWS_VOLUME_STEPSIZE
		for i := 0; i < maxSteps; i++ {
			volumeDown()
			time.Sleep(1 * time.Millisecond)
		}
		volumeState = 0
	}

	volume = volume / 2 * 2 // Round to even numbers
	if volumeState < volume {
		for i := volumeState; i < volume; i += WINDOWS_VOLUME_STEPSIZE {
			volumeUp()
			time.Sleep(1 * time.Millisecond)
		}
	} else if volumeState > volume {
		for i := volumeState; i > volume; i -= WINDOWS_VOLUME_STEPSIZE {
			volumeDown()
			time.Sleep(1 * time.Millisecond)
		}
	}
	volumeState = volume
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

func IncreaseVolume(volume int) (err error) {
	for i=0; i < volume; i+=WINDOWS_VOLUME_STEPSIZE{
		volumeUp()
	}
}

func DecreaseVolume(volume int) (err error) {
	for i=0; i < volume; i+=WINDOWS_VOLUME_STEPSIZE{
		volumeDown()
	}
}
