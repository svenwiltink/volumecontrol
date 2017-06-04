package volumecontrol

import (
	"testing"
	"time"
)

func TestSetVolume(t *testing.T) {
	err := SetVolume(20)
	t.Logf("%v", err)

	time.Sleep(2 * time.Second)

	err = SetVolume(5)
	t.Logf("%v", err)

	time.Sleep(2 * time.Second)

	err = SetVolume(22)
	t.Logf("%v", err)

	time.Sleep(2 * time.Second)

	err = SetVolume(12)
	t.Logf("%v", err)

	time.Sleep(2 * time.Second)
}

func TestIncreaseVolume(t *testing.T) {
	err := IncreaseVolume(10)
	t.Logf("%v", err)
}

func TestDecreaseVolume(t *testing.T) {
	err := DecreaseVolume(10)
	t.Logf("%v", err)
}
