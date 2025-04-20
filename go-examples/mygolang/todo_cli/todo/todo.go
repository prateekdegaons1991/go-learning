package todo

// create a struct to represent a todo item
type Todo struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

// create a slice to hold the todo items

var todos []Todo

// create a function to add a todo item
func AddTodo(title string, description string) {
	todo := Todo{
		ID:          len(todos) + 1,
		Title:       title,
		Description: description,
		Done:        false,
	}
	todos = append(todos, todo)
}

// create a function to list all todo items
func ListTodos() {
	for _, todo := range todos {
		done := " "
		if todo.Done {
			done = "x"
		}
		println(todo.ID, done, todo.Title)
	}
}

// create a function to mark a todo item as done
func MarkTodoDone(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			break
		}
	}
}

// create a function to delete a todo item
func DeleteTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
}

// create a function to get a todo item by id
func GetTodoByID(id int) *Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo
		}
	}
	return nil
}
