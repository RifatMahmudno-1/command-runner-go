package internal

import (
	"os"
)

// CreateNewFile creates a new file with the specified filename.
func CreateNewFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
