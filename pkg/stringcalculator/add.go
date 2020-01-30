package stringcalculator

import "strconv"


func Add(numbers string) string {
  if(numbers == "") {
    return "0"
  }

  result := 0.
  for _, number_string := range SplitNumbersString(numbers) {
    number_float, _ := strconv.ParseFloat(number_string, 64)
    result += number_float
  }

  return strconv.FormatFloat(result, 'f', -1, 64)
}
