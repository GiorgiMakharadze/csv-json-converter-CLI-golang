# csv-json-converter-CLI-golang

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This repository houses a simple CLI (Command Line Interface) application written in Go. The tool allows for the conversion of CSV files to JSON format and vice versa.

## Prerequisites

- Go 1.16 or higher. [Download here](https://golang.org/dl/)

## Getting Started

Clone the repository to your local machine using the following command:

```sh
$ git clone https://github.com/GiorgiMakharadze/csv-json-converter-CLI-golang.git

$ cd csv-json-converter-CLI-golang
```

## Building the App

```sh
$ make build
```

## Usage

The CLI tool takes three arguments:

1. Conversion direction (either "csvtojson" or "jsontocsv").
2. Input file path.
3. Output file path.

Use the tool as follows

```sh
$ ./bin/csvjsonconverter csvtojson input.csv output.json
```

or

```sh
$ ./bin/csvjsonconverter jsontocsv input.json output.csv
```

If the conversion is successful, it will output: `Conversion successful!`

## Test the App

```sh
$ make test
```

This will execute the test cases and output the results.

## Cleaning Up

You can clean up generated files using commands from the Makefile:

1. To remove a specific file:

```sh
$ make clean file=filename.csv/json
```

2. To remove all CSV and JSON files:

```sh
$ make clean-all
```

3. To delete all files with a specified extension:

```sh
$ make delete type=csv/json
```
