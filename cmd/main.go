package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Fprint

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)
		switch {
		case command == "exit 0":
			os.Exit(0)
		case strings.HasPrefix(command, "echo "):
			fmt.Println(strings.TrimPrefix(command, "echo "))
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
