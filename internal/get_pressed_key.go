package internal

import (
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

// GetPressedKey waits for a key press and returns the string representation of the pressed key.
func GetPressedKey() string {
	pressedKey := ""
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		pressedKey = key.String()
		return true, nil
	})
	return pressedKey
}
