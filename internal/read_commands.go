package internal

import (
	"bufio"
	"os"
	"strings"
)

// ReadCommandsFromFile reads commands from a file and returns them as a slice of strings.
// It ignores empty lines and lines starting with '#'.
// If an error occurs while reading the file, it returns nil.
func ReadCommandsFromFile(file *os.File) ([]string, error) {
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
