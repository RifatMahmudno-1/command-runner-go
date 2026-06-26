package internal

import "fmt"

// ClearConsole clears the console screen.
func ClearConsole() {
	// \033[H: Move to top
	// \033[2J: Clear screen
	// \033[3J: Clear scrollback buffer
	fmt.Print("\033[H\033[2J\033[3J")
}
