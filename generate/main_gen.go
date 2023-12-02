package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func assertErrNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmpl, err := template.New("main").Parse(MAIN_TEMPLATE)
	assertErrNil(err)

	solutionsFound, err := os.ReadDir("./solutions")
	assertErrNil(err)

	templateData := struct {
		Dirs []string
	}{
		Dirs: []string{},
	}

	for _, soln := range solutionsFound {
		if !soln.IsDir() {
			continue
		}

		number, _ := strings.CutSuffix(soln.Name(), ".go")
		templateData.Dirs = append(templateData.Dirs, number)
	}

	mainFile, err := os.Create("./solutions/main.go")
	assertErrNil(err)

	tmpl.Execute(mainFile, templateData)
	assertErrNil(err)
}
