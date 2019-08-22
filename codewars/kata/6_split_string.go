package kata

import "regexp"

func Solution(str string) (output []string) {
  re, _ := regexp.Compile(`\w{2}`)
  output = re.FindAllString(str, -1)
  if len(str) & 1 == 1 {
    output = append(output, string(str[len(str)-1])+"_")
  }
  return
}
