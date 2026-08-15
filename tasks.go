package main

import (
	"fmt"
	"strconv"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) addTask(description string) (int, error) {
	num := fmt.Sprintf("%d", len(tl.Tasks)+1)
	id, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Cannot convert ID: ", err)
		return -1, err
	}

	newTask := Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tl.Tasks = append(tl.Tasks, newTask)

	return id, nil
}
