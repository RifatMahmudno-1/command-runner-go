package internal

import (
	"bufio"
	"os"
	"strings"
)

// ReadCommands reads commands from a file and returns them as a slice of strings.
// It ignores empty lines and lines starting with '#'.
func ReadCommands(file *os.File) ([]string, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	commands := []string{}
	scanner := bufio.NewScanner(file)

	// read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}
		commands = append(commands, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}
