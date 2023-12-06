package solution_6

import (
	"fmt"
	"strconv"
	"strings"
)

func getWins(t int, d int) int {
	l, r := make(chan int, 1), make(chan int, 1)

	go func() {
		for timeHeld := 0; timeHeld < t; timeHeld++ {
			travelled := timeHeld * (t - timeHeld)
			if travelled > d {
				l <- timeHeld
				return
			}
		}
	}()

	go func() {
		for timeHeld := t; timeHeld >= 0; timeHeld-- {
			travelled := timeHeld * (t - timeHeld)
			if travelled > d {
				r <- timeHeld
				return
			}
		}
	}()

	return <-r - <-l + 1
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
