package main

import (
	"command-runner/internal"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var actions = [...]string{"Run Commands", "Add Command", "Edit Command", "Delete Command", "Exit"}

func main() {
	commandsFilePath, err := internal.GetCommandsPath()
	if err != nil {
		fmt.Println("Failed to determine the path for 'commands.cfg'.")
		internal.ConfirmExit()
		return
	}

	if !internal.FileExists(commandsFilePath) {
		fmt.Println("'commands.cfg' not found. Creating a new one...")
		if err := internal.CreateNewFile(commandsFilePath); err != nil {
			fmt.Println("Failed to create 'commands.cfg'")
			internal.ConfirmExit()
			return
		}
		fmt.Println("'commands.cfg' created successfully.")
		internal.ConfirmContinue()
	}

	file, err := os.OpenFile(commandsFilePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Failed to open 'commands.cfg'")
		internal.ConfirmExit()
		return
	}
	defer file.Close()

	commands, err := internal.ReadCommandsFromFile(file)
	if err != nil {
		fmt.Println("Failed to read commands from 'commands.cfg'")
		internal.ConfirmExit()
		return
	}

	for {
		internal.ClearConsole()
		switch internal.SelectCommand("Select an action:", actions[:]) {
		case -1:
			fmt.Println("No action selected.")
			internal.ConfirmExit()
			return
		case 0: // Run Commands
			selectedIndexes := internal.SelectCommands("Select commands to run:", commands)
			if selectedIndexes == nil {
				fmt.Println("No commands selected to run.")
				internal.ConfirmContinue()
				continue
			}
			fmt.Println("Press Y to execute the selected commands or any other key to cancel.")
			if key := internal.GetPressedKey(); key != "Y" && key != "y" {
				fmt.Println("Execution cancelled.")
				internal.ConfirmContinue()
				return
			}
			internal.ClearConsole()
			for _, i := range selectedIndexes {
				command := commands[i]
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
			internal.ConfirmContinue()
			continue
		case 1: // Add Command
			newCommand, err := internal.TakeInput("Enter the new command (One Line):")
			if err != nil {
				fmt.Println("Failed to add command.")
				internal.ConfirmContinue()
				continue
			}
			commands = append(commands, newCommand)
			internal.WriteCommands(commands, file)
			fmt.Println("Command added successfully.")
			internal.ConfirmContinue()
			continue
		case 2: // Edit Command
			if len(commands) == 0 {
				fmt.Println("No commands are available to edit.")
				internal.ConfirmContinue()
				continue
			}
			selectedIndex := internal.SelectCommand("Select a command to edit:", commands)
			if selectedIndex == -1 {
				fmt.Println("No command selected for editing.")
				internal.ConfirmContinue()
				continue
			}
			newCommand, err := internal.TakeInput("Enter the updated command (One Line):")
			if err != nil {
				fmt.Println("Failed to edit command.")
				internal.ConfirmContinue()
				continue
			}
			commands[selectedIndex] = newCommand
			internal.WriteCommands(commands, file)
			fmt.Println("Command updated successfully.")
			internal.ConfirmContinue()
			continue
		case 3: // Delete Command
			if len(commands) == 0 {
				fmt.Println("No commands are available to delete.")
				internal.ConfirmContinue()
				continue
			}
			selectedIndexes := internal.SelectCommands("Select commands to delete:", commands)
			if selectedIndexes == nil {
				fmt.Println("No commands selected for deletion.")
				internal.ConfirmContinue()
				continue
			}
			for _, i := range selectedIndexes {
				commands = append(commands[:i], commands[i+1:]...)
			}
			internal.WriteCommands(commands, file)
			fmt.Println("Selected commands deleted successfully.")
			internal.ConfirmContinue()
			continue
		case 4: // Exit
			internal.ConfirmExit()
			return
		default:
			fmt.Println("Invalid action selected.")
			internal.ConfirmContinue()
			continue
		}
	}
}
