package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/prateekdegaons1991/todo_cli/todo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [<args>]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo add <title> <description>")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		todo.AddTodo(title, description)
	case "list":
		todo.ListTodos()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <id>")
			return
		}
		id := os.Args[2]
		intID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid ID. Please provide a numeric ID.")
			return
		}
		todo.MarkTodoDone(intID)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			return
		}
		id := os.Args[2]
		intID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid ID. Please provide a numeric ID.")
			return
		}
		todo.DeleteTodo(intID)
	default:
		fmt.Println("Unknown command:", command)
	}
}

// how do i run this code
