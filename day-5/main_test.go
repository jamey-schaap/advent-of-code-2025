package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnswer(t *testing.T) {
	input := "3-5\n" +
		"10-14\n" +
		"16-20\n" +
		"12-18\n" +
		"\n" +
		"1\n" +
		"5\n" +
		"8\n" +
		"11\n" +
		"17\n" +
		"32"

	r := bytes.NewReader([]byte(input))
	answer, err := GetAnswer(r)
	assert.NoError(t, err)
	assert.Equal(t, 3, answer)
}
