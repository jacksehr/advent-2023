package solution_5

import (
	"slices"
	"strconv"
	"strings"
)

type mapRange struct {
	src    int
	dest   int
	length int
}

type Map struct {
	ranges []mapRange
}

func mapFromInput(s []string) Map {
	newMap := Map{[]mapRange{}}

	for _, l := range s {
		parts := strings.Fields(l)
		destStart, _ := strconv.Atoi(parts[0])
		srcStart, _ := strconv.Atoi(parts[1])
		ln, _ := strconv.Atoi(parts[2])

		newMap.ranges = append(newMap.ranges, mapRange{src: srcStart, dest: destStart, length: ln})
	}

	slices.SortFunc(newMap.ranges, func(a, b mapRange) int {
		return a.src - b.src
	})

	return newMap
}

func (m *Map) Invert() {
	for i := 0; i < len(m.ranges); i++ {
		m.ranges[i].src, m.ranges[i].dest = m.ranges[i].dest, m.ranges[i].src
	}
}

func (m *Map) Get(k int) int {
	v := k

	for _, rng := range m.ranges {
		if k >= rng.src && k <= rng.src+rng.length {
			v = rng.dest + (k - rng.src)
			break
		}
	}

	return v
}
