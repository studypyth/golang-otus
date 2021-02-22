package hw02_unpack_string

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sourceString string) (string, error) {
	var result strings.Builder
	var previousRune rune
	for index, currentRune := range sourceString {
		if unicode.IsDigit(currentRune) {
			if unicode.IsDigit(previousRune) || index == 0 {
				return "", ErrInvalidString
			} else {
				multiplier, _ := strconv.Atoi(string(currentRune))
				multipleRune := strings.Repeat(string(previousRune), multiplier)
				result.WriteString(multipleRune)
				previousRune = currentRune
			}
		} else {
			if !unicode.IsDigit(previousRune) {
				result.WriteRune(previousRune)
			}
			if index+1 == len(sourceString) {
				result.WriteRune(currentRune)
			}
			previousRune = currentRune
		}
	}
	if len(result.String()) > 0 {
		return result.String()[1:], nil
	} else {
		return result.String(), nil
	}
}
