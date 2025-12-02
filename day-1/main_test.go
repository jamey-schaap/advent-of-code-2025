package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDial(t *testing.T) {
	// Test L-shift, didn't pass zero
	pos, passedZeroCount, err := Dial(78, 'L', 5)
	assert.NoError(t, err)
	assert.Equal(t, 73, pos)
	assert.Equal(t, 0, passedZeroCount)

	// Test L-shift, passed zero once
	pos, passedZeroCount, err = Dial(0, 'L', 1)
	assert.NoError(t, err)
	assert.Equal(t, 99, pos)
	assert.Equal(t, 0, passedZeroCount)

	// Test L-shift, passed zero multiple times
	pos, passedZeroCount, err = Dial(67, 'L', 265)
	assert.NoError(t, err)
	assert.Equal(t, 2, pos)
	assert.Equal(t, 2, passedZeroCount)

	// Test R-shift, didn't pass zero
	pos, passedZeroCount, err = Dial(23, 'R', 5)
	assert.NoError(t, err)
	assert.Equal(t, 28, pos)
	assert.Equal(t, 0, passedZeroCount)

	// Test R-shift, passed zero once
	pos, passedZeroCount, err = Dial(99, 'R', 1)
	assert.NoError(t, err)
	assert.Equal(t, 0, pos)
	assert.Equal(t, 1, passedZeroCount)

	// Test R-shift, passed zero multiple times
	pos, passedZeroCount, err = Dial(52, 'R', 201)
	assert.NoError(t, err)
	assert.Equal(t, 53, pos)
	assert.Equal(t, 2, passedZeroCount)

	// Test the full example from part 2
	passedZeroCount = 0
	pos = 50

	pos, n, err := Dial(pos, 'L', 68)
	assert.Equal(t, 82, pos)
	assert.Equal(t, 1, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'L', 30)
	assert.Equal(t, 52, pos)
	assert.Equal(t, 0, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'R', 48)
	assert.Equal(t, 0, pos)
	assert.Equal(t, 1, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'L', 5)
	assert.Equal(t, 95, pos)
	assert.Equal(t, 0, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'R', 60)
	assert.Equal(t, 55, pos)
	assert.Equal(t, 1, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'L', 55)
	assert.Equal(t, 0, pos)
	assert.Equal(t, 1, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'L', 1)
	assert.Equal(t, 99, pos)
	assert.Equal(t, 0, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'L', 99)
	assert.Equal(t, 0, pos)
	assert.Equal(t, 1, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'R', 14)
	assert.Equal(t, 14, pos)
	assert.Equal(t, 0, n)
	passedZeroCount += n

	pos, n, err = Dial(pos, 'L', 82)
	assert.Equal(t, 32, pos)
	assert.Equal(t, 1, n)
	passedZeroCount += n

	assert.Equal(t, 6, passedZeroCount)
}
