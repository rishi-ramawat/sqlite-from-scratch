// Package repl implements the interactive read-eval-print loop for the database shell.
package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Run reads commands from in and writes results to out until EOF or ".exit".
func Run(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, "db > ")
		if !scanner.Scan() {
			fmt.Fprintln(out)
			return scanner.Err()
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, ".") {
			if done := handleMetaCommand(line, out); done {
				return nil
			}
			continue
		}
		fmt.Fprintf(out, "Unrecognized statement: %q\n", line)
	}
}

// handleMetaCommand processes dot-commands such as ".exit".
// It returns true when the REPL should terminate.
func handleMetaCommand(cmd string, out io.Writer) bool {
	switch cmd {
	case ".exit":
		return true
	default:
		fmt.Fprintf(out, "Unrecognized command: %q\n", cmd)
		return false
	}
}
