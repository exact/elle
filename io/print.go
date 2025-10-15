package io

import (
	"fmt"
	"os"
)

// Puts writes a to standard output, and adds a newline at the end.
// It returns amount of bytes written and any write error encountered.
func Puts(a ...any) (int, error) {
	return fmt.Fprintln(os.Stdout, a...)
}

// Warn writes a to standard error, and adds a newline at the end.
// It returns amount of bytes written and any write error encountered.
func Warn(a ...any) (int, error) {
	return fmt.Fprintln(os.Stderr, a...)
}
