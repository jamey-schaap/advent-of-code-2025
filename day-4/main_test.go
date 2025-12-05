package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnswer(t *testing.T) {
	input := "..@@.@@@@.\n" +
		"@@@.@.@.@@\n" +
		"@@@@@.@.@@\n" +
		"@.@@@@..@.\n" +
		"@@.@@@@.@@\n" +
		".@@@@@@@.@\n" +
		".@.@.@.@@@\n" +
		"@.@@@.@@@@\n" +
		".@@@@@@@@.\n" +
		"@.@.@@@.@."

	r := bytes.NewReader([]byte(input))
	answer, err := GetAnswer(r)
	assert.Equal(t, 13, answer)
	assert.NoError(t, err)
}
