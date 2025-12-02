package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnswer(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	actual := GetAnswer(input)
	assert.Equal(t, 1227775554, actual)
}

func TestGetInvalidIds(t *testing.T) {
	actual := GetInvalidIds("11-22")
	assert.Equal(t, []int{11, 22}, actual)

	actual = GetInvalidIds("99-115")
	assert.Equal(t, []int{99}, actual)

	actual = GetInvalidIds("998-1012")
	assert.Equal(t, []int{1010}, actual)

	actual = GetInvalidIds("1188511880-1188511890")
	assert.Equal(t, []int{1188511885}, actual)

	actual = GetInvalidIds("222220-222224")
	assert.Equal(t, []int{222222}, actual)

	actual = GetInvalidIds("1698522-1698528")
	assert.Len(t, actual, 0)

	actual = GetInvalidIds("446443-446449")
	assert.Equal(t, []int{446446}, actual)

	actual = GetInvalidIds("38593856-38593862")
	assert.Equal(t, []int{38593859}, actual)
}

func TestGetRange(t *testing.T) {
	start := 11
	end := 22
	var expected []int
	for i := start; i <= end; i++ {
		expected = append(expected, i)
	}

	actual := GetRange(start, end+1)
	assert.EqualValues(t, expected, actual)
}
