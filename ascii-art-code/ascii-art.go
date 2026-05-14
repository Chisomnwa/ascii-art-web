package asciiartcode

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func AsciiArt(text, banner string) (string, error) {
	// ensure the correct banner is used
	valid := banner == "standard" || banner == "thinkertoy" || banner == "shadow"
	if !valid {
		fmt.Println("wrong banner usage")
		return "", errors.New("invalid Banner")
	}

	// read the content of the banner file
	filepath := banner + ".txt"
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("error reading file:", err)
		return "", errors.New("unable to read content")
	}
	data := strings.Split(string(content), "\n")

	inputText := strings.ReplaceAll(text, "\\n", "\n")

	if inputText == "" {
		return "", nil 
	}
	
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
	return res, nil
}
