package main

import "strings"

func inputFile(input string) []string {
	if input == "" {
		return []string{}
	}
	text := strings.ReplaceAll(input, "\\n", "\n")
	if text == "\n" {
		return []string{""}
	}
	return strings.Split(text, "\n")
}
