package pkg

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
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
