package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnswerPart1(t *testing.T) {
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
	answer, err := GetAnswerPart1(r)
	assert.NoError(t, err)
	assert.Equal(t, 3, answer)
}

func TestGetAnswerPart2(t *testing.T) {
	input := "3-5\n" +
		"10-14\n" +
		"16-20\n" +
		"12-18"

	r := bytes.NewReader([]byte(input))
	answer, err := GetAnswerPart2(r)
	assert.NoError(t, err)
	assert.Equal(t, 14, answer)
}

func TestSquashIdRanges(t *testing.T) {
	input := []IdRange{
		{min: 0, max: 10},
		{min: 5, max: 15},
		{min: 12, max: 20},
	}

	expected := []IdRange{
		{min: 0, max: 20},
	}

	result := SquashIdRanges(input)
	assert.EqualValues(t, expected, result)

	input = []IdRange{
		{min: 20, max: 40},
		{min: 30, max: 50},
		{min: 40, max: 60},
		{min: 50, max: 70},
	}

	expected = []IdRange{
		{min: 20, max: 70},
	}

	result = SquashIdRanges(input)
	assert.EqualValues(t, expected, result)

	input = []IdRange{
		{min: 20, max: 40},
		{min: 30, max: 50},
		{min: 140, max: 160},
		{min: 150, max: 170},
	}

	expected = []IdRange{
		{min: 20, max: 50},
		{min: 140, max: 170},
	}

	result = SquashIdRanges(input)
	assert.EqualValues(t, expected, result)
}

func TestSquashIdRangesOnce(t *testing.T) {
	input := []IdRange{
		{min: 0, max: 10},
		{min: 5, max: 15},
		{min: 12, max: 20},
	}

	expected := []IdRange{
		{min: 0, max: 15},
		{min: 12, max: 20},
	}

	result, n := SquashIdRangesOnce(input)
	assert.Equal(t, 1, n)
	assert.EqualValues(t, expected, result)

	input = []IdRange{
		{min: 20, max: 40},
		{min: 30, max: 50},
		{min: 40, max: 60},
		{min: 50, max: 70},
	}

	expected = []IdRange{
		{min: 20, max: 50},
		{min: 40, max: 70},
	}

	result, n = SquashIdRangesOnce(input)
	assert.Equal(t, 2, n)
	assert.EqualValues(t, expected, result)

	input = []IdRange{
		{min: 20, max: 40},
		{min: 30, max: 50},
		{min: 140, max: 160},
		{min: 150, max: 170},
	}

	expected = []IdRange{
		{min: 20, max: 50},
		{min: 140, max: 170},
	}

	result, n = SquashIdRangesOnce(input)
	assert.Equal(t, 2, n)
	assert.EqualValues(t, expected, result)
}
