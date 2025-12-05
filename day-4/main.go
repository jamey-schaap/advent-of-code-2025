package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	answer, err := GetAnswer(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer)
}

func GetAnswer(r io.Reader) (int, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	grid := bytes.Split(b, []byte("\n"))
	cnt := GetAccessibleCount(grid, 1, 4)
	return cnt, nil
}

func GetAccessibleCount(grid [][]byte, radius, threshold int) int {
	const toiletRoll = byte('@')

	accessibleCnt := 0
	for y := range grid {
		row := grid[y]
		for x := range row {
			cell := row[x]
			if cell != toiletRoll {
				continue
			}

			c := Coordinate{x, y}
			coords := GetNeighborCoordinates(c, radius, len(row)-1, len(grid)-1)

			cnt := 0
			for _, c := range coords {
				if grid[c.y][c.x] == toiletRoll {
					cnt++
				}
			}

			if cnt < threshold {
				accessibleCnt++
			}
		}
	}

	return accessibleCnt
}

type Coordinate struct {
	x, y int
}

func GetNeighborCoordinates(c Coordinate, radius, maxX, maxY int) []Coordinate {
	var xs, ys []int
	for r := range radius {
		offset := r + 1
		xs = append(xs, c.x-offset, c.x+offset)
		ys = append(ys, c.y-offset, c.y+offset)
	}

	xs = append(xs, c.x)
	xs = Filter(xs, func(x int) bool {
		return x >= 0 && x <= maxX
	})

	ys = append(ys, c.y)
	ys = Filter(ys, func(y int) bool {
		return y >= 0 && y <= maxY
	})

	product := CartesianProduct(xs, ys)
	coords := make([]Coordinate, len(product)-1)
	i := 0
	for _, p := range product {
		x := p[0]
		y := p[1]

		if x == c.x && y == c.y {
			continue
		}
		coords[i] = Coordinate{x, y}
		i++
	}
	return coords
}

func CartesianProduct[T any](s1, s2 []T) [][]T {
	capacity := len(s1) * len(s2)
	product := make([][]T, capacity)
	i := 0

	for _, v1 := range s1 {
		for _, v2 := range s2 {
			product[i] = []T{v1, v2}
			i++
		}
	}
	return product
}

func Filter[T any](items []T, f func(T) bool) []T {
	var out []T
	for _, item := range items {
		if f(item) {
			out = append(out, item)
		}
	}
	return out
}
