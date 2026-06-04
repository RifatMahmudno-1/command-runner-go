package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func fileExists(filename *string) bool {
	_, err := os.Stat(*filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func createNewFile(filename *string) bool {
	// create the file
	file, err := os.Create(*filename)
	if err != nil {
		return false
	}

	// ensure the file is closed after writing
	defer file.Close()

	// prepare the content to write to the file
	var stringBuilder strings.Builder
	stringBuilder.WriteString("# Add your commands here, one per line.\n")
	stringBuilder.WriteString("# Lines starting with # are comments and will be ignored.\n")
	stringBuilder.WriteString("# Example:\n")
	stringBuilder.WriteString("# echo Hello, World!\n")

	// write the content to the file
	_, err = file.WriteString(stringBuilder.String())
	if err != nil {
		return false
	}
	return true
}

func readCommandsFromFile(filename *string) []string {
	// open the file for reading
	file, err := os.Open(*filename)
	if err != nil {
		return nil
	}
	defer file.Close()

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
	return commands
}

func getPressedKey() string {
	pressedKey := ""
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		pressedKey = key.String()
		return true, nil
	})
	return pressedKey
}

func getCommandsPath() string {
	if IsDev {
		return "commands.cfg"
	}
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	exeDir := filepath.Dir(exe)
	return filepath.Join(exeDir, "commands.cfg")
}

func main() {
	commandsFilePath := getCommandsPath()

	if commandsFilePath == "" {
		fmt.Println("Failed to determine the path for 'commands.cfg'.")
		return
	}

	if !fileExists(&commandsFilePath) {
		fmt.Println("'commands.cfg' not found. Creating a new one with instructions...")
		if !createNewFile(&commandsFilePath) {
			fmt.Println("Failed to create 'commands.cfg'")
			return
		}
		fmt.Println("'commands.cfg' created successfully. Add your commands to the file and run this script again.")
		return
	}

	commands := readCommandsFromFile(&commandsFilePath)
	if commands == nil {
		fmt.Println("Failed to read commands from 'commands.cfg'")
		return
	}
	if len(commands) == 0 {
		fmt.Println("No commands found in 'commands.cfg'. Please add some commands and run this script again.")
		return
	}

	fmt.Println("The following commands will be executed:")
	for idx, cmd := range commands {
		num := fmt.Sprintf("%3d", idx+1)
		fmt.Printf("%s) %s\n", num, cmd)
	}

	fmt.Println("Press Y to execute the above commands, or any other key to cancel: ")
	if key := getPressedKey(); key != "Y" && key != "y" {
		fmt.Println("Execution cancelled.")
		return
	}

	for _, command := range commands {
		fmt.Printf("\nExecuting: %s\n", command)
		// Prepare the command to be executed in the Windows command prompt
		cmd := exec.Command("cmd", "/C", command)
		var stdout, stderr strings.Builder
		// Capture the standard output
		cmd.Stdout = &stdout
		// Capture the standard error
		cmd.Stderr = &stderr
		// Run the command
		err := cmd.Run()

		errStr := strings.TrimSpace(stderr.String())
		outStr := strings.TrimSpace(stdout.String())

		if err != nil {
			fmt.Printf("Error executing command: %s\n", strings.TrimSpace(err.Error()))
			if len(errStr) > 0 {
				fmt.Printf("Error Output:\n%s\n", errStr)
			}
		} else if len(errStr) > 0 {
			fmt.Printf("Error Output:\n%s\n", errStr)
		} else if len(outStr) > 0 {
			fmt.Printf("%s\n", outStr)
		}
	}
}
