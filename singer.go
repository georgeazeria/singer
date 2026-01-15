package singer

import (
	"fmt"
	"regexp"
	"strings"
	"math/rand"
	"time"
)

const (
	CharacterFilter   = `[\x00-\x1F\/\\:\*\?\"<>\|]`
	UnicodeWhitespace = `[[:space:]]+`
)

var (
	WindowsReservedNames = [...]string{
		"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5",
		"COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5",
		"LPT6", "LPT7", "LPT8", "LPT9",
	}
)

func randomLetters(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	output := make([]byte, length)
	for i := range output {
		output[i] = letters[rand.Intn(len(letters))]
	}
	return string(output)
}


func Sanitise(input string) string {
	return sanitise(input, 0, "")
}

func Sanitize(input string) string { // American Engish compatability layer
	return sanitise(input, 0, "")
}


func SanitiseLength(input string, pad int) string {
	return sanitise(input, pad, "")
}

func SanitizeLength(input string, pad int) string { // American Engish compatability layer
	return sanitise(input, pad, "")
}


func sanitise(input string, pad int) string {
	input = clean(input)
	length := len(input)

	if pad > length {
		return input
	}

	if length > 255 {
		length = 255
	}

	if pad != 0 {
		length -= pad
	}

	return input[0:length]
}

func replace(input string, pattern string, replacement string) string {
	output := regexp.MustCompile(pattern)
	return strings.TrimSpace(output.ReplaceAllString(input, replacement))
}

func clean(input string) string {
	input = replace(input, UnicodeWhitespace, " ")
	input = replace(input, CharacterFilter, "")
	input = replace(input, UnicodeWhitespace, " ")

	return filter(input)
}

func filter(input string) string {
	input = filterWindowsReservedNames(input)
	input = filterBlank(input)
	input = replaceIllegalFinalRune(input)

	return input
}

func filterWindowsReservedNames(input string) string {
	caps := strings.ToUpper(input)

	for i := range WindowsReservedNames {
		reserved := WindowsReservedNames[i]

		if reserved == caps {
			return input + "_" + randomLetters(8)
		}
	}

	return input
}

func filterBlank(input string) string {
	if input == "" {
		return "blank_"+randomLetters(8)
	}
 	return input
}

func replaceIllegalFinalRune(input string) string {
	runes := []rune(input)
	if len(runes) == 0 {
		return input
	}
	if runes[len(runes)-1] == '.' || runes[len(runes)-1] == ',' {
		runes[len(runes)-1] = '_'
	}
	return string(runes)
}