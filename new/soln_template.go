package main

const SOLUTION_TEMPLATE = `package solution_{{.N}}

import (
	"fmt"
)

func partOne(inputLines []string) {
	fmt.Println(inputLines)
}

func partTwo(inputLines []string) {}

func Solution(inputLines []string) {
	fmt.Println("Part one")
	partOne(inputLines)

	fmt.Println("Part two")
	partTwo(inputLines)
}`
