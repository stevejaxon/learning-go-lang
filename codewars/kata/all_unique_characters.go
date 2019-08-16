package kata

func HasUniqueChar (str string) bool {
  seen := make(map[rune]struct{})
  for _, char := range str {
	if _, prevSeen := seen[char]; prevSeen {
		return false
	}
	seen[char] = struct{}{}
  }
  return true
}