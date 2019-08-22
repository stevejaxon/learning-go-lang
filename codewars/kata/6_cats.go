package kata

import "math"

func PeacefulYard(yard []string, minDistance int) bool {
	if minDistance <= 0 { return true }
	l := [2]int{-1, -1}
	m := [2]int{-1, -1}
	r := [2]int{-1, -1}
	for row, rowData := range yard {
		for col, colData := range rowData {
			if colData != '-' {
				switch colData {
					case 'L':
						l = [2]int{row, col}
						if ! (comfortableDistance(l, m, minDistance) && comfortableDistance(l, r, minDistance)) {
							return false
						}
					case 'R':
						r = [2]int{row, col}
						if ! (comfortableDistance(r, m, minDistance) && comfortableDistance(r, l, minDistance)) {
							return false
						}
					case 'M':
						m = [2]int{row, col}
						if ! (comfortableDistance(m, r, minDistance) && comfortableDistance(m, l, minDistance)) {
							return false
						}
				}
			}
		}
	}
	return true	
}

func measureDistance(cat1, cat2 [2]int) int {
	return int(math.Sqrt(math.Pow(float64(cat1[0]) - float64(cat2[0]), 2) + math.Pow(float64(cat1[1]) - float64(cat2[1]), 2)))
}

func comfortableDistance(cat1, cat2 [2]int, minDistance int) bool {
	return cat2 == [2]int{-1, -1} || measureDistance(cat1, cat2) >= minDistance
}
