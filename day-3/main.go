package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const numberSize = 12

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	answer, err := GetAnswer(f, numberSize)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}

func GetAnswer(r io.Reader, numberSize int) (sum int, err error) {
	number := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		number, err = FindHighestNumber(text, numberSize)
		if err != nil {
			return 0, err
		}
		sum += number
	}

	return sum, nil
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

	const char9 = uint8('9')
	for i := range text {
		c := text[i]
		if i+rightNeighbours > len(text)-1 || c <= char {
			continue
		}

		char = c
		idx = i
		if c == char9 {
			return char, idx, nil
		}
	}

	return char, idx, nil
}
