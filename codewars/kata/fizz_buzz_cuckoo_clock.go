package kata

import (
	"strings"
	"strconv"
)

func FizzBuzzCuckooClock(time string) string {
	minute, _ := strconv.Atoi(time[3:])
	switch {
		case minute == 0:
			hour, _ := strconv.Atoi(time[:2])
			switch {
				case hour == 0:
					return strings.Repeat("Cuckoo ", 11) + "Cuckoo"
				case hour > 12:
					hour -= 12
					fallthrough
				default:
					return strings.Repeat("Cuckoo ", hour-1) + "Cuckoo"
			}
		case minute == 30:
			return "Cuckoo"
		case minute % 3 == 0 && minute % 5 == 0:
			return "Fizz Buzz"
		case minute % 3 == 0:
			return "Fizz"
		case minute % 5 == 0:
			return "Buzz"
		default:
			return "tick"
	}
}
