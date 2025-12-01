package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDial(t *testing.T) {
	result, err := Dial(0, 'L', 1)
	assert.NoError(t, err)
	assert.Equal(t, 99, result)

	result, err = Dial(67, 'L', 265)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	result, err = Dial(99, 'R', 1)
	assert.NoError(t, err)
	assert.Equal(t, 0, result)

	result, err = Dial(52, 'R', 50)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)
}
