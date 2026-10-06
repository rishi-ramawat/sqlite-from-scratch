// Command sqlite is the entry point for the sqlite-from-scratch database shell.
package main

import (
	"fmt"
	"os"

	"github.com/rishi-ramawat/sqlite-from-scratch/internal/repl"
)

func main() {
	if err := repl.Run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
