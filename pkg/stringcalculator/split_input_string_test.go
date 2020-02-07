package stringcalculator

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSingleIntegerString(t *testing.T) {
  assertSplitInputString(t, "19", []string{"19"})
}

func TestSingleFloatWithDecimalsString(t *testing.T) {
  assertSplitInputString(t, "19.9", []string{"19.9"})
}

func TestSomeNumbersSeparatedByComma(t *testing.T) {
  assertSplitInputString(t, "19,46,3.1415", []string{"19", "46", "3.1415"})
}

func TestEmptyString(t *testing.T) {
  assertSplitInputString(t, "", []string{})
}

func TestSupportNewLineSparator(t *testing.T) {
  assertSplitInputString(t, "9,12\n42.54", []string{"9", "12", "42.54"})
}

func assertSplitInputString(t *testing.T, input string, expectedResult []string) {
  assertSplitInputStringWithErrors(t, input, expectedResult, nil)
}

func assertSplitInputStringWithErrors(t *testing.T, input string, expectedResult []string, expectedError *string) {
  actualResult, actualError := SplitInputString(input)
  assert.Equal(t, expectedResult, actualResult)
  assert.Equal(t, expectedError, actualError)
}
