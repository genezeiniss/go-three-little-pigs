package utils

import (
	"strings"
	"unicode"
)

func AmendString(str string) string {

	var amendedString string

	for _, letter := range str {
		if unicode.IsUpper(letter){
			amendedString = amendedString + " "
		}
		amendedString = amendedString + string(letter)
	}

	return strings.Title(amendedString)
}
