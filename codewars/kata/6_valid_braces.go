package kata

func ValidBraces(str string) bool {
	stack := []rune{rune(str[0])}
  	for _, brace := range str[1:] {
    	switch brace {
			case '(', '[', '{':
        		stack = append(stack, brace)
      		case ')', ']', '}':
				if len(stack) == 0 {
					return false
				}
				n := len(stack)-1
				head := stack[n]
				diff := int(brace) - int(head)
				// the char codes for () are contiguous, but for {} and [] they are separated by one char 
				if diff < 0 || diff > 2 {
					return false
				}
				// pop
				stack = stack[:n]
    	}
	}
	return len(stack) == 0
}