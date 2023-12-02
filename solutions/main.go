package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	
	solution_1 "jshe.dev/advent_template/solutions/1"
	
)

func main() {
	usage := "Usage: go run main.go [solution_to_run]"
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	nthSoln, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	functions := []func(inputLines []string){
			
			solution_1.Solution,
			
	}

	if nthSoln < 1 || nthSoln > len(functions) {
		fmt.Println(usage)
		os.Exit(1)
	}

	input := readInput(nthSoln)
	solnToRun := functions[nthSoln-1]

	solnToRun(input)
}

func readInput(n int) []string {
	inputFilePath := fmt.Sprintf("./solutions/%d/input.txt", n)

	file, err := os.Open(inputFilePath)

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Fatal(err)
		}

		lines = append(lines, scanner.Text())
	}

	return lines
}