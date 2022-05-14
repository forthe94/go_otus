package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func CheckNumber(num uint8) bool {
	if num >= 48 && num < 58 {
		return true
	}
	return false
}

func Unpack(input string) (string, error) {
	prevLetter := -1
	ret := ""
	for i := 0; i < len(input); i++ {
		if CheckNumber(input[i]) {
			if prevLetter < 0 {
				return "", ErrInvalidString
			}
			if input[i] == 48 {
				ret = ret[:len(ret)-1]
			} else {
				for j := 0; j < int(input[i]-49); j++ {
					ret += string(uint8(prevLetter))
				}
			}
			prevLetter = -1
		} else {
			prevLetter = int(input[i])
			ret += string(input[i])
		}
	}
	// Place your code here.
	return ret, nil
}
