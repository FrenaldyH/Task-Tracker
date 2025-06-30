package main

import (
	"encoding/json"
	"errors"
	"os"
)

type taskS struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

var arrT []taskS

const DB = "Task.json"

func loadTask() error {
	if _, err := os.Stat(DB); err != nil {
		os.WriteFile(DB, []byte("[]"), 0644)
	} else if readData, err := os.ReadFile(DB); err != nil {
		return errors.New("error reading the data base file")
	} else if err := json.Unmarshal([]byte(readData), &arrT); err != nil {
		return errors.New("error unmarshaling JSON")
	}
	return nil
}

func saveTask() error {
	if newData, err := json.MarshalIndent(arrT, "", "  "); err != nil {
		return errors.New("error marshaling JSON")
	} else if err := os.WriteFile(DB, newData, 0644); err != nil {
		return errors.New("error writing the data base file")
	}
	return nil
}
