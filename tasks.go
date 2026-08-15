package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
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

func (tl *TaskList) showTasks(status Status) error {
	if (len(tl.Tasks) == 0) {
		fmt.Println("No tasks was found.")
		return nil
	}
	// Sort by UpdatedAt
	sort.Slice(tl.Tasks, func(i, j int) bool {
		return tl.Tasks[i].UpdatedAt.After(tl.Tasks[j].UpdatedAt)
	})

	found := false

	fmt.Println("\n========================================================")
	for _, t := range tl.Tasks {
		if status != "" && t.Status != status {
			continue
		}

		found = true

		fmt.Printf(
			"ID: %d | %s | %s | updated: %s\n",
			t.ID,
			t.Status,
			t.Description,
			t.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}
	if !found {
		fmt.Printf("no %s tasks was found...\n", status)
	}
	fmt.Println("========================================================")

	return nil
}
