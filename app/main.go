package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command := scanner.Text()
		fmt.Fprintf(os.Stdout, "%s: command not found", command)
	}

}
