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

	err = SetVolume(33)
	t.Logf("%v", err)

	time.Sleep(2 * time.Second)
}
