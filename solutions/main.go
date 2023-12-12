package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	
	solution_1 "jshe.dev/advent_template/solutions/1"
	
	solution_2 "jshe.dev/advent_template/solutions/2"
	
	solution_3 "jshe.dev/advent_template/solutions/3"
	
	solution_4 "jshe.dev/advent_template/solutions/4"
	
	solution_5 "jshe.dev/advent_template/solutions/5"
	
	solution_6 "jshe.dev/advent_template/solutions/6"
	
	solution_7 "jshe.dev/advent_template/solutions/7"
	
	solution_8 "jshe.dev/advent_template/solutions/8"
	
	solution_9 "jshe.dev/advent_template/solutions/9"
	
	solution_10 "jshe.dev/advent_template/solutions/10"
	
	solution_11 "jshe.dev/advent_template/solutions/11"
	
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
			
			solution_2.Solution,
			
			solution_3.Solution,
			
			solution_4.Solution,
			
			solution_5.Solution,
			
			solution_6.Solution,
			
			solution_7.Solution,
			
			solution_8.Solution,
			
			solution_9.Solution,
			
			solution_10.Solution,
			
			solution_11.Solution,
			
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