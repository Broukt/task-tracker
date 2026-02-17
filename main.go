package main

import (
	"fmt"
	"os"
	"strings"
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
	}
}

func handleAdd(options []string) {
	if options == nil || len(options) != 1 {
		fmt.Println("Invalid usage for add command")
		fmt.Println("Usage: go run main.go add [description]")
	}
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
