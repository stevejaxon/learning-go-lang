package kata

func StringConstructing(a, s string) int {
	return evaluate(a,a,s, 1)
}

func evaluate(origAlphabet, remainingAlphabet, word string, cost int) int {
	if len(word) == 0 { 
		return cost + len(remainingAlphabet)
	}
	if len(remainingAlphabet) == 0 {
		return evaluate(origAlphabet, origAlphabet, word, cost+1)
	}
	if remainingAlphabet[0] == word[0] {
		return evaluate(origAlphabet, remainingAlphabet[1:], word[1:], cost)
	} else {
		return evaluate(origAlphabet, remainingAlphabet[1:], word, cost+1)
	}
}