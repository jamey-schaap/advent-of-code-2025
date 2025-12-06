package main

import (
	"advent-of-code-2025/utils"
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

	idRange := utils.GetRange(start, end+1)
	invalidIds := utils.Filter(idRange, func(_, id int) bool {
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
		if utils.AllEqual(parts...) {
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
