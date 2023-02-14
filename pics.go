package main

import "encoding/json"

type Comics []*Comic

type Comic struct {
	Title      string `json:"title"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Img        string `json:"img"`
	Transcript string `json:"transcript"`
}

func (pd Comic) String() string {
	marshal, err := json.Marshal(pd)
	HandleError(err, "Cannot marshal json: ")
	return string(marshal)
}
