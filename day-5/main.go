package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//answer, err := GetAnswerPart1(f)
	answer, err := GetAnswerPart2(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
}

type IdRange struct {
	min, max int
}

func GetAnswerPart1(r io.Reader) (int, error) {
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

func GetAnswerPart2(r io.Reader) (int, error) {
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

	ranges = SquashIdRanges(ranges)

	result := 0
	for _, r := range ranges {
		result += r.max - r.min + 1 // to account for the upper boundary which is inclusive
	}

	return result, nil
}

func SquashIdRanges(ranges []IdRange) []IdRange {
	var out []IdRange
	var squashedCnt int
	in := ranges

	for {
		out, squashedCnt = SquashIdRangesOnce(in)
		if squashedCnt == 0 {
			break
		}
		in = out
	}

	return out
}

func SquashIdRangesOnce(ranges []IdRange) ([]IdRange, int) {
	out := make([]IdRange, 0)
	checkedIndexes := make([]int, 0)
	squashedCnt := 0

	for idxA, a := range ranges {
		rangeToAdd := a
		if slices.Contains(checkedIndexes, idxA) {
			continue
		}

		for idxB, b := range ranges {
			if idxA == idxB || slices.Contains(checkedIndexes, idxB) {
				continue
			}

			// a.min in b.min .. b.max or b.min in a.min .. a.max
			if (b.min <= a.min && a.min <= b.max) ||
				a.min <= b.min && b.min <= a.max {
				rangeToAdd = IdRange{
					min: min(a.min, b.min),
					max: max(a.max, b.max),
				}
				checkedIndexes = append(checkedIndexes, idxB)
				squashedCnt++
				break
			}
		}

		out = append(out, rangeToAdd)
		checkedIndexes = append(checkedIndexes, idxA)
	}

	return out, squashedCnt
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
