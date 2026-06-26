package internal

import "fmt"

// SelectCommand displays a list of commands and allows the user to select one using the keyboard.
// It returns the index of the selected command, or -1 if the user cancels the selection.
func SelectCommand(label string, commands []string) int {
	var cursor = 0

	for {
		ClearConsole()
		if label != "" {
			fmt.Println(label)
		}
		fmt.Println("Use UP/DOWN, SPACE to toggle, ENTER to confirm, ESC to cancel:")

		for idx, cmd := range commands {
			cursorMark := " "
			if idx == cursor {
				cursorMark = ">"
			}
			fmt.Printf("%s %s\n", cursorMark, cmd)
		}

		key := GetPressedKey()
		switch key {
		case "up":
			if cursor > 0 {
				cursor--
			}
		case "down":
			if cursor < len(commands)-1 {
				cursor++
			}
		case "enter":
			return cursor
		case "esc":
			return -1
		default:
			continue
		}
	}
}
