package stringcalculator

import (
  //"fmt"
  //"reflect"
  "strconv"
)

//fmt.Println(numbers)
//fmt.Println(reflect.TypeOf(numbers))

func Add(numbers string) string {

  if(numbers == "") {
    return "0"
  }

  result, _ := strconv.ParseFloat(numbers, 32)

  return strconv.FormatFloat(result, 'f', -1, 32)
}
