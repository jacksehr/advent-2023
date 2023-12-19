package solution_14

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

const TEST_INPUT = `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

func col(grid []string, x int) string {
	sb := strings.Builder{}

	for _, row := range grid {
		sb.WriteByte(row[x])
	}

	return sb.String()
}

// first row becomes last column
func rotateCW(grid []string) []string {
	rotated := []string{}
	for x := 0; x < len(grid[0]); x++ {
		sb := strings.Builder{}
		for y := len(grid); y >= 0; y-- {
			sb.WriteByte(grid[y][x])
		}

		rotated = append(rotated, sb.String())
	}

	return rotated
}

func rotateColsCW(cols []string) []string {
	rotated := []string{}
	for _, col := range cols {
		colAsRunes := []rune(col)
		slices.Reverse(colAsRunes)
		rotated = append(rotated, string(colAsRunes))
	}

	return rotated
}

func tilt(grid []string) []string {
	cols := []string{}

	for x := 0; x < len(grid[0]); x++ {
		cols = append(cols, col(grid, x))
	}

	pattern := regexp.MustCompile(`\.+$`)

	for colIndex, c := range cols {
		col := c

		for i := 0; i < len(c); i++ {
			if col[i] == 'O' {
				match := pattern.FindString(col[:i])

				n := len(match)
				if n == 0 {
					continue
				}

				toReplace := []rune(col)

				toReplace[i], toReplace[i-n] = toReplace[i-n], toReplace[i]

				col = string(toReplace)
			}
		}

		cols[colIndex] = col
	}

	return cols
}

func partOne(inputLines []string) {
	// inputLines = strings.Split(strings.TrimSpace(TEST_INPUT), "\n")

	cols := tilt(inputLines)

	weightSum := 0
	for _, col := range cols {
		for w, c := range col {
			if c == 'O' {
				weightSum += len(col) - w
			}
		}
	}

	fmt.Println("Weight sum is", weightSum)
}

func partTwo(inputLines []string) {
	// inputLines = strings.Split(strings.TrimSpace(TEST_INPUT), "\n")

	m := map[string]int{}

	c0, c1 := 0, 0

	var cols []string
	grid := inputLines
	m[strings.Join(grid, "\n")] = 0

	max := 1_000_000_000

	for i := 0; i < max; i++ {
		for range "NWSE" {
			cols = tilt(grid)
			grid = rotateColsCW(cols)
		}

		if seenAt, ok := m[strings.Join(grid, "\n")]; ok {
			c0 = seenAt
			c1 = i + 1
			break
		}

		m[strings.Join(grid, "\n")] = i + 1
	}

	length := c1 - c0
	lookup := c0 + (max-c1)%length

	for k, v := range m {
		if v == lookup {
			g := strings.Split(k, "\n")
			cols := []string{}

			for x := 0; x < len(g[0]); x++ {
				cols = append(cols, col(g, x))
			}

			weightSum := 0
			for _, col := range cols {
				for w, c := range col {
					if c == 'O' {
						weightSum += len(col) - w
					}
				}
			}

			fmt.Println("Weight sum is", weightSum)
		}
	}
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
