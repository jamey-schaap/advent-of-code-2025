package main

import "fmt"

func main() {
	//f, err := os.Open("./input.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//scanner := bufio.NewScanner(f)
	//for scanner.Scan() {
	//	scanner.Text()
	//}
	res := FindHighestNumber("1237193", 3)
	fmt.Println(res)

	res = FindHighestNumber("12937193", 3)
	fmt.Println(res)
	// 72571894
	// 725
	// 727
	// -8
	// 778
	// 78...
	// 8...

	// -9
	// 789
	// 89...
	// 9...

	// from i -> [n] numbers are higher
	// find the highest i, with [size] numbers to the right
	// then repeat for x := range [size], then search for [size-1] to the right

	// 3257597823087123
	//      ^ finds the highest number in the string with at least [size] to the right
	// 3257597823087123
	//        ^
	// 3257597823087123
	//            ^
	// does it find the next highest (and first? i.e. ....8....8....)
}

func FindHighestNumber(text string, size int) string {
	if len(text) == size {
		return text
	}

	highestNumber := make([]int32, size)

	for idxChar, char := range text {
		for idxHighest, v := range highestNumber {
			if char <= v {
				continue
			}

			if idxHighest == 0 && idxChar == len(text)-size {
				return text[idxChar:]
			}

			highestNumber[idxHighest] = char
			// highestNumber[idxHighest+1:] reset all to -1
			break
		}
	}

	return string(highestNumber)
}
