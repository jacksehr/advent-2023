package solution_4

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"
)

var r = regexp.MustCompile(`Card *\d*: (.*) \| (.*)`)

type parsedCard struct {
	Want []string
	Have []string
}

func parseCard(card string) parsedCard {
	match := r.FindStringSubmatch(card)

	return parsedCard{
		Want: strings.Fields(match[1]),
		Have: strings.Fields(match[2]),
	}
}

func partOne(inputLines []string) {
	pointSum := 0

	for _, s := range inputLines {
		parsedCard := parseCard(s)

		winnersFound := 0
		for _, card := range parsedCard.Have {
			if slices.Contains(parsedCard.Want, card) {
				winnersFound++
			}
		}

		if winnersFound > 0 {
			pointSum += int(math.Pow(2, float64(winnersFound-1)))
		}
	}

	fmt.Println("Point sum is", pointSum)
}

type processedCard struct {
	Matches int
}

func partTwo(inputLines []string) {
	cards := []processedCard{}
	queue := []int{}

	for i, s := range inputLines {
		parsedCard := parseCard(s)

		winnersFound := 0
		for _, num := range parsedCard.Have {
			if slices.Contains(parsedCard.Want, num) {
				winnersFound++
			}
		}

		cards = append(cards, processedCard{Matches: winnersFound})
		queue = append(queue, i)
	}

	for i := 0; i < len(queue); i++ {
		index := queue[i]
		matches := cards[index].Matches

		for j := 1; j <= matches; j++ {
			if index+j >= len(cards) {
				break
			}

			queue = append(queue, index+j)
		}
	}

	fmt.Println("Card total is", len(queue))
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
