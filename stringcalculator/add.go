package stringcalculator

import (
  //"fmt"
  //"reflect"
  "strconv"
  "strings"
)

//fmt.Println(numbers)
//fmt.Println(reflect.TypeOf(numbers))

func Add(numbers string) string {

  if(numbers == "") {
    return "0"
  }

  pieces := strings.Split(numbers, ",")

  result := 0.
  for _, piece := range pieces {
    number, _ := strconv.ParseFloat(piece, 64)
    result += number
  }

  return strconv.FormatFloat(result, 'f', -1, 64)
}
