package main

import "fmt"

var currentId int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Create a MySQL db for golang"})
	RepoCreateTodo(Todo{Name: "Learn how to test in golang")
	RepoCreateTodo(Todo{Name: "Refactor this file")
	RepoCreateTodo(Todo{Name: "Make a package to save functions in")
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
