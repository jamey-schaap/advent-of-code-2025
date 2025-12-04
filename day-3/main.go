package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const numberSize = 2

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sum, number := 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		number, err = FindHighestNumber(text, numberSize)
		sum += number
	}
	fmt.Println(sum)
}

func FindHighestNumber(text string, size int) (int, error) {
	chars := make([]uint8, size)
	curr := text
	for i := range size {
		c, idx, err := FindHighestChar(curr, size-i-1)
		if err != nil {
			return 0, err
		}
		chars[i] = c
		curr = curr[idx+1:]
	}

	return strconv.Atoi(string(chars))
}

func FindHighestChar(text string, rightNeighbours int) (char uint8, idx int, err error) {
	if len(text) == 0 {
		return 0, 0, fmt.Errorf("string shouldn't be empty")
	}

	if len(text) == 1 {
		return text[0], 0, nil
	}

	for i := range text {
		c := text[i]
		if i+rightNeighbours <= len(text)-1 && c > char {
			char = c
			idx = i
		}
	}

	return char, idx, nil
}
