package internal

import (
	"os"
	"path/filepath"
)

// GetCommandsPath returns the path to the commands configuration file.
func GetCommandsPath() (string, error) {
	if IsDev {
		return "commands.cfg", nil
	}
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exe)
	return filepath.Join(exeDir, "commands.cfg"), nil
}
