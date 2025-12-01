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
	pos := 50
	pwd := 0
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

		pos, err = Dial(pos, dir, clicks)
		if err != nil {
			panic(err)
		}

		if pos == 0 {
			pwd++
		}
	}

	fmt.Printf("Password: %d\n", pwd)
}

func Dial(pos int, dir uint8, clicks int) (int, error) {
	const amountOfPossibleValues = upperBoundary + 1

	offset := clicks % amountOfPossibleValues
	switch dir {
	case 'L':
		pos = (pos - offset + amountOfPossibleValues) % amountOfPossibleValues
	case 'R':
		pos = (pos + offset) % amountOfPossibleValues
	default:
		return 0, fmt.Errorf("unexpected char '%d'", dir)
	}

	return pos, nil
}
