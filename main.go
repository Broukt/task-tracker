package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

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
		fmt.Println("Usage: go run . add [description]")
	}
	id := GetCurrentID()
	description := options[0]
	status := "in-progress"
	createdAt := time.Now().Format(time.RFC1123)
	updatedAt := ""
	task := Task{
		id, description, status, createdAt, updatedAt,
	}
	SaveTask(task)
	fmt.Printf("Adding task with ID %d and description: %s\n", id, options[0])
}

func handleUpdate(options []string) {
	if options == nil || len(options) != 2 {
		fmt.Println("Invalid usage for update command")
		fmt.Println("Usage: go run . update [id] [new description]")
	}
}

func handleDelete(options []string) {
	if options == nil || len(options) != 1 {
		fmt.Println("Invalid usage for delete command")
		fmt.Println("Usage: go run . delete [id]")
	}
}

func handleMark(options []string) {
	if options == nil || len(options) != 2 {
		fmt.Println("Invalid usage for mark command")
		fmt.Println("Usage: go run . mark [id] [status]")
	}
}

func handleList(options []string) {
	if len(options) > 1 {
		fmt.Println("Invalid usage for list command")
		fmt.Println("Usage: go run . list [status]")
	}
}
