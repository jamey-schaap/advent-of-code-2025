package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	answer, err := GetAnswer(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}

type IdRange struct {
	min, max int
}

func GetAnswer(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	var ranges []IdRange
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		idRange, err := ParseRangeFromString(line)
		if err != nil {
			return 0, nil
		}
		ranges = append(ranges, *idRange)
	}

	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()

		n, err := strconv.Atoi(line)
		if err != nil {
			return 0, nil
		}

		for _, r := range ranges {
			if n >= r.min && n <= r.max {
				cnt++
				break
			}
		}
	}
	return cnt, nil
}

func ParseRangeFromString(text string) (*IdRange, error) {
	parts := strings.Split(text, "-")
	minStr, maxStr := parts[0], parts[1]
	minId, err := strconv.Atoi(minStr)
	if err != nil {
		return nil, err
	}

	maxId, err := strconv.Atoi(maxStr)
	if err != nil {
		return nil, err
	}

	return &IdRange{minId, maxId}, nil
}
