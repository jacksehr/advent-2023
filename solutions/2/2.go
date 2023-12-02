package solution_2

import (
	"fmt"
	"strconv"
	"strings"
)

func partOne(inputLines []string) {
	idSum := 0

	for i, l := range inputLines {
		_, game := parseGame(l)

		if isGamePossible(game, 12, 13, 14) {
			idSum += i + 1
		}
	}

	fmt.Println("Game sum is", idSum)
}

func partTwo(inputLines []string) {
	powerSum := 0

	for _, l := range inputLines {
		minNeededState := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		_, game := parseGame(l)
		state := getGameState(game)
		for color, n := range state {
			minNeededState[color] = max(minNeededState[color], n)
		}

		powerSum += (minNeededState["red"] * minNeededState["green"] * minNeededState["blue"])
	}

	fmt.Println("Power sum is", powerSum)
}

func getGameState(game [][]string) map[string]int {
	state := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, move := range game {
		for _, sm := range move {
			moveParts := strings.Split(sm, " ")
			n, _ := strconv.Atoi(moveParts[0])
			color := moveParts[1]

			state[color] = max(state[color], n)
		}
	}

	return state
}

func isGamePossible(game [][]string, red int, green int, blue int) bool {
	state := getGameState(game)

	for color, n := range state {
		switch color {
		case "red":
			if n > red {
				return false
			}
		case "green":
			if n > green {
				return false
			}
		case "blue":
			if n > blue {
				return false
			}
		}
	}

	return true
}

func parseGame(s string) (int, [][]string) {
	parts := strings.Split(s, ": ")

	gameId, _ := strconv.Atoi(strings.Replace(parts[0], "Game ", "", 1))

	moves := strings.Split(parts[1], "; ")

	splitMoves := [][]string{}
	for _, move := range moves {
		splitMoves = append(splitMoves, strings.Split(move, ", "))
	}

	return gameId, splitMoves
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
