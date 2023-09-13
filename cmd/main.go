package main

import (
	"fmt"
	"os"

	"github.com/GiorgiMakharadze/csv-json-converter-CLI-golang/pkg"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: csvtojson <csv_file_path> <json_file_path>")
		return
	}
	csvFilePath := os.Args[1]
	jsonFilePath := os.Args[2]

	err := pkg.ConvertCSVToJSON(csvFilePath, jsonFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Conversion successful!")
}
