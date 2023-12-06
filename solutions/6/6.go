package solution_6

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getWins(t, d int) int {
	tF, dF := float64(t), float64(d)

	roots := [2]float64{
		// quadratic formula to solve for x in: x(t - x) > d
		(tF - math.Sqrt(math.Pow(tF, 2)-4*dF)) / 2,
		(tF + math.Sqrt(math.Pow(tF, 2)-4*dF)) / 2,
	}

	return int(math.Floor(roots[1])-math.Ceil(roots[0])) + 1
}

func partOne(inputLines []string) {
	times := strings.Fields(inputLines[0])[1:]
	distances := strings.Fields(inputLines[1])[1:]

	winProduct := 1

	for i := 0; i < len(distances); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])

		winProduct *= getWins(t, d)
	}

	fmt.Println("Win product is", winProduct)
}

func partTwo(inputLines []string) {
	t, _ := strconv.Atoi(strings.Join(strings.Fields(inputLines[0])[1:], ""))
	d, _ := strconv.Atoi(strings.Join(strings.Fields(inputLines[1])[1:], ""))

	fmt.Println("There are", getWins(t, d), "wins")
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
