package kata

// Solution for https://www.codewars.com/kata/a-rule-of-divisibility-by-7/train/go - see for more details

type Divisible struct {
	lastNum int64
	steps int
}

func Seven(n int64) (result []int) {
	if n < 91 {
		return []int{0,0}
	}
	data := Divisible{n, 0}
	seven(&data)
	return []int{int(data.lastNum), data.steps}
}

func seven(result *Divisible) {
	lastDigit := result.lastNum % 10
	carryForward := result.lastNum / 10
	carryForward -= int64(lastDigit * 2)

	// The carry forward string is of length 2 or less, so we terminate
	if carryForward < 100 {
		if carryForward % 7 != 0 {
			// Set error output format
			result.lastNum = 0
			result.steps = 0
		} else {
			result.lastNum = carryForward
			result.steps = result.steps + 1
		}
		return 
	}

	result.lastNum = carryForward
	result.steps = result.steps + 1

	seven(result)

	return
}