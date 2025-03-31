package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// checks whether the text consists only of dots and dashes.
func checkMorse(text string) bool {
	arr := []rune(text)
	for _, v := range arr {
		if v != '.' && v != '-' && v != ' ' {
			return false
		}
	}
	return true
}

// checks for the use of valid characters for Morse code conversion.
func checkToMorse(text string) bool {
	text = strings.ToUpper(text)
	text = strings.ReplaceAll(text, " ", "")
	arr := []rune(text)
	for _, v := range arr {
		_, ok := morse.DefaultMorse[v]
		if !ok {
			return false
		}
	}
	return true
}

// CheckMorseAndConvert converts Morse to text if the input consists of dots and dashes. Otherwise, it converts the input string to Morse code.
func СheckMorseAndConvert(text string) (string, error) {
	if checkMorse(text) {
		return morse.ToText(text), nil
	} else if checkToMorse(text) {
		return morse.ToMorse(text), nil
	}
	return "", fmt.Errorf("the file cannot be converted. Invalid characters")
}
