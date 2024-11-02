package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStrRangeToInt_SingleValue(t *testing.T) {
	start, end, err := convertStrRangeToInt("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, start)
	assert.Equal(t, 1, end)
}

func TestConvertStrRangeToInt_StartRange(t *testing.T) {
	start, end, err := convertStrRangeToInt("1-")
	assert.Nil(t, err)
	assert.Equal(t, 1, start)
	assert.Equal(t, 0, end)
}

func TestConvertStrRangeToInt_StartEndRange(t *testing.T) {
	start, end, err := convertStrRangeToInt("1-2")
	assert.Nil(t, err)
	assert.Equal(t, 1, start)
	assert.Equal(t, 2, end)
}
