package internal

import "fmt"

// ConfirmContinue prompts the user to press any key to continue.
func ConfirmContinue() {
	fmt.Println("\nPress any key to continue...")
	GetPressedKey()
}
