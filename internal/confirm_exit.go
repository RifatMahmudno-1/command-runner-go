package internal

import "fmt"

// ConfirmExit prompts the user to press any key to exit.
func ConfirmExit() {
	fmt.Println("\nPress any key to exit...")
	GetPressedKey()
}
