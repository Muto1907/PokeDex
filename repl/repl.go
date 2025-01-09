package repl

import "strings"

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
	return []string{}
}
