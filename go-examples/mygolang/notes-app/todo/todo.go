package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct to represent a todo
// The todo struct contains fields for the title, content, and creation time of the todo.

type todo struct {
	Content   string    `json:"content"` //this is called as struct tags
	CreatedAt time.Time `json:"created_at"`
}

// Method to display the todo details
// This method prints the todo's title, content, and creation time in a formatted manner.

func (todo todo) Display() {
	fmt.Println("\n\n===================================")
	fmt.Println("todo Details:")
	fmt.Println("===================================")
	fmt.Printf("Content: %s\n\n", todo.Content)
	fmt.Printf("Created At: %s\n\n", todo.CreatedAt.Format(time.RFC1123))
	fmt.Println("===================================")
}

// Constructor function to create a new todo
// This function takes a title and content as input, validates them, and returns a new todo instance.
// If the title or content is empty, it returns an error.

func New(title, content string) (todo, error) {
	if title == "" || content == "" {
		return todo{}, errors.New("title and content cannot be empty")
	}

	return todo{
		Content:   content,
		CreatedAt: time.Now().Local().In(time.Now().Location()),
	}, nil
}

// Method to save the todo to a file
// This method takes a file name as input and writes the todo's details to the specified file.
// If the file already exists, it appends the todo to the file.
// If the file does not exist, it creates a new file and writes the todo to it.
// If there is an error during file operations, it returns an error.

func (todo todo) SaveToFile() error {
	fileName := todo.Content
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)

}
