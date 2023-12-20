package solution_16

import (
	"fmt"
	"log"
)

const TEST_INPUT = `
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`

type Direction int

const (
	U Direction = iota
	D
	L
	R
)

type Pos struct {
	x int
	y int

	from *Pos
}

type Key struct {
	x   int
	y   int
	dir Direction
}

func (p *Pos) direction() Direction {
	if p.from == nil {
		return R
	}

	// no diagonal movement is allowed
	dx := p.x - p.from.x
	switch {
	case dx > 0:
		return R
	case dx < 0:
		return L
	}

	dy := p.y - p.from.y
	switch {
	case dy > 0:
		return D
	case dy < 0:
		return U
	}

	log.Fatalf("This direction makes no sense: %v", p)

	return R
}

func pop[T any](stack []T) (T, []T) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func goNext(cell rune, curr *Pos) []*Pos {
	dir := curr.direction()

	var (
		up    = &Pos{curr.x, curr.y - 1, curr}
		down  = &Pos{curr.x, curr.y + 1, curr}
		left  = &Pos{curr.x - 1, curr.y, curr}
		right = &Pos{curr.x + 1, curr.y, curr}
	)

	switch cell {
	case '.':
		switch dir {
		case U:
			return []*Pos{up}
		case D:
			return []*Pos{down}
		case L:
			return []*Pos{left}
		case R:
			return []*Pos{right}
		}
	case '|':
		switch dir {
		case U:
			return []*Pos{up}
		case D:
			return []*Pos{down}
		}

		return []*Pos{up, down}
	case '-':
		switch dir {
		case L:
			return []*Pos{left}
		case R:
			return []*Pos{right}
		}

		return []*Pos{left, right}
	case '\\':
		switch dir {
		case U:
			return []*Pos{left}
		case D:
			return []*Pos{right}
		case L:
			return []*Pos{up}
		case R:
			return []*Pos{down}
		}
	case '/':
		switch dir {
		case U:
			return []*Pos{right}
		case D:
			return []*Pos{left}
		case L:
			return []*Pos{down}
		case R:
			return []*Pos{up}
		}
	}

	return []*Pos{}
}

func cellsCovered(grid []string, start *Pos) int {
	seen := map[Key]struct{}{}
	stack := []*Pos{start}

	var curr *Pos
	for len(stack) > 0 {
		curr, stack = pop(stack)

		if curr.y < 0 || curr.y >= len(grid) || curr.x < 0 || curr.x >= len(grid[0]) {
			continue
		}

		cacheKey := Key{curr.x, curr.y, curr.direction()}
		if _, ok := seen[cacheKey]; ok {
			continue
		} else {
			seen[cacheKey] = struct{}{}
		}

		stack = append(stack, goNext(rune(grid[curr.y][curr.x]), curr)...)
	}

	uniqueCells := map[[2]int]struct{}{}
	for k := range seen {
		uniqueCells[[2]int{k.x, k.y}] = struct{}{}
	}

	return len(uniqueCells)
}

func partOne(inputLines []string) {
	fmt.Println("Saw", cellsCovered(inputLines, &Pos{0, 0, nil}), "uniques")
}

func partTwo(inputLines []string) {
	maxCovered := 0

	topRow := inputLines[0]
	for x := range topRow {
		maxCovered = max(maxCovered, cellsCovered(inputLines, &Pos{x, 0, &Pos{x, -1, nil}}))
		maxCovered = max(maxCovered, cellsCovered(inputLines, &Pos{x, len(inputLines) - 1, &Pos{x, len(inputLines), nil}}))
	}

	for y := range inputLines {
		maxCovered = max(maxCovered, cellsCovered(inputLines, &Pos{0, y, &Pos{-1, y, nil}}))
		maxCovered = max(maxCovered, cellsCovered(inputLines, &Pos{len(inputLines[0]) - 1, y, &Pos{len(inputLines[0]), y, nil}}))
	}

	fmt.Println("Max is", maxCovered, "uniques")
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
