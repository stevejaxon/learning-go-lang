package kata

import "math"

func Century(year int) int {
  return int(math.Ceil(float64(year)/100))
}