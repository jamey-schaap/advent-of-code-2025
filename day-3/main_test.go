package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHighestNumber(t *testing.T) {
	actual := FindHighestNumber("987654321111111", 2)
	assert.Equal(t, "98", actual)
}
