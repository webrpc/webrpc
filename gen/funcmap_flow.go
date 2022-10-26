package gen

import (
	"fmt"
	"os"
)

func printfStderr(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(os.Stderr, format, a...)
	return err
}

func exit(code int) error {
	os.Exit(code)
	return nil
}
