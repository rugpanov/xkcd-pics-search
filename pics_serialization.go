package main

import (
	"encoding/json"
	"os"
)

func (comicsRef *Comics) SaveTo(fileName string) {
	bytes, err := json.Marshal(comicsRef)
	HandleError(err, "Cannot serialize comicsRef ")

	err = os.WriteFile(fileName, bytes, 0644)
	HandleError(err, "Cannot open file "+fileName+": ")
}

func Read(fileName string) (comics Comics) {
	reader, err := os.Open(fileName)
	HandleError(err, "Cannot open file "+fileName+": ")

	err = json.NewDecoder(reader).Decode(&comics)
	HandleError(err, "Cannot decode comics from file "+fileName+": ")
	return
}
