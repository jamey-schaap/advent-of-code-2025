package main

import (
	"fmt"
	"io"
	"log"
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

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	answer := GetAnswer(string(b))
	fmt.Println(answer)
}

func GetAnswer(text string) int {
	idRanges := strings.Split(text, ",")
	sum := 0
	for _, r := range idRanges {
		invalidIds := GetInvalidIds(r)

		for _, id := range invalidIds {
			sum += id
		}
	}

	return sum
}

func GetInvalidIds(r string) []int {
	parts := strings.Split(r, "-")
	if len(parts) != 2 {
		log.Fatalf("Invalid range: %s", r)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	idRange := GetRange(start, end+1)
	invalidIds := Filter(idRange, func(id int) bool {
		s := strconv.Itoa(id)
		return !isValidId(s)
	})
	return invalidIds
}

func isValidId(id string) bool {
	if len(id) == 0 {
		return false
	}

	if id[0] == '0' {
		return false
	}

	if len(id)%2 != 0 {
		return true
	}

	middleIdx := len(id) / 2
	part1 := id[:middleIdx]
	part2 := id[middleIdx:]

	return part1 != part2
}

func GetRange(start, end int) []int {
	if start > end {
		start, end = end, start
	}

	r := make([]int, 0, end-start+1)
	for i := start; i < end; i++ {
		r = append(r, i)
	}
	return r
}

func Filter[T any](items []T, f func(item T) bool) []T {
	var out []T
	for _, item := range items {
		if f(item) {
			out = append(out, item)
		}
	}
	return out
}
