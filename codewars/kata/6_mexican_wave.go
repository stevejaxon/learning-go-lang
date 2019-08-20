package kata

import "strings"
import "unicode"

func Wave(words string) (output []string) {
	if strings.TrimSpace(words) == "" {
		output = []string{}
		return
	}
	runes := []rune(words)
	for i, word := range runes {
		if word == ' ' { continue }
		orig := word
		runes[i] = unicode.ToUpper(word)
		output = append(output, string(runes))
		runes[i] = orig
	}
	return
}