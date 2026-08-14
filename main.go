package main

import (
	"fmt"
	"os"
)

func main() {
	sigs := make(chan os.Signal, 1)
	msg := make(chan string, 1)
	go func () {
		// Receive input in a loop
		for {
			var s string
			fmt.Print("task-cli ")
			fmt.Scan(&s)
			msg <- s
		}
	}()

	loop:
	for {
		select {
		case <-sigs:
		fmt.Println("Program shutdown, byebye...")
		break loop

		case s := <-msg:
			switch {
			case s == "add":
				fmt.Println("Task added successfully (ID: )", msg, s)
			}	

		}
	}
	
}

