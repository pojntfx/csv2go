# csv-to-go-generator

Generate Go code containing a multi-dimensional array with the contents of a CSV file.

## Overview

`csv-to-go-generator` is a Go code generator that takes a CSV file as it's input and creates a multidimensional array is it's output.

## Installation

A Go package [is available](https://pkg.go.dev/mod/github.com/pojntfx/csv-to-go-generator).

## Usage

```bash
% csv-to-go-generator -help
Usage of csv-to-go-generator:
  -in string
        Path to the CSV input file (default "data.csv")
  -out string
        Path of the generated output file (default "csv.go")
  -package string
        Package name of the generated output file (default "csv")
```

## License

csv-to-go-generator (c) 2020 Felicitas Pojtinger

SPDX-License-Identifier: AGPL-3.0
