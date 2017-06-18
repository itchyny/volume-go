package volume

// GetVolume returns the current volume (0 to 100).
func GetVolume() (int, error) {
	panic("not implemented on Windows")
	return 0, nil
}

// SetVolume sets the sound volume to the specified value.
func SetVolume(volume int) error {
	panic("not implemented on Windows")
	return nil
}

// IncreaseVolume increases (or decreases) the audio volume by the specified value.
func IncreaseVolume(diff int) error {
	panic("not implemented on Windows")
	return nil
}

// GetMuted returns the current muted status.
func GetMuted() (bool, error) {
	panic("not implemented on Windows")
	return false, nil
}

// Mute mutes the audio.
func Mute() error {
	panic("not implemented on Windows")
	return nil
}

// Unmute unmutes the audio.
func Unmute() error {
	panic("not implemented on Windows")
	return nil
}
