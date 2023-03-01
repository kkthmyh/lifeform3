package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(path string) []string {

	var addressList []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	for _, line := range csvData {
		addressList = append(addressList, line[0])
	}
	return addressList
}
