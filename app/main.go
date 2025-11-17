package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

var builtInCommands = []string{"echo", "exit", "type", "pwd"}

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
	cmd := parseCommand(c)
	args := parseArgs(c)

	switch cmd {
	case "exit":
		handleExitCommand(cmd, args)
	case "echo":
		handleEchoCommand(cmd, args)
	case "type":
		handleTypeCommand(cmd, args)
	case "pwd":
		handlePWDCommand(cmd, args)
	case "cd":
		handleCDCommand(cmd, args)
	default:
		handleOtherCommand(cmd, args)
	}
}

func handleExitCommand(_ string, args []string) {
	exitCode, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Fprintf(os.Stdout, "Cannot handle exit code: %v, error: %v", args, err)
		return
	}
	os.Exit(exitCode)
}

func handleEchoCommand(_ string, args []string) {
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
}

func handleTypeCommand(_ string, args []string) {
	cmd := args[0]
	if slices.Contains(builtInCommands, cmd) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
		return
	}

	if path, ok := searchExecPath(cmd); ok {
		fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, path)
		return
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
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

func handleOtherCommand(command string, args []string) {
	if _, ok := searchExecPath(command); ok {
		execCommand(command, args)
		return
	}
	fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
}

func execCommand(path string, args []string) {
	cmd := exec.Command(path, args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprint(os.Stdout, string(out))
		return
	}
	fmt.Fprint(os.Stdout, string(out))
}

func handlePWDCommand(_ string, _ []string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "%s\n", path)
}

func handleCDCommand(_ string, args []string) {
	path := correctPath(args[0])
	errMsg := fmt.Sprintf("cd: %s: No such file or directory\n", path)

	info, err := os.Stat(path)

	if err != nil {
		fmt.Fprint(os.Stdout, errMsg)
		return
	}

	if !info.IsDir() {
		fmt.Fprint(os.Stdout, errMsg)
		return
	}

	err = os.Chdir(path)
	if err != nil {
		fmt.Fprint(os.Stdout, errMsg)
		return
	}
}

func correctPath(path string) string {
	homeSign := "~"
	if strings.Contains(path, homeSign) {
		home := os.Getenv("HOME")
		return strings.ReplaceAll(path, homeSign, home)
	}
	return path
}

func parseCommand(command string) string {
	return strings.Split(command, " ")[0]
}

func parseArgs(command string) []string {
	args := strings.Join(strings.SplitN(command, " ", 2)[1:], " ")

	inSingleQuotes := false
	inDoubleQuotes := false
	isQuoted := false
	var acc []string
	var buf []rune

	flush := func() {
		if len(buf) > 0 {
			acc = append(acc, string(buf))
			buf = buf[:0]
		}
	}

	for _, c := range args {
		if isQuoted {
			isQuoted = false

			if inSingleQuotes {
				buf = append(buf, '\\')
				buf = append(buf, c)
				continue
			}

			if inDoubleQuotes {
				if c == '"' || c == '\\' {
					buf = append(buf, c)
				} else {
					buf = append(buf, '\\')
					buf = append(buf, c)
				}
				continue
			}

			// default
			buf = append(buf, c)
			continue
		}
		
		switch c {
		case '\\':
			isQuoted = true
		case ' ':
			if inDoubleQuotes || inSingleQuotes {
				buf = append(buf, c)
				continue
			}
			flush()
		case '"':
			if inSingleQuotes {
				buf = append(buf, c)
				continue
			}
			inDoubleQuotes = !inDoubleQuotes
		case '\'':
			if inDoubleQuotes {
				buf = append(buf, c)
				continue
			}
			inSingleQuotes = !inSingleQuotes
		default:
			buf = append(buf, c)
		}
	}

	flush()

	for i := range acc {
		acc[i] = correctPath(acc[i])
	}
	return acc
}
