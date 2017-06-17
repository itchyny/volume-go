package volume

import (
	"testing"
)

func TestGetVolume(t *testing.T) {
	_, err := GetVolume()
	if err != nil {
		t.Errorf("get volume failed: %+v", err)
	}
}

func TestSetVolume(t *testing.T) {
	vol, err := GetVolume()
	defer SetVolume(vol)
	if err != nil {
		t.Errorf("get volume failed: %+v", err)
	}
	for _, vol := range []int{0, 37, 54, 20, 10} {
		err = SetVolume(vol)
		if err != nil {
			t.Errorf("set volume failed: %+v", err)
		}
		v, err := GetVolume()
		if err != nil {
			t.Errorf("set volume failed: %+v", err)
		}
		if vol != v {
			t.Errorf("set volume failed: (got: %+v, expected: %+v)", v, vol)
		}
	}
}
