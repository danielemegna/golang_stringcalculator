package stringcalculator

import "strconv"

func Add(numbers string) string {
  var result float64 = 0.

  for _, piece := range SplitNumbersString(numbers) {
    result += parseStringPiece(piece)
  }

  return formatResult(result)
}

func formatResult(number float64) string {
  return strconv.FormatFloat(number, 'f', -1, 64)
}

func parseStringPiece(piece string) float64 {
  number, _ := strconv.ParseFloat(piece, 64)
  return number
}
