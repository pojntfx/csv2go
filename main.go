package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/dave/jennifer/jen"
)

func main() {
	// Parse flags
	inFilePath := flag.String("in", "data.csv", "Path to the CSV input file")
	packageName := flag.String("package", "csv", "Package name of the generated output file")
	outFilePath := flag.String("out", "csv.go", "Path of the generated output file")

	flag.Parse()

	// Read CSV file into two-dimensional array
	csvFile, err := os.Open(*inFilePath)
	if err != nil {
		log.Fatal(err)
	}

	contents := [][]string{}

	csvReader := csv.NewReader(csvFile)
	currentLine := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		recordLenght := len(record)

		lineContents := []string{}

		for i := 0; i < recordLenght; i++ {
			lineContents = append(lineContents, record[i])
		}

		contents = append(contents, lineContents)

		currentLine++
	}

	// Generate Go code with two-dimensionsional array
	f := jen.NewFile(*packageName)

	f.Var().Defs(
		jen.Id("Content").Op("=").Index().Index().String().ValuesFunc(func(g *jen.Group) {
			for _, row := range contents {
				g.Index().String().Values(jen.Lit(row[0]), jen.Lit(row[1]))
			}
		}),
	)

	if err := f.Save(*outFilePath); err != nil {
		log.Fatal(err)
	}
}
