package solution_7

import (
	"strconv"
	"strings"
)

func parseHand(h string) map[rune]int {
	handMap := map[rune]int{}

	for _, c := range h {
		_, ok := handMap[c]
		if !ok {
			handMap[c] = 1
		} else {
			handMap[c]++
		}
	}

	return handMap
}

type Hand int

const ( // iota is reset to 0
	HIGH Hand = iota
	ONE_PAIR
	TWO_PAIR
	THREE
	FULL_HOUSE
	FOUR
	FIVE
)

func identifyHand(h string) Hand {
	hMap := parseHand(h)

	numDistinctCards := len(hMap)
	switch numDistinctCards {
	case 1:
		return FIVE
	case 2:
		for _, nCards := range hMap {
			if nCards == 4 {
				return FOUR
			}
		}
		return FULL_HOUSE
	case 3:
		for _, nCards := range hMap {
			if nCards == 3 {
				return THREE
			}
		}
		return TWO_PAIR
	case 4:
		return ONE_PAIR
	default:
		return HIGH
	}
}

var cardToIntMap = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
}

func cardToInt(c rune, joker bool) int {
	if joker && c == 'J' {
		return 1
	}

	if i, ok := cardToIntMap[c]; ok {
		return i
	}

	i, _ := strconv.Atoi(string(c))
	return i
}

var jokerCache = map[string]string{}

func replaceJokers(h string) string {
	if !strings.ContainsRune(h, 'J') {
		return h
	}

	if cachedJoker, ok := jokerCache[h]; ok {
		return cachedJoker
	}

	hMap := parseHand(h)

	replacementType := identifyHand(h)
	replacementH := h

	for c := range hMap {
		if c == 'J' {
			continue
		}

		newHand := strings.ReplaceAll(h, "J", string(c))
		nhType := identifyHand(newHand)
		if int(nhType-replacementType) > 0 {
			replacementH = newHand
			replacementType = nhType
		}
	}

	jokerCache[h] = replacementH

	return replacementH
}
