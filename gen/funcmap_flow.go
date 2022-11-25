package gen

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Similar to "printf" but instead of writing into the generated
// output file, stderrPrintf writes to webrpc-gen CLI stderr.
// Useful for printing template errors or for template debugging.
func stderrPrintf(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(os.Stderr, format, a...)
	return err
}

// Terminate template execution. Useful for fatal errors.
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
