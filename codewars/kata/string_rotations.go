package kata

// Solution to https://www.codewars.com/kata/all-inclusive/train/go
func ContainAllRots(input string, arr []string) bool { 
  if input == "" {
		return true
	}
	rotations := rotations(input)
	for _, perm := range rotations {
		if ! contains(perm, arr) {
			return false
		}
	}
	return true
}

func rotations(input string) []string {
	var rotations []string
	rotations = append(rotations, input)
	for i := 1; i < len(input); i++ {
		rotations = append(rotations, input[i:] + input[:i])
	}
	return rotations
}

func contains(needle string, haystack []string) bool {
	for _, word := range haystack {
		if word == needle {
			return true
		}
	}
	return false
}