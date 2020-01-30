package stringcalculator

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSingleIntegerString(t *testing.T) {
  assert.Equal(t, []string{"19"}, SplitInputString("19"))
}

func TestSingleFloatWithDecimalsString(t *testing.T) {
  assert.Equal(t, []string{"19.9"}, SplitInputString("19.9"))
}

func TestSomeNumbersSeparatedByComma(t *testing.T) {
  expected := []string{"19", "46", "3.1415"}
  assert.Equal(t, expected, SplitInputString("19,46,3.1415"))
}

func TestEmptyString(t *testing.T) {
  assert.Equal(t, []string{}, SplitInputString(""))
}

func TestSupportNewLineSparator(t *testing.T) {
  assert.Equal(t, []string{"9", "12", "42.54"}, SplitInputString("9,12\n42.54"))
}
