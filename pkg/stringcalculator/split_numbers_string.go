package stringcalculator

import "strings"

func SplitNumbersString(numbers string) []string {
  if(numbers == "") {
    return []string{}
  }

  result := strings.Split(numbers, ",")
  return result
}
