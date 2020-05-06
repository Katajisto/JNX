package main

import (
	"strings"
)

func stringToSlice(x string, sep rune) []string {
	var msgSlice []string
	curText := ""
	for _, char := range x {
		if char == sep {
			msgSlice = append(msgSlice, strings.ToLower(curText))
			curText = ""
			continue
		}
		curText += string(char)
	}
	msgSlice = append(msgSlice, strings.ToLower(curText))
	return msgSlice
}
