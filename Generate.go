package main

import "strings"

func Generate(lines []string, banner map[rune][]string) string {
	var b strings.Builder
	for _, line := range lines {
		if line == "" {
			b.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range line {
				if ascii, ok := banner[ch]; ok {
					b.WriteString(ascii[row])
				}
			}
			b.WriteString("\n")
		}
	}
	return b.String()
}
