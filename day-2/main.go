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
		return !IsValidId(s)
	})
	return invalidIds
}

func IsValidId(id string) bool {
	if len(id) == 0 {
		return false
	}

	if id[0] == '0' {
		return false
	}

	middleIdx := len(id) / 2
	for size := 1; size <= middleIdx; size++ {
		if len(id)%size != 0 {
			continue
		}

		parts, _ := EvenlySplitString(id, size)
		if AllEqual(parts...) {
			return false
		}
	}

	return true
}

func EvenlySplitString(str string, size int) ([]string, error) {
	if size == 0 {
		return nil, fmt.Errorf("invalid size: %d", size)
	}

	if len(str)%size != 0 {
		return nil, fmt.Errorf("string of len %d cannot be evenly dived into substrings of len %d", len(str), size)
	}

	count := len(str) / size
	var parts []string
	for i := 0; i < count; i++ {
		idx := i * size
		substr := str[idx : idx+size]
		parts = append(parts, substr)
	}
	return parts, nil
}

func GetRange(start, end int) []int {
	return GetRangeWithStep(start, end, 1)
}

func GetRangeWithStep(start, end, step int) []int {
	if start > end {
		start, end = end, start
	}

	r := make([]int, 0, end-start+1)
	for i := start; i < end; i += step {
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

func AllEqual[T comparable](values ...T) bool {
	if len(values) < 2 {
		return true
	}

	first := values[0]
	for _, v := range values[1:] {
		if v != first {
			return false
		}
	}
	return true
}
