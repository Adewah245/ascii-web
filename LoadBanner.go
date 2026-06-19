package main

import (
	"fmt"
	"os"
	"strings"
)

var banners = make(map[string]map[rune][]string)

func LoadAllBanners() error {
	file := map[string]string{
		"standard":   "banners/standard.txt",
		"shadow":     "banners/shadow.txt",
		"thinkertoy": "banners/thinkertoy.txt",
	}
	for name, ascii := range file {
		m, err := LoadBanner(ascii)
		if err != nil {
			return err
		}
		banners[name] = m
	}
	return nil
}
func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error File Not Found!")
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println("File Present is Empty")
		os.Exit(1)
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) < 855 || len(lines) > 855 {
		fmt.Println("Error: Incomplete Lines")
		os.Exit(1)
	}
	result := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		start := i*9 + 1
		r := rune(32 + i)
		result[r] = lines[start : start+8]
	}
	return result, nil
}
