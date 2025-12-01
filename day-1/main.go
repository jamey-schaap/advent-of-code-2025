package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	upperBoundary = 99
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var passedZeroCount int
	pos := 50
	passwordA := 0
	passwordB := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		pos, passedZeroCount, err = Dial(pos, dir, clicks)
		if err != nil {
			panic(err)
		}

		if pos == 0 {
			passwordA++
			passwordB++
		}
		passwordB += passedZeroCount
	}

	fmt.Printf("Password A: %d\n", passwordA)
	fmt.Printf("Password B: %d\n", passwordB)
}

func Dial(pos int, dir uint8, clicks int) (remainder, quotient int, err error) {
	const amountOfPossibleValues = upperBoundary + 1

	offset := clicks % amountOfPossibleValues
	switch dir {
	case 'L':
		quotient += (clicks + amountOfPossibleValues) / amountOfPossibleValues
		remainder = (pos - offset + amountOfPossibleValues) % amountOfPossibleValues
	case 'R':
		quotient = (pos + clicks) / amountOfPossibleValues
		remainder = (pos + offset) % amountOfPossibleValues
	default:
		return 0, quotient, fmt.Errorf("unexpected char '%d'", dir)
	}

	return remainder, quotient, nil
}
