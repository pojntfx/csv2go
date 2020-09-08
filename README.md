# csv2go

Generates Go code containing a multi-dimensional array with the contents of a CSV file.

## Overview

`csv2go` is a Go code generator that takes a CSV file as it's input and creates a multidimensional array is it's output.

## Installation

### Prebuilt Binaries

Linux, macOS and Windows binaries are available on [GitHub Releases](https://github.com/pojntfx/csv2go/releases).

### Go Package

A Go package [is available](https://pkg.go.dev/mod/github.com/pojntfx/csv2go).

## Usage

```bash
% csv2go -help
Usage of csv2go:
  -in string
        Path to the CSV input file (default "data.csv")
  -out string
        Path of the generated output file (default "csv.go")
  -package string
        Package name of the generated output file (default "csv")
```

## License

csv2go (c) 2020 Felicitas Pojtinger

SPDX-License-Identifier: AGPL-3.0
