package gen

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Similar to "printf" but instead of writing into the generated
// output file, stderrPrintf writes to webrpc-gen CLI stderr.
// Useful for printing template errors / debugging.
func stderrPrintf(format string, a ...interface{}) string {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	return ""
}

// Similar to "print" but instead of writing into the generated
// output file, stderrPrint writes to webrpc-gen CLI stderr.
// Useful for printing template errors / debugging.
func stderrPrint(a ...interface{}) string {
	_, _ = fmt.Fprint(os.Stderr, a...)
	return ""
}

// Terminate template execution with a status code.
// Useful for exiting early or for printing fatal errors from within templates.
func exit(code int) error {
	os.Exit(code)
	return nil
}

func minVersion(version string, minVersion string) bool {
	major, minor, err := parseMajorMinorVersion(version)
	if err != nil {
		panic(fmt.Sprintf("minVersion: unexpected version %q", version))
	}

	minMajor, minMinor, err := parseMajorMinorVersion(minVersion)
	if err != nil {
		panic(fmt.Sprintf("minVersion: unexpected min version %q", minVersion))
	}

	if minMajor > major {
		return false
	}

	if minMinor > minor {
		return false
	}

	return true
}

func parseMajorMinorVersion(version string) (major int, minor int, err error) {
	version = strings.TrimPrefix(version, "v")
	parts := strings.Split(version, ".")

	major, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	if len(parts) > 1 {
		minor, err = strconv.Atoi(parts[1])
		if err != nil {
			return
		}
	}

	return
}

func hasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}
