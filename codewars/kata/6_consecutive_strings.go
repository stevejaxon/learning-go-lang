package kata

import "strings"

func LongestConsec(strarr []string, k int) string {
	if (len(strarr) == 0 || len(strarr) < k) {
		return ""
	}
	outputStart := 0
	maxLength := 0

	for i := 0; i <= len(strarr) - k; i++ {
		wordsLength := 0
		for j := 0; j < k; j++ {
			wordsLength += len(strarr[i+j])
		}
		if wordsLength > maxLength {
			outputStart = i
			maxLength = wordsLength
		}
	}
	return strings.Join(strarr[outputStart:outputStart+k],"")
}