package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func getTasks() []Task {
	jsonData, err := os.ReadFile("tasks.json")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Error reading tasks.json: %v", err)
	} else if errors.Is(err, os.ErrNotExist) {
		return []Task{}
	}
	var tasks []Task
	if len(jsonData) > 0 {
		err = json.Unmarshal(jsonData, &tasks)
		if err != nil {
			log.Fatalf("Error parsing tasks.json: %v", err)
		}
	}
	return tasks
}

func getCurrentID() int {
	lastID := tasks[len(tasks)-1].ID
	return lastID + 1
}

func saveTasks() {
	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling tasks: %v", err)
	}
	err = os.WriteFile("tasks.json", updatedData, 0o644)
	if err != nil {
		log.Fatalf("Error writing to tasks.json: %v", err)
	}
}

func getTaskByID(id int) (*Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
