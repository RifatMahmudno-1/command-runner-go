package internal

import (
	"os"
	"strings"
)

func WriteCommands(commands []string, file *os.File) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.WriteString(strings.Join(commands, "\n"))
	if err != nil {
		return err
	}

	return nil
}
