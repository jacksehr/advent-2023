package solution_15

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const TEST_INPUT = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func hash(s string) int {
	curr := 0
	for _, c := range s {
		curr += int(c)
		curr *= 17
		curr = curr % 256
	}

	return curr
}

func partOne(inputLines []string) {
	// inputLines = strings.Split(TEST_INPUT, "\n")

	l := strings.Split(inputLines[0], ",")

	hashSum := 0
	for _, s := range l {
		hashSum += hash(s)
	}

	fmt.Println("Hash sum is", hashSum)
}

type lens struct {
	label       string
	focalLength int
}

func partTwo(inputLines []string) {
	// inputLines = strings.Split(TEST_INPUT, "\n")

	l := strings.Split(inputLines[0], ",")
	m := map[int][]lens{}

	for _, s := range l {
		parts := strings.Split(s, "=")
		switch len(parts) {
		case 1:
			label, _ := strings.CutSuffix(s, "-")
			if sm, ok := m[hash(label)]; ok {
				m[hash(label)] = slices.DeleteFunc(sm, func(l lens) bool { return l.label == label })
			}
		case 2:
			label := parts[0]
			n, _ := strconv.Atoi(parts[1])

			box, ok := m[hash(label)]
			if !ok {
				m[hash(label)] = []lens{{label, n}}
			} else {
				added := false
				for i, light := range box {
					if light.label == label {
						m[hash(label)] = slices.Replace(box, i, i+1, lens{label, n})
						added = true
						break
					}
				}
				if !added {
					m[hash(label)] = append(box, lens{label, n})
				}
			}
		}
	}

	fpSum := 0
	for k, v := range m {
		if len(v) == 0 {
			continue
		}

		for i, light := range v {
			fpSum += (k + 1) * (i + 1) * light.focalLength
		}
	}

	fmt.Println("Focusing power sum is", fpSum)
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
