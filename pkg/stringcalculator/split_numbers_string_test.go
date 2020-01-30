package stringcalculator

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSingleIntegerString(t *testing.T) {
  assert.Equal(t, []string{"19"}, SplitNumbersString("19"))
}

func TestSingleFloatWithDecimalsString(t *testing.T) {
  assert.Equal(t, []string{"19.9"}, SplitNumbersString("19.9"))
}

func TestSomeNumbersSeparatedByComma(t *testing.T) {
  expected := []string{"19", "46", "3.1415"}
  assert.Equal(t, expected, SplitNumbersString("19,46,3.1415"))
}

func TestEmptyString(t *testing.T) {
  assert.Equal(t, []string{}, SplitNumbersString(""))
}
