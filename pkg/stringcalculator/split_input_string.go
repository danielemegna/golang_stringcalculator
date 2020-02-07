package stringcalculator

import "strings"

func SplitInputString(numbers string) ([]string, *string) {
  if(numbers == "") {
    return []string{}, nil
  }

  result := strings.FieldsFunc(numbers, func(c rune) bool {
    return c == ',' || c == '\n'
  })

  return result, nil
}
