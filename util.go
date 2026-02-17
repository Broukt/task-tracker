package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func GetCurrentID() int {
	jsonData, err := os.ReadFile("tasks.json")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Error reading tasks.json: %v", err)
	} else if errors.Is(err, os.ErrNotExist) {
		return 1
	}
	var tasks []Task
	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		log.Fatalf("Error parsing tasks.json: %v", err)
	}
	lastID := tasks[len(tasks)-1].ID
	return lastID + 1
}

func SaveTask(task Task) {
	jsonData, err := os.ReadFile("tasks.json")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Error reading tasks.json: %v", err)
	} else if errors.Is(err, os.ErrNotExist) {
		jsonData = []byte("[]")
	}
	var tasks []Task
	if len(jsonData) > 0 {
		err = json.Unmarshal(jsonData, &tasks)
		if err != nil {
			log.Fatalf("Error parsing tasks.json: %v", err)
		}
	}
	tasks = append(tasks, task)
	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling tasks: %v", err)
	}
	err = os.WriteFile("tasks.json", updatedData, 0o644)
	if err != nil {
		log.Fatalf("Error writing to tasks.json: %v", err)
	}
}
