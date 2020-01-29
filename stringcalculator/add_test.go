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

func TestAddAnSingleNumberString(t *testing.T) {
  assert.Equal(t, "1", Add("1"))
  assert.Equal(t, "42", Add("42"))
  assert.Equal(t, "19.9", Add("19.9"))
}
