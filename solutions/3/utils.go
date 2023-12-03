package solution_3

const TEST_INPUT = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func hasAdjSymbol(board []string, xL int, xR int, y int) bool {
	for rowY := y - 1; rowY <= y+1; rowY++ {
		if rowY < 0 || rowY > len(board)-1 {
			continue
		}

		for colX := xL - 1; colX < xR+1; colX++ {
			if colX < 0 || colX > len(board[rowY])-1 {
				continue
			}

			c := board[rowY][colX]

			if !isDigit(rune(c)) && c != '.' {
				return true
			}
		}
	}

	return false
}

type coord struct {
	Row int
	Col int
}

type numLoc struct {
	L int
	R int
}

type numCoord struct {
	Row int
	Col numLoc
}

func getAdjGears(board []string, xL int, xR int, y int) []coord {
	attachedGears := []coord{}

	for rowY := y - 1; rowY <= y+1; rowY++ {
		if rowY < 0 || rowY > len(board)-1 {
			continue
		}

		for colX := xL - 1; colX < xR+1; colX++ {
			if colX < 0 || colX > len(board[rowY])-1 {
				continue
			}

			c := board[rowY][colX]

			if c == '*' {
				attachedGears = append(attachedGears, coord{Row: rowY, Col: colX})
			}
		}
	}

	return attachedGears
}
