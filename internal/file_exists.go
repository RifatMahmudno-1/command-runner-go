package internal

import "os"

// FileExists checks if a file exists at the specified path.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
