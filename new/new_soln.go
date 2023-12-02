package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const AOC_SESSION_KEY = "AOC_SESSION"

var year = time.Now().Year()

func main() {
	flag.IntVar(&year, "year", year, "-year [year] for a custom year")
	flag.Parse()

	solns, err := os.ReadDir("./solutions")
	if err != nil {
		log.Fatal(err)
	}
	solns = slices.DeleteFunc(solns, func(soln os.DirEntry) bool {
		return !soln.IsDir()
	})

	numNewSoln := 1
	if len(solns) > 0 {
		lastSoln := solns[len(solns)-1]
		lastSolnName, _ := strings.CutSuffix(lastSoln.Name(), ".go")
		numLastSoln, err := strconv.Atoi(lastSolnName)
		if err != nil {
			log.Fatal(err)
		}

		numNewSoln = numLastSoln + 1
	}

	newDirPath := fmt.Sprintf("solutions/%d", numNewSoln)
	err = os.Mkdir(newDirPath, 0770)
	if err != nil {
		log.Fatal(err)
	}

	newSrcFile, err := os.Create(path.Join(newDirPath, fmt.Sprintf("%d.go", numNewSoln)))
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("solution").Parse(string(SOLUTION_TEMPLATE))
	if err != nil {
		log.Fatal(err)
	}

	templateData := struct{ N int }{N: numNewSoln}
	tmpl.Execute(newSrcFile, templateData)

	newInputFile, err := os.Create(path.Join(newDirPath, "input.txt"))
	if err != nil {
		log.Fatalf("input.txt could not be generated: %v", err)
	}

	newInput, err := getInput(year, numNewSoln)
	if err != nil {
		log.Fatalf("input could not be retrieved: %v", err)
	}

	if _, err := newInputFile.WriteString(newInput); err != nil {
		log.Fatalf("input could not be written to file: %v", err)
	}
}

func getInput(year int, day int) (string, error) {
	aocSession, ok := os.LookupEnv(AOC_SESSION_KEY)
	if !ok {
		return "", fmt.Errorf("%s not found in env", AOC_SESSION_KEY)
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return "", fmt.Errorf("failed to make req for input: %w", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: aocSession})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("input request failed: %w", err)
	}

	input, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strings.TrimSpace(string(input)), nil
}
