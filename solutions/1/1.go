package solution_1

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func partOne(inputLines []string) {
	sum := 0

	for _, l := range inputLines {
		r := []rune(l)

		d, _ := findDigit(string(r))
		slices.Reverse(r)
		dL, _ := findDigit(string(r))

		val, _ := strconv.Atoi(fmt.Sprintf("%c%c", d, dL))
		sum += val
	}

	fmt.Println("Sum is", sum)
}

func findDigit(s string) (rune, error) {
	for _, c := range s {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			return c, nil
		}
	}
	return rune(0), errors.New("no digit found")
}

var DIGIT_REPLACER = strings.NewReplacer(
	"one", "1",
	"two", "2",
	"three", "3",
	"four", "4",
	"five", "5",
	"six", "6",
	"seven", "7",
	"eight", "8",
	"nine", "9",
)

func partTwo(inputLines []string) {
	newLines := []string{}
	for _, s := range inputLines {
		newLines = append(newLines, DIGIT_REPLACER.Replace(s))
	}

	partOne(newLines)
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
