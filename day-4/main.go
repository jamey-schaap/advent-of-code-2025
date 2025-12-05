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

	cnt := 0
	for {
		c := GetAccessibleCount(grid, 1, 4)
		if c == 0 {
			break
		}
		cnt += c
	}
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
				grid[y][x] = byte('.')
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
	xs := MakeAxis(c.x, maxX, radius)
	ys := MakeAxis(c.y, maxY, radius)

	product := CartesianProduct(xs, ys)
	coords := make([]Coordinate, 0, len(product)-1)
	for _, p := range product {
		x, y := p[0], p[1]
		if x == c.x && y == c.y {
			continue
		}
		coords = append(coords, Coordinate{x, y})
	}
	return coords
}

func MakeAxis(center, max, radius int) []int {
	out := make([]int, 0, 2*radius+1)
	for offset := -radius; offset <= radius; offset++ {
		v := center + offset
		if v >= 0 && v <= max {
			out = append(out, v)
		}
	}
	return out
}

func CartesianProduct[T any](s1, s2 []T) [][]T {
	capacity := len(s1) * len(s2)
	product := make([][]T, 0, capacity)
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			product = append(product, []T{v1, v2})
		}
	}
	return product
}
