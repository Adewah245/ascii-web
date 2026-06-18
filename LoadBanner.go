package main

import (
	"os"
	"strings"
)

var Banners = make(map[string]map[rune][]string)

func LoadAllBanners() error {
	files := map[string]string{
		"standard":   "banners/standard.txt",
		"shadow":     "banners/shadow.txt",
		"thinkertoy": "banners/thinkertoy.txt",
	}

	for name, path := range files {
		m, err := LoadBanner(path)
		if err != nil {
			return err
		}
		Banners[name] = m
	}
	return nil
}

// === YOUR ORIGINAL FUNCTION (unchanged) ===
func LoadBanner(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, os.ErrInvalid
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) != 855 {
		return nil, os.ErrInvalid
	}
	result := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		start := i*9 + 1
		r := rune(32 + i)
		result[r] = lines[start : start+8]
	}
	return result, nil
}
