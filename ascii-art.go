package main

import (
	"fmt"
	"os"
	"strings"
)

func asciiArt(text, banner string) string {
	// banner = "standard"
	valid := banner == "standard" || banner == "thinkertoy" || banner == "shadow"
	if !valid {
		fmt.Println("wrong banner usage")
		return ""
	}

	content, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return ""
	}
	data := strings.Split(string(content), "\n")

	inputText := strings.ReplaceAll(text, "\\n", "\n")
	wordSlice := strings.Split(inputText, "\n")

	var res string

	for _, words := range wordSlice {
		for i := 0; i < 8; i++ {
			for _, char := range words {
				ascidx := int(char-' ')*9 + 1 + i
				charidx := data[ascidx]
				res += charidx
			}
			res += "\n"
		}
	}
	return res
}
