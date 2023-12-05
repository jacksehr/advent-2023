package solution_5

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func partOne(inputLines []string) {
	input := strings.Split(strings.Join(inputLines, "\n"), "\n\n")

	s := strings.Fields(strings.Split(input[0], ": ")[1])
	seeds := []int{}
	for i := 0; i < len(s); i++ {
		seed, _ := strconv.Atoi(s[i])
		seeds = append(seeds, seed)
	}

	maps := []Map{}
	for _, mapInput := range input[1:] {
		splitMapInput := strings.Split(mapInput, "\n")

		newMap := mapFromInput(splitMapInput[1:])
		maps = append(maps, newMap)
	}

	lowestLoc := math.MaxInt
	for _, seed := range seeds {
		curr := seed
		for _, m := range maps {
			curr = m.Get(curr)
		}

		lowestLoc = min(curr, lowestLoc)
	}

	fmt.Println("Lowest loc is", lowestLoc)
}

func checkLoc(l int, maps []Map, seedRanges [][2]int, c chan int) {
	curr := l
	for _, m := range maps {
		curr = m.Get(curr)
	}

	for _, rng := range seedRanges {
		if curr >= rng[0] && curr <= rng[0]+rng[1] {
			c <- l
			return
		}
	}

	go checkLoc(l+1, maps, seedRanges, c)
}

func partTwo(inputLines []string) {
	input := strings.Split(strings.Join(inputLines, "\n"), "\n\n")

	matches := regexp.MustCompile(`(\d+ \d+)`).FindAllString(input[0], -1)
	seedRanges := [][2]int{}
	for _, match := range matches {
		rng := strings.Fields(match)
		start, _ := strconv.Atoi(rng[0])
		length, _ := strconv.Atoi(rng[1])
		seedRanges = append(seedRanges, [2]int{start, length})
	}

	maps := []Map{}
	for _, mapInput := range input[1:] {
		splitMapInput := strings.Split(mapInput, "\n")

		newMap := mapFromInput(splitMapInput[1:])
		newMap.Invert()
		maps = append(maps, newMap)
	}

	slices.Reverse(maps)

	c := make(chan int, 1)

	go checkLoc(0, maps, seedRanges, c)

	result := <-c
	fmt.Println("Lowest loc is", result)
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
