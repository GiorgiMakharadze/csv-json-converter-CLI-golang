package pkg

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"testing"
)

func TestConvertCSVToJSON(t *testing.T) {
	csvFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(csvFile.Name())

	writer := csv.NewWriter(csvFile)
	writer.Write([]string{"Name", "Age"})
	writer.Write([]string{"Alice", "25"})
	writer.Write([]string{"Bob", "30"})
	writer.Flush()
	csvFile.Close()

	jsonFile, err := os.CreateTemp("", "test.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(jsonFile.Name())
	jsonFile.Close()

	err = ConvertCSVToJSON(csvFile.Name(), jsonFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open(jsonFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var data []map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		t.Fatal(err)
	}

	if len(data) != 2 {
		t.Fatalf("expected 2 rows but got %d", len(data))
	}

	if data[0]["Name"] != "Alice" {
		t.Fatalf("expected Alice but got %s", data[0]["Name"])
	}
}

func TestConvertJSONToCSV(t *testing.T) {
	jsonFile, err := os.CreateTemp("", "test.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(jsonFile.Name())

	jsonData := []map[string]interface{}{
		{"Name": "Alice", "Age": 25},
		{"Name": "Bob", "Age": 30},
	}

	encoder := json.NewEncoder(jsonFile)
	if err := encoder.Encode(&jsonData); err != nil {
		t.Fatal(err)
	}
	jsonFile.Close()

	csvFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(csvFile.Name())
	csvFile.Close()

	err = ConvertJSONToCSV(jsonFile.Name(), csvFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open(csvFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	header := map[string]int{}
	for idx, name := range records[0] {
		header[name] = idx
	}

	if _, ok := header["Name"]; !ok {
		t.Fatal("Name header not found in CSV file")
	}

	if _, ok := header["Age"]; !ok {
		t.Fatal("Age header not found in CSV file")
	}

	expectedData := []map[string]string{
		{"Name": "Alice", "Age": "25"},
		{"Name": "Bob", "Age": "30"},
	}

	for i, expected := range expectedData {
		for key, value := range expected {
			if records[i+1][header[key]] != value {
				t.Fatalf("unexpected data in CSV file: got %+v, want %+v", records[i+1][header[key]], value)
			}
		}
	}
}