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

	go func ()  {
		<-sigChan
		fmt.Println("\ntask-cli is abandoned, goodbye...")
		os.Exit(0)
	}()

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

		fmt.Println("command: ", command)
		fmt.Println("description: ", description)
	}
}


