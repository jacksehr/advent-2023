package solution_7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func cmpHands(h1, h2 string, joker bool) int {
	h1Type, h2Type := identifyHand(h1), identifyHand(h2)

	var cmpV int
	if !joker || (!strings.ContainsRune(h1, 'J') && !strings.ContainsRune(h2, 'J')) {
		cmpV = int(h1Type - h2Type)
		if cmpV != 0 {
			return cmpV
		}
	} else {
		jh1, jh2 := replaceJokers(h1), replaceJokers(h2)
		jh1Type, jh2Type := identifyHand(jh1), identifyHand(jh2)

		cmpV = int(jh1Type - jh2Type)
		if cmpV != 0 {
			return cmpV
		}
	}

	for i := 0; i < len(h1); i++ {
		h1i, h2i := cardToInt(rune(h1[i]), joker), cardToInt(rune(h2[i]), joker)
		if h1i == h2i {
			continue
		}

		return h1i - h2i
	}

	return 0
}

func partOne(inputLines []string) {
	handsWithBids := [][]string{}

	for _, l := range inputLines {
		handsWithBids = append(handsWithBids, strings.Fields(l))
	}

	slices.SortStableFunc(
		handsWithBids,
		func(hb1, hb2 []string) int { return cmpHands(hb1[0], hb2[0], false) },
	)

	winningsSum := 0
	for i, hb := range handsWithBids {
		bid, _ := strconv.Atoi(hb[1])
		winningsSum += (i + 1) * bid
	}

	fmt.Println("Winnings sum is", winningsSum)
}

func partTwo(inputLines []string) {
	handsWithBids := [][]string{}

	for _, l := range inputLines {
		handsWithBids = append(handsWithBids, strings.Fields(l))
	}

	slices.SortStableFunc(
		handsWithBids,
		func(hb1, hb2 []string) int { return cmpHands(hb1[0], hb2[0], true) },
	)

	winningsSum := 0
	for i, hb := range handsWithBids {
		bid, _ := strconv.Atoi(hb[1])
		winningsSum += (i + 1) * bid
	}

	fmt.Println("Winnings sum is", winningsSum)
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
