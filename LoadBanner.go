package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(fileName string) (map[rune][]string, error) {
	// Reading the content of the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: File not found!")
		return nil, err
	}
	// checking if the file is empty
	if len(data) == 0 {
		fmt.Println("Empty banner File")
		return nil, err
	}
	fmt.Println("File Loaded Successfully...")

	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(text, "\n")
	// checking if lines are complete
	if len(lines) != 855 {
		fmt.Println("Incomplete banner file")
		return nil, err
	}

	result := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		start := i*9 + 1

		r := rune(32 + i)
		result[r] = lines[start : start+8]
	}
	return result, nil
}
