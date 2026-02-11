package control

import "strings"

// Each command is an array of words
type Command []string

// ParseCommand given a string line returns its command
func ParseCommand(line string) Command {

	// Trim and split by spaces
	return strings.Fields(strings.TrimSpace(line))
}
