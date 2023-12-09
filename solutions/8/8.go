package solution_8

import (
	"fmt"
	"strings"
)

func partOne(inputLines []string) {
	directions := strings.Split(inputLines[0], "")
	tree := map[string][2]string{}

	replacer := strings.NewReplacer("=", "", "(", "", ",", "", ")", "")
	nodes := inputLines[2:]
	for _, l := range nodes {
		nodeParts := strings.Fields(replacer.Replace(l))
		tree[nodeParts[0]] = [2]string{nodeParts[1], nodeParts[2]}
	}

	nSteps := getNStepsToZ(directions, tree)
	fmt.Println(nSteps("AAA"), "steps were required")
}

func partTwo(inputLines []string) {
	directions := strings.Split(inputLines[0], "")
	tree := map[string][2]string{}

	replacer := strings.NewReplacer("=", "", "(", "", ",", "", ")", "")
	nodes := inputLines[2:]

	endingInA := []string{}
	for _, l := range nodes {
		nodeParts := strings.Fields(replacer.Replace(l))
		tree[nodeParts[0]] = [2]string{nodeParts[1], nodeParts[2]}

		if strings.HasSuffix(nodeParts[0], "A") {
			endingInA = append(endingInA, nodeParts[0])
		}
	}

	nSteps := getNStepsToZ(directions, tree)
	steps := []int{}
	for _, n := range endingInA {
		steps = append(steps, nSteps(n))
	}

	lcm := LCM(steps[0], steps[1], steps[2:]...)

	fmt.Println(lcm, "steps required")
}

func getNStepsToZ(directions []string, tree map[string][2]string) func(start string) int {
	return func(start string) int {
		curr := start
		i := 0
		for ; !strings.HasSuffix(curr, "Z"); i++ {
			switch directions[i%len(directions)] {
			case "L":
				curr = tree[curr][0]
			case "R":
				curr = tree[curr][1]
			}
		}

		return i
	}
}

// some of Google's finest below

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
