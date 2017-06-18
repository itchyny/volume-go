package volume

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	origVolume, err := GetVolume()
	if err != nil {
		fmt.Fprintf(os.Stderr, "get volume failed: %+v\n", err)
		os.Exit(1)
	}
	origMuted, err := GetMuted()
	if err != nil {
		fmt.Fprintf(os.Stderr, "get muted failed: %+v\n", err)
		os.Exit(1)
	}
	code := m.Run()
	_ = SetVolume(origVolume)
	if origMuted {
		_ = Mute()
	} else {
		_ = Unmute()
	}
	os.Exit(code)
}

func TestGetVolume(t *testing.T) {
	_, err := GetVolume()
	if err != nil {
		t.Errorf("get volume failed: %+v", err)
	}
}

func TestSetVolume(t *testing.T) {
	for _, vol := range []int{0, 37, 54, 20, 10} {
		err := SetVolume(vol)
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

func TestMute(t *testing.T) {
	err := Mute()
	if err != nil {
		t.Errorf("mute failed: %+v", err)
	}
	muted, err := GetMuted()
	if err != nil {
		t.Errorf("get muted failed: %+v", err)
	}
	if !muted {
		t.Errorf("mute failed: %t", muted)
	}
}

func TestUnmute(t *testing.T) {
	err := Unmute()
	if err != nil {
		t.Errorf("unmute failed: %+v", err)
	}
	muted, err := GetMuted()
	if err != nil {
		t.Errorf("get muted failed: %+v", err)
	}
	if muted {
		t.Errorf("unmute failed: %t", muted)
	}
}
