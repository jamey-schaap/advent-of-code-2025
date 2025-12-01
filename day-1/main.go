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

		offset := clicks % (upperBoundary + 1)
		switch dir {
		case 'L':
			pos = (pos - offset) % (upperBoundary + 1)
		case 'R':
			pos = (pos + offset) % (upperBoundary + 1)
		default:
			panic(fmt.Errorf("unexpected char '%d'", dir))
		}
		if pos == 0 {
			pwd++
		}
	}

	fmt.Printf("Password: %d\n", pwd)
}
