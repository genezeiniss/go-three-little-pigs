package utils

import (
	"strings"
	"unicode"
)

func AmendString(str string) string {

	var amendedString string

	for index, letter := range str {
		if index > 0 && unicode.IsUpper(letter) {
			amendedString = amendedString + " "
		}
		amendedString = amendedString + string(letter)
	}

	return strings.Title(amendedString)
}
