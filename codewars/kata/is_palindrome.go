package kata

import (
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return false
	}
	str = strings.ToLower(str)
	rightPtr := len(str) - 1
	runes := []rune(str)
	for leftPtr := 0 ; leftPtr < rightPtr; leftPtr, rightPtr = leftPtr +1, rightPtr -1 {
		if runes[leftPtr] != runes[rightPtr] {
			return false
		}
	}
  	return true
}