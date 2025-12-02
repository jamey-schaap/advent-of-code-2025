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

	var passedZeroCount int
	pos := 50
	passwordA := 0
	passwordB := 0

	scanner := bufio.NewScanner(f)
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
		shift := (amountOfPossibleValues - pos) % amountOfPossibleValues
		quotient = (clicks + shift) / amountOfPossibleValues
		remainder = (pos - offset + amountOfPossibleValues) % amountOfPossibleValues
	case 'R':
		quotient = (pos + clicks) / amountOfPossibleValues
		remainder = (pos + offset) % amountOfPossibleValues

	default:
		return 0, 0, fmt.Errorf("unexpected char '%d'", dir)
	}

	return remainder, quotient, nil
}
