package stringcalculator

import "strings"

func SplitInputString(numbers string) []string {
  if(numbers == "") {
    return []string{}
  }

  result := strings.FieldsFunc(numbers, func(c rune) bool {
    return c == ',' || c == '\n'
  })

  return result
}
