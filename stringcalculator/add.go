package stringcalculator

import (
  "fmt"
  "reflect"
)


func Add(numbers string) string {
  fmt.Println(numbers)
  fmt.Println(reflect.TypeOf(numbers))

  if(numbers == "") {
    return "0"
  }

  return numbers
}
