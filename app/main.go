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

func handleCommand(command string) {
	if strings.HasPrefix(command, "exit") {
		handleExitCommand(command)
	}
	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
}

func handleExitCommand(command string) {
	cmd := strings.Split(command, " ")
	arg := cmd[1]

	exitCode, err := strconv.Atoi(arg)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Cannot handle exit code: %v, error: %v", arg, err)
		return
	}
	os.Exit(exitCode)

}
