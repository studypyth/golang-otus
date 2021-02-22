package hw02

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sourceString string) (string, error) {
	var builder strings.Builder
	var previousRune rune
	var err error
	for index, currentRune := range sourceString {
		if (unicode.IsDigit(currentRune) && unicode.IsDigit(previousRune)) || (unicode.IsDigit(currentRune) && index == 0) {
			err = ErrInvalidString
		}
		if unicode.IsDigit(currentRune) {
			multiplier, _ := strconv.Atoi(string(currentRune))
			multipleRune := strings.Repeat(string(previousRune), multiplier)
			builder.WriteString(multipleRune)
			previousRune = currentRune
		} else {
			if !unicode.IsDigit(previousRune) {
				builder.WriteRune(previousRune)
			}
			if index+1 == len(sourceString) {
				builder.WriteRune(currentRune)
			}
			previousRune = currentRune
		}
	}
	result := builder.String()
	if len(builder.String()) > 0 {
		result = result[1:]
	}
	return result, err
}
