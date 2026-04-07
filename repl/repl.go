package repl

import "strings"

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}
