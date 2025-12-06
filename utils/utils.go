package utils

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

func Filter[T any](items []T, f func(i int, item T) bool) []T {
	var out []T
	for i, item := range items {
		if f(i, item) {
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
