package io

import "fmt"

func S(f string, a ...any) string {
	return fmt.Sprintf(f, a...)
}
