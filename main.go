package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	getInput()
}

func getInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("task-cli ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}
