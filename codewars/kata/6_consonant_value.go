package kata

func Solve(str string) int {
	return maxConsonantValue(consonantSections(str))
}

func consonantSections(str string) (consonants []string) {
	vowels := map[rune]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
	from := 0
	for i, r := range str {
		if vowels[r] {
			if i > from {
				consonants = append(consonants, str[from:i])
			}
			from = i + 1
		}
	}
	consonants = append(consonants, str[from:len(str)])
	return
}

func consonantsValue(consonant string) (value int) {
	for _, r := range consonant {
		value += int(r) - 96
	}
	return
}

func maxConsonantValue(consonants []string) (max int) {
	for _, consonant := range consonants {
		value := consonantsValue(consonant)
		if value > max {
			max = value
		}
	}
	return
}