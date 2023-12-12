package solution_11

import (
	"fmt"
	"slices"
	"strings"
)

type EmptySpace struct {
	rows []int
	cols []int
}

func findEmpty(grid []string) EmptySpace {
	es := EmptySpace{[]int{}, []int{}}

	emptyRow := strings.Repeat(".", len(grid[0]))

	for y, row := range grid {
		if row == emptyRow {
			es.rows = append(es.rows, y)
		}
	}

	for x := 0; x < len(grid[0]); x++ {
		colIsEmpty := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] != '.' {
				colIsEmpty = false
			}
		}

		if colIsEmpty {
			es.cols = append(es.cols, x)
		}
	}

	return es
}

func getPairDistanceSum(grid []string, emptyScale int) int {
	hashes := [][2]int{}

	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				hashes = append(hashes, [2]int{x, y})
			}
		}
	}

	emptySpace := findEmpty(grid)

	pathSum := 0
	for i := 0; i < len(hashes); i++ {
		for _, pairedHash := range hashes[i:] {
			x := []int{pairedHash[0], hashes[i][0]}
			y := []int{pairedHash[1], hashes[i][1]}

			slices.Sort(x)
			slices.Sort(y)

			xDist := 0
			for i := x[0]; i < x[1]; i++ {
				if slices.Contains(emptySpace.cols, i) {
					xDist += emptyScale
				} else {
					xDist += 1
				}
			}

			yDist := 0
			for i := y[0]; i < y[1]; i++ {
				if slices.Contains(emptySpace.rows, i) {
					yDist += emptyScale
				} else {
					yDist += 1
				}
			}

			pathSum += xDist + yDist
		}
	}

	return pathSum
}

func partOne(inputLines []string) {
	pathSum := getPairDistanceSum(inputLines, 2)

	fmt.Println("Path length sum is", pathSum)
}

func partTwo(inputLines []string) {
	pathSum := getPairDistanceSum(inputLines, 1_000_000)

	fmt.Println("Path length sum is", pathSum)
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
