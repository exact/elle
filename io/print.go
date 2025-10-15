package io

import (
	"fmt"
	"os"
)

// Puts writes `a` to the standard output, and adds a newline at the end.
// It returns amount of bytes written and any write error encountered.
func Puts(a ...any) (int, error) {
	return fmt.Fprintln(os.Stdout, a...)
}
