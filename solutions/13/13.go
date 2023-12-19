package solution_13

import (
	"fmt"
	"slices"
	"strings"
)

// const TEST_INPUT = `
// #.##..##.
// ..#.##.#.
// ##......#
// ##......#
// ..#.##.#.
// ..##..##.
// #.#.##.#.

// #...##..#
// #....#..#
// ..##..###
// #####.##.
// #####.##.
// ..##..###
// #....#..#
// `

const TEST_INPUT = `
.##.#.....#..
......#...###
.##..#..#....
.##.##..#....
.......#...##
.##..#.##...#
.....####..##
#..#.#..#####
####.#.#.####
.##.#.##..#..
.##.#####.#..
#..#.##..##..
####.###..#..
`

type Dir int

const (
	H Dir = iota
	V
)

type mirror struct {
	direction Dir
	lower     int
	upper     int
}

func (m *mirror) dist() int {
	multi := 1
	if m.direction == H {
		multi = 100
	}

	return multi * (m.lower + ((m.upper - m.lower) / 2) + 1)
}

func (m *mirror) size() int {
	return m.upper - m.lower
}

func (m *mirror) print(grid []string) {
	switch m.direction {
	case H:
		fmt.Println(strings.Join(grid[:m.lower], "\n"))
		fmt.Println(strings.Repeat("-", len(grid[0])))
		fmt.Println(strings.Join(grid[m.lower:m.upper+1], "\n"))
		fmt.Println(strings.Repeat("-", len(grid[0])))
		fmt.Println(strings.Join(grid[m.upper+1:], "\n"))
	case V:
		for _, row := range grid {
			fmt.Print(row[:m.lower])
			fmt.Print("|")
			fmt.Print(row[m.lower : m.upper+1])
			fmt.Print("|")
			fmt.Println(row[m.upper+1:])
		}
	}
}

func col(grid []string, x int) string {
	sb := strings.Builder{}

	for _, row := range grid {
		sb.WriteByte(row[x])
	}

	return sb.String()
}

func getLargestMirror(grid []string, toExclude *mirror) *mirror {
	hm := &mirror{H, 0, 0}
	vm := &mirror{V, 0, 0}

	for y := range grid {
		if y == len(grid)-1 {
			break
		}

		t, b := y, y+1
		for t >= 0 && b < len(grid) {
			r, r1 := grid[t], grid[b]
			if r != r1 {
				break
			}

			t--
			b++
		}

		if y > t && (t == -1 || b == len(grid)) {
			curr := &mirror{H, t + 1, b - 1}
			if toExclude != nil && curr.direction == toExclude.direction && curr.lower == toExclude.lower && curr.upper == toExclude.upper {
				continue
			}

			if curr.size() > hm.size() {
				hm = curr
			}
		}
	}

	if hm.size() > 0 {
		return hm
	}

	for x := 0; x < len(grid[0])-1; x++ {
		l, r := x, x+1
		for l >= 0 && r < len(grid[0]) {
			c, c1 := col(grid, l), col(grid, r)
			if c != c1 {
				break
			}

			l--
			r++
		}

		if x > l && (l == -1 || r == len(grid[0])) {
			curr := &mirror{V, l + 1, r - 1}

			if toExclude != nil && curr.direction == toExclude.direction && curr.lower == toExclude.lower && curr.upper == toExclude.upper {
				continue
			}

			if curr.size() > vm.size() {
				vm = curr
			}
		}
	}

	return vm
}

func partOne(grids [][]string) {
	summary := 0

	for _, grid := range grids {
		m := getLargestMirror(grid, nil)

		summary += m.dist()
	}

	fmt.Println("Summary is", summary)
}

func partTwo(grids [][]string) {
	summary := 0

	for _, grid := range grids {
		m0 := getLargestMirror(grid, nil)

	gridloop:
		for y, row := range grid {
			for x := range row {
				newGrid := slices.Clone(grid)
				rowToReplace := newGrid[y]
				switch rowToReplace[x] {
				case '#':
					newGrid[y] = rowToReplace[:x] + "." + rowToReplace[x+1:]
				case '.':
					newGrid[y] = rowToReplace[:x] + "#" + rowToReplace[x+1:]
				}

				var m *mirror
				if y == 5 && x == 11 {
					m = getLargestMirror(newGrid, m0)
				} else {
					m = getLargestMirror(newGrid, m0)
				}

				if m.size() != 0 {
					summary += m.dist()
					break gridloop
				}
			}
		}
	}

	fmt.Println("Summary is", summary)
}

func Solution(inputLines []string) {
	// inputLines = strings.Split(strings.TrimSpace(TEST_INPUT), "\n")
	inputLines = strings.Split(strings.Join(inputLines, "\n"), "\n\n")
	splitLines := [][]string{}

	for _, l := range inputLines {
		splitLines = append(splitLines, strings.Split(l, "\n"))
	}

	fmt.Println("Part one")
	partOne(splitLines)

	fmt.Println("Part two")
	partTwo(splitLines)
}
