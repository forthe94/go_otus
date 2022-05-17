package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	prevLetter := -1
	ret := ""
	for _, char := range input {
		if unicode.IsDigit(char) {
			if prevLetter < 0 {
				return "", ErrInvalidString
			}
			if char == '0' {
				ret = ret[:len(ret)-1]
			} else {
				for j := 0; j < int(char-49); j++ {
					ret += string(rune(prevLetter))
				}
			}
			prevLetter = -1
		} else {
			prevLetter = int(char)
			ret += string(char)
		}
	}
	// Place your code here.
	return ret, nil
}
