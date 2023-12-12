package solution_12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

const TEST_INPUT = `
???.### 1,1,3
.??..??...?##. 1,1,3
`

// .??..??...?##. 1,1,3
// ?#?#?#?#?#?#?#? 1,3,1,6
// ????.#...#... 4,1,1
// ????.######..#####. 1,6,5
// ?###???????? 3,2,1

func getBinaryPerms(line string, hashLimit int) []string {
	for i, c := range line {
		if c == '?' {
			head := line[:i]

			dotTails := getBinaryPerms(line[i+1:], hashLimit)
			hashTails := getBinaryPerms(line[i+1:], hashLimit-1)

			perms := []string{}
			for _, t := range dotTails {
				perms = append(perms,
					strings.Join([]string{head, ".", t}, ""),
				)
			}

			for _, t := range hashTails {
				if hashLimit > 0 {
					perms = append(perms,
						strings.Join([]string{head, "#", t}, ""),
					)
				}
			}

			return perms
		}
	}

	return []string{line}
}

func partOne(inputLines []string) {
	// inputLines = strings.Split(strings.TrimSpace(TEST_INPUT), "\n")

	eg := errgroup.Group{}

	validSum := 0

	c := make(chan int)

	go func() {
		for range c {
			validSum++
		}
	}()

	for _, l := range inputLines {
		parts := strings.Fields(l)

		line := parts[0]
		blockLengths := []int{}
		blSum := 0
		for _, n := range strings.Split(parts[1], ",") {
			bl, _ := strconv.Atoi(n)
			blSum += bl
			blockLengths = append(blockLengths, bl)
		}

		calc := func() error {
			perms := getBinaryPerms(line, blSum-strings.Count(line, "#"))

			patternParts := []string{}
			for _, blockLen := range blockLengths {
				patternParts = append(patternParts, fmt.Sprintf(`#{%d}`, blockLen))
			}
			pattern := regexp.MustCompile(strings.Join(patternParts, `\.+`))

			for _, p := range perms {
				if pattern.MatchString(p) {
					c <- 1
				}
			}

			return nil
		}

		eg.Go(calc)
	}

	eg.Wait()

	fmt.Println("Arrangement count sum was", validSum)
}

func partTwo(inputLines []string) {
	inputLines = strings.Split(strings.TrimSpace(TEST_INPUT), "\n")

	eg := errgroup.Group{}

	validSum := 0

	c := make(chan int)

	go func() {
		for range c {
			validSum++
		}
	}()

	for _, l := range inputLines {
		parts := strings.Fields(l)

		line := parts[0]
		unfoldedLine, unfoldedBlocks := []string{}, []string{}
		for i := 0; i < 5; i++ {
			unfoldedLine = append(unfoldedLine, line)
			unfoldedBlocks = append(unfoldedBlocks, parts[1])
		}
		line = strings.Join(unfoldedLine, "?")
		blocks := strings.Join(unfoldedBlocks, ",")

		blockLengths := []int{}
		blSum := 0
		for _, n := range strings.Split(blocks, ",") {
			bl, _ := strconv.Atoi(n)
			blSum += bl
			blockLengths = append(blockLengths, bl)
		}

		calc := func() error {
			perms := getBinaryPerms(line, blSum-strings.Count(line, "#"))

			patternParts := []string{}
			for _, blockLen := range blockLengths {
				patternParts = append(patternParts, fmt.Sprintf(`#{%d}`, blockLen))
			}
			pattern := regexp.MustCompile(strings.Join(patternParts, `\.+`))

			for _, p := range perms {
				if pattern.MatchString(p) {
					c <- 1
				}
			}

			return nil
		}

		eg.Go(calc)
	}

	eg.Wait()

	fmt.Println("Arrangement count sum was", validSum)
}

func Solution(inputLines []string) {
	// fmt.Println("Part one")
	// partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
