package internal

import "fmt"

// SelectCommands displays a list of commands and allows the user to select multiple commands using the keyboard.
// It returns a slice of the selected commands indexes, or nil slice if the user cancels the selection.
func SelectCommands(label string, commands []string) []int {
	selected := make([]bool, len(commands))
	cursor := 0

	for {
		ClearConsole()
		if label != "" {
			fmt.Println(label)
		}
		fmt.Println("Use UP/DOWN, SPACE to toggle, ENTER to confirm, ESC to cancel:")

		for idx, cmd := range commands {
			checkMark, cursorMark := "[ ]", " "
			if selected[idx] {
				checkMark = "[x]"
			}
			if idx == cursor {
				cursorMark = ">"
			}
			fmt.Printf("%s %s %s\n", cursorMark, checkMark, cmd)
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
		case "space":
			selected[cursor] = !selected[cursor]
		case "enter":
			var selectedIndexes []int
			for i, s := range selected {
				if s {
					selectedIndexes = append(selectedIndexes, i)
				}
			}
			if len(selectedIndexes) == 0 {
				return nil
			}
			return selectedIndexes
		case "esc":
			return nil
		default:
			continue
		}
	}
}
