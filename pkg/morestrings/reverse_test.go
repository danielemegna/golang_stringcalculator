package morestrings

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestGoIsRunnung(t *testing.T) {
  assert.True(t, true, "True is true!")
}

func TestDoNotTouchPalindromes(t *testing.T) {
	assert.Equal(t, "repaper", ReverseRunes("repaper"))
	assert.Equal(t, "radar", ReverseRunes("radar"))
}

func TestDoubleConversion(t *testing.T) {
	assert.Equal(t, "Double conversion", ReverseRunes(ReverseRunes("Double conversion")))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, "eleinaD", ReverseRunes("Daniele"))
  assert.Equal(t, "Hello, Go!", ReverseRunes("!oG ,olleH"))
	assert.Equal(t, "wow radar level", ReverseRunes("level radar wow"))
}
