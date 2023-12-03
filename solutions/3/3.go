package solution_3

import (
	"fmt"
	"strconv"
)

func partOne(inputLines []string) {
	partNumSum := 0

	for y, row := range inputLines {
		for x := 0; x < len(row); x++ {
			if !isDigit(rune(row[x])) {
				continue
			}

			rowRemainder := row[x:]
			skipN := 0

			for i, c := range rowRemainder {
				if !isDigit(c) {
					skipN += i
					break
				}

				if x+i == len(row)-1 {
					skipN += i + 1
					break
				}
			}

			if hasAdjSymbol(inputLines, x, x+skipN, y) {
				partNum, _ := strconv.Atoi(row[x : x+skipN])
				partNumSum += partNum
			}

			x += skipN - 1
		}
	}

	fmt.Println("Sum is", partNumSum)
}

func partTwo(inputLines []string) {
	gearRatioSum := 0

	gearMap := map[coord][]numCoord{}

	for y, row := range inputLines {
		for x := 0; x < len(row); x++ {
			if !isDigit(rune(row[x])) {
				continue
			}

			rowRemainder := row[x:]
			skipN := 0

			for i, c := range rowRemainder {
				if !isDigit(c) {
					skipN += i
					break
				}

				if x+i == len(row)-1 {
					skipN += i + 1
					break
				}
			}

			adjGears := getAdjGears(inputLines, x, x+skipN, y)
			if len(adjGears) > 0 {
				for _, gear := range adjGears {
					newCoord := numCoord{Row: y, Col: numLoc{L: x, R: x + skipN}}

					_, ok := gearMap[gear]
					if !ok {
						gearMap[gear] = []numCoord{newCoord}
					} else {
						gearMap[gear] = append(gearMap[gear], newCoord)
					}
				}
			}

			x += skipN - 1
		}
	}

	for _, nums := range gearMap {
		if len(nums) != 2 {
			continue
		}

		num1, _ := strconv.Atoi(inputLines[nums[0].Row][nums[0].Col.L:nums[0].Col.R])
		num2, _ := strconv.Atoi(inputLines[nums[1].Row][nums[1].Col.L:nums[1].Col.R])

		gearRatioSum += num1 * num2
	}

	fmt.Println("Gear ratio sum is", gearRatioSum)
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
