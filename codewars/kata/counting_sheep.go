package kata

import (
	"strings"
	"strconv"
)

const suffix = " sheep..."

func CountSheep(num int) string {
  if num <= 0 {
	  return ""
  }

  builder := strings.Builder{}
  for i := 1; i <= num; i++ {
	builder.WriteString(strconv.Itoa(i))
	builder.WriteString(suffix)
  }
  return builder.String()
}