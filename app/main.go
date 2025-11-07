package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		if scanner.Scan() {
			command := scanner.Text()
			handleCommand(command)
		}
	}
}

func handleCommand(c string) {
	command := strings.SplitN(c, " ", 2)
	cmd := command[0]

	switch cmd {
	case "exit":
		handleExitCommand(cmd, command[1])
	case "echo":
		handleEchoCommand(cmd, command[1])
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
	}

}

func handleExitCommand(_ string, args string) {
	exitCode, err := strconv.Atoi(args)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Cannot handle exit code: %v, error: %v", args, err)
		return
	}
	os.Exit(exitCode)
}

func handleEchoCommand(_ string, args string) {
	fmt.Fprintln(os.Stdout, args)
}
