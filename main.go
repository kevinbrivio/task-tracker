package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	// read the line from runtime
	reader := bufio.NewReader(os.Stdin)

	// Listen for every time control+c was typed
	sigChan := make(chan os.Signal, 1)

	// If the OS sends SIGINT (ctrl+C) or SIGTERM
	// We might need to drop it into sigChan instead of killing the program
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\ntask-cli is abandoned, goodbye...")
		os.Exit(0)
	}()

	// Init the Task
	taskList := TaskList{}

	for {
		fmt.Print("task-cli ")

		line, _ := reader.ReadString('\n') // Read until we got newline a.k.a enter
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if line == "exit" {
			break
		}
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		description := ""

		if len(parts) > 1 {
			description = parts[1]
		}

		switch command {
		case "add":
			if description == "" {
				fmt.Println("Oh no, what do you want to add...?")
				fmt.Println("Use this format: add [task]")
				continue
			}

			id, err := taskList.addTask(description)
			if err != nil {
				break
			}
			fmt.Printf("Task added successfully (ID: %d)\n", id)

		case "list":
			// empty description means list all tasks
			if description == "" {
				fmt.Println("============================")
				for _, t := range taskList.Tasks {
					fmt.Printf("[%d]. %s\n", t.ID, t.Description)
				}
				fmt.Println("============================")
			}
		}
	}
}
