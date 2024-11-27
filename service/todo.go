package service

import (
	"fmt"
	"time"
	"todo-cc/port"
)

type Todo struct {
	persistence port.PersistencePort
}

func NewTodo(todoPersistence port.PersistencePort) Todo {
	return Todo{
		todoPersistence,
	}
}

func (t Todo) CreateNewTask(title, description string, deadline time.Time, completed bool) error {
	err := t.persistence.NewTask(title, description, deadline, completed)
	if err != nil {
		return fmt.Errorf("error while saving task: %v+", err.Error())
	}

	return nil
}
