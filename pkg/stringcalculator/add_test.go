package stringcalculator

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAreRunning(t *testing.T) {
  assert.True(t, true, "True is true!")
}

func TestAddAnEmptyString_ShouldReturnZero(t *testing.T) {
  assert.Equal(t, "0", Add(""))
}

func TestAddASingleNumberString(t *testing.T) {
  assert.Equal(t, "1", Add("1"))
  assert.Equal(t, "42", Add("42"))
  assert.Equal(t, "19.9", Add("19.9"))
}

func TestSumCoupleOfNumbers(t *testing.T) {
  assert.Equal(t, "3", Add("1,2"))
  assert.Equal(t, "16.9", Add("16.2,0.7"))
}

func TestSupportNewLineSeparator(t *testing.T) {
  assert.Equal(t, "5", Add("3\n2"))
  assert.Equal(t, "6", Add("1\n2,3"))
}
