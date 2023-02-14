package main

import (
	"encoding/json"
	"os"
)

func (pics *Pics) SaveTo(fileName string) {
	bytes, err := json.Marshal(pics)
	HandleError(err, "Cannot serialize pics ")

	err = os.WriteFile(fileName, bytes, 0644)
	HandleError(err, "Cannot open file "+fileName+": ")
}

func Read(fileName string) (pics Pics) {
	reader, err := os.Open(fileName)
	HandleError(err, "Cannot open file "+fileName+": ")

	err = json.NewDecoder(reader).Decode(&pics)
	HandleError(err, "Cannot decode pics from file "+fileName+": ")
	return
}
