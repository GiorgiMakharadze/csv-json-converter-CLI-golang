package main

import (
	"fmt"
	"os"

	"github.com/GiorgiMakharadze/csv-json-converter-CLI-golang/pkg"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: csvtojson <conversion_direction> <input_file_path> <output_file_path>")
		return
	}

	conversionDirection := os.Args[1]
	inputFilePath := os.Args[2]
	outputFilePath := os.Args[3]

	var err error
	if conversionDirection == "csvtojson" {
		err = pkg.ConvertCSVToJSON(inputFilePath, outputFilePath)
	} else if conversionDirection == "jsontocsv" {
		err = pkg.ConvertJSONToCSV(inputFilePath, outputFilePath)
	} else {
		fmt.Println("Unknown conversion direction. Use 'csvtojson' or 'jsontocsv'.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Conversion successful!")
}
