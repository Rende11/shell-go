package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

var builtInCommands = []string{"echo", "exit", "type"}

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
	case "type":
		handleTypeCommand(cmd, command[1])
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
	}

}

func handleExitCommand(_, args string) {
	exitCode, err := strconv.Atoi(args)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Cannot handle exit code: %v, error: %v", args, err)
		return
	}
	os.Exit(exitCode)
}

func handleEchoCommand(_, args string) {
	fmt.Fprintln(os.Stdout, args)
}

func handleTypeCommand(_, args string) {
	if slices.Contains(builtInCommands, args) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args)
		return
	}

	if path, ok := searchExecPath(args); ok {
		fmt.Fprintf(os.Stdout, "%s is %s\n", args, path)
		return
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", args)
}

func searchExecPath(execName string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))

	for _, p := range paths {
		entries, _ := os.ReadDir(p)
		for _, entry := range entries {
			if execName == entry.Name() {
				path := filepath.Join(p, entry.Name())
				stats, err := os.Stat(path)
				if err != nil {
					continue
				}
				mode := stats.Mode()
				if isExecutable(mode) {
					return path, true
				}
			}
		}
	}
	return "", false
}

func isExecutable(mode fs.FileMode) bool {
	modeStr := mode.String()
	ownerCanExec := modeStr[3] == 'x'
	groupCanExec := modeStr[6] == 'x'
	otherCanExec := modeStr[9] == 'x'
	return ownerCanExec || groupCanExec || otherCanExec
}
