package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func getCurrentID() int {
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

func saveTask(task Task) {
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

func main() {
	args := os.Args
	if len(args) > 1 {
		cli(args[1:])
		return
	}
	fmt.Println("Usage: go run main.go [command] [options]")
}

func cli(stringArgs []string) {
	command := strings.ToLower(stringArgs[0])
	options := stringArgs[1:]
	switch command {
	case "add":
		handleAdd(options)
	case "update":
		handleUpdate(options)
	case "delete":
		handleDelete(options)
	case "mark":
		handleMark(options)
	case "list":
		handleList(options)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func handleAdd(options []string) {
	if options == nil || len(options) != 1 {
		fmt.Println("Invalid usage for add command")
		fmt.Println("Usage: go run main.go add [description]")
	}
	id := getCurrentID()
	description := options[0]
	status := "in-progress"
	createdAt := time.Now().Format(time.RFC1123)
	updatedAt := ""
	task := Task{
		id, description, status, createdAt, updatedAt,
	}
	saveTask(task)
	fmt.Printf("Adding task with ID %d and description: %s\n", id, options[0])
}

func handleUpdate(options []string) {
	if options == nil || len(options) != 2 {
		fmt.Println("Invalid usage for update command")
		fmt.Println("Usage: go run main.go update [id] [new description]")
	}
}

func handleDelete(options []string) {
	if options == nil || len(options) != 1 {
		fmt.Println("Invalid usage for delete command")
		fmt.Println("Usage: go run main.go delete [id]")
	}
}

func handleMark(options []string) {
	if options == nil || len(options) != 2 {
		fmt.Println("Invalid usage for mark command")
		fmt.Println("Usage: go run main.go mark [id] [status]")
	}
}

func handleList(options []string) {
	if len(options) > 1 {
		fmt.Println("Invalid usage for list command")
		fmt.Println("Usage: go run main.go list [status]")
	}
}
