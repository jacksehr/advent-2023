package solution_9

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func buildSeqStack(seq []int) [][]int {
	seqStack := [][]int{seq}
	seqIdx := 0

	for i := 0; i < len(seqStack); i++ {
		curr := seqStack[seqIdx]
		newSeq := []int{}

		allZeroes := true
		for j := 0; j < len(curr)-1; j++ {
			diff := curr[j+1] - curr[j]
			if allZeroes && diff != 0 {
				allZeroes = false
			}
			newSeq = append(newSeq, diff)
		}

		if !allZeroes {
			seqStack = append(seqStack, newSeq)
			seqIdx++
		}
	}

	return seqStack
}

func getNextValue(seq []int, front bool) int {
	seqStack := buildSeqStack(seq)
	slices.Reverse(seqStack)

	newValue := 0
	if !front {
		for _, sq := range seqStack {
			newValue += sq[len(sq)-1]
		}
	} else {
		for _, sq := range seqStack {
			newValue = sq[0] - newValue
		}
	}

	return newValue
}

func partOne(inputLines []string) {
	valueSum := 0

	for _, l := range inputLines {
		seq := []int{}
		for _, s := range strings.Fields(l) {
			num, _ := strconv.Atoi(s)
			seq = append(seq, num)
		}

		valueSum += getNextValue(seq, false)
	}

	fmt.Println("Sum of extrapolated values is", valueSum)
}

func partTwo(inputLines []string) {
	valueSum := 0

	for _, l := range inputLines {
		seq := []int{}
		for _, s := range strings.Fields(l) {
			num, _ := strconv.Atoi(s)
			seq = append(seq, num)
		}

		valueSum += getNextValue(seq, true)
	}

	fmt.Println("Sum of extrapolated values is", valueSum)
}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}
