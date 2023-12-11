package solution_10

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

const TEST_INPUT = `
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`

type Empty struct{}

type Pos struct {
	x    int
	y    int
	from *Pos
}
type PosSet map[Pos]struct{}

func findOrigin(grid []string) Pos {
	for y, row := range grid {
		for x, c := range row {
			if c == 'S' {
				return Pos{x, y, nil}
			}
		}
	}

	return Pos{0, 0, nil}
}

func pop[T any](stack []T) (T, []T) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

type Direction int

const (
	U Direction = iota
	D
	L
	R
	None
)

func getDirection(p *Pos) Direction {
	if p.from == nil {
		return None
	}

	switch {
	case p.from.x == p.x && p.from.y-1 == p.y:
		return U
	case p.from.x == p.x && p.from.y+1 == p.y:
		return D
	case p.from.y == p.y && p.from.x-1 == p.x:
		return L
	case p.from.y == p.y && p.from.x+1 == p.x:
		return R
	}

	return None
}

func getAdjPipes(pos Pos, grid []string) []Pos {
	var (
		up    = Pos{pos.x, pos.y - 1, &pos}
		down  = Pos{pos.x, pos.y + 1, &pos}
		left  = Pos{pos.x - 1, pos.y, &pos}
		right = Pos{pos.x + 1, pos.y, &pos}
	)

	dir := getDirection(&pos)
	cell := grid[pos.y][pos.x]

	possible := []Pos{}
	switch cell {
	case 'S':
		possible = append(possible, up, down, left, right)
	case '|':
		switch dir {
		case U:
			possible = append(possible, up)
		case D:
			possible = append(possible, down)
		}
	case '-':
		switch dir {
		case L:
			possible = append(possible, left)
		case R:
			possible = append(possible, right)
		}
	case 'F':
		switch dir {
		case L:
			possible = append(possible, down)
		case U:
			possible = append(possible, right)
		}
	case '7':
		switch dir {
		case R:
			possible = append(possible, down)
		case U:
			possible = append(possible, left)
		}
	case 'L':
		switch dir {
		case D:
			possible = append(possible, right)
		case L:
			possible = append(possible, up)
		}
	case 'J':
		switch dir {
		case D:
			possible = append(possible, left)
		case R:
			possible = append(possible, up)
		}
	}

	filtered := slices.DeleteFunc(possible, func(p Pos) bool {
		return p.x < 0 || p.x >= len(grid[0]) || p.y < 0 || p.y >= len(grid) || grid[p.y][p.x] == '.'
	})

	return filtered
}

func getLoop(grid []string) (PosSet, Pos) {
	origin := findOrigin(grid)
	visited := map[Pos]struct{}{}
	stack := getAdjPipes(origin, grid)

	var curr Pos
	for len(stack) > 0 {
		curr, stack = pop(stack)

		if grid[curr.y][curr.x] == 'S' {
			break
		}

		if _, seen := visited[curr]; seen {
			continue
		}

		nextPositions := getAdjPipes(curr, grid)

		stack = append(stack, nextPositions...)

		visited[curr] = struct{}{}
	}

	end := curr

	loop := PosSet{}
	for curr.from != nil {
		cleaned := Pos{curr.x, curr.y, nil}
		loop[cleaned] = Empty{}
		curr = *curr.from
	}

	return loop, end
}

func partOne(inputLines []string) {
	loop, _ := getLoop(inputLines)

	fmt.Println("Furthest point is", math.Ceil(float64(len(loop))/2), "away")
}

func scaleGrid(grid []string) []string {
	char := "_"

	enlarged := slices.Clone(grid)

	for y, row := range grid {
		enlarged[y] = strings.Join(strings.Split(row, ""), char)
	}

	return strings.Split(strings.Join(enlarged, "\n"+strings.Repeat(char, len(enlarged[0]))+"\n"), "\n")
}

func partTwo(inputLines []string) {
	// inputLines = strings.Split(strings.TrimSpace(TEST_INPUT), "\n")

	enlarged := scaleGrid(inputLines)

	_, end := getLoop(inputLines)
	newLoop := PosSet{}

	for end.from != nil {
		newLoop[Pos{2 * end.x, 2 * end.y, nil}] = Empty{}

		if end.x == end.from.x {
			yToFill := (2*end.y + 2*end.from.y) / 2
			rowAsRunes := []rune(enlarged[yToFill])
			rowAsRunes[2*end.x] = '|'
			enlarged[yToFill] = string(rowAsRunes)

			newLoop[Pos{2 * end.x, yToFill, nil}] = Empty{}
		} else {
			xToFill := (2*end.x + 2*end.from.x) / 2
			rowAsRunes := []rune(enlarged[2*end.y])
			rowAsRunes[xToFill] = '-'
			enlarged[2*end.y] = string(rowAsRunes)

			newLoop[Pos{xToFill, 2 * end.y, nil}] = Empty{}
		}

		end = *end.from
	}

	painted := make([][]rune, len(enlarged))
	for y := range enlarged {
		painted[y] = make([]rune, len(enlarged[0]))
	}

	visited := PosSet{}

	// haphazardly assume that S is not at the origin, for general use you couldn't do this
	curr := Pos{0, 0, nil}
	floodStack := []Pos{curr}
	for len(floodStack) > 0 {
		curr, floodStack = pop(floodStack)

		if curr.x < 0 || curr.x >= len(enlarged[0]) || curr.y < 0 || curr.y >= len(enlarged) {
			continue
		}

		if _, isLoop := newLoop[curr]; isLoop {
			continue
		}

		if _, seen := visited[curr]; seen {
			continue
		}

		visited[curr] = Empty{}

		painted[curr.y][curr.x] = '0'

		var (
			up    = Pos{curr.x, curr.y - 1, nil}
			down  = Pos{curr.x, curr.y + 1, nil}
			left  = Pos{curr.x - 1, curr.y, nil}
			right = Pos{curr.x + 1, curr.y, nil}
		)
		floodStack = append(floodStack, up, down, left, right)
	}

	tilesFound := 0
	for y, row := range painted {
		for x, c := range row {
			if c == '0' || enlarged[y][x] == '_' {
				continue
			}

			if _, inLoop := newLoop[Pos{x, y, nil}]; inLoop {
				continue
			}

			tilesFound++
		}
	}

	fmt.Println(tilesFound, "tiles found")
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
