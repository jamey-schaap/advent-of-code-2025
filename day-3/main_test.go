package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHighestNumber(t *testing.T) {
	actual, err := FindHighestNumber("987654321111111", 2)
	assert.Equal(t, 98, actual)
	assert.NoError(t, err)

	actual, err = FindHighestNumber("811111111111119", 2)
	assert.Equal(t, 89, actual)
	assert.NoError(t, err)

	actual, err = FindHighestNumber("234234234234278", 2)
	assert.Equal(t, 78, actual)
	assert.NoError(t, err)

	actual, err = FindHighestNumber("818181911112111", 2)
	assert.Equal(t, 92, actual)
	assert.NoError(t, err)
}

func TestFindHighestChar(t *testing.T) {
	char, n, err := FindHighestChar("987654321111111", 2)
	assert.Equal(t, 0, n)
	assert.Equal(t, uint8('9'), char)
	assert.NoError(t, err)

	char, n, err = FindHighestChar("11111111111119", 1)
	assert.Equal(t, 0, n)
	assert.Equal(t, uint8('1'), char)
	assert.NoError(t, err)

	char, n, err = FindHighestChar("11111111111119", 0)
	assert.Equal(t, 13, n)
	assert.Equal(t, uint8('9'), char)
	assert.NoError(t, err)
}
