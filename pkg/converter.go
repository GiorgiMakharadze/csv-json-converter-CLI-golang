package pkg

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func ConvertCSVToJSON(csvFilePath string, jsonFilePath string) error {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var jsonData []map[string]string

	headers := csvData[0]

	for _, row := range csvData[1:] {
		entry := make(map[string]string)
		for i, header := range headers {
			entry[header] = row[i]
		}
		jsonData = append(jsonData, entry)
	}

	jsonFile, err := os.Create(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonEncoder := json.NewEncoder(jsonFile)
	err = jsonEncoder.Encode(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func ConvertJSONToCSV(jsonFilePath string, csvFilePath string) error {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	var jsonData []map[string]interface{}

	jsonDecoder := json.NewDecoder(jsonFile)
	err = jsonDecoder.Decode(&jsonData)
	if err != nil {
		log.Fatal(err)
	}

	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	headers := make([]string, 0)
	for header := range jsonData[0] {
		headers = append(headers, header)
	}
	sort.Strings(headers)
	err = writer.Write(headers)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range jsonData {
		var row []string
		for _, header := range headers {
			value := entry[header]
			strValue := fmt.Sprintf("%v", value)
			row = append(row, strValue)
		}
		err = writer.Write(row)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
