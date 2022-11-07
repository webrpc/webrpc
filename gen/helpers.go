package gen

import (
	"fmt"
	"strings"

	"golang.org/x/tools/imports"
)

// Backward compatibility with webrpc-gen v0.6.0.
func getOldTarget(target string) string {
	switch target {
	case "go":
		return "github.com/webrpc/gen-golang@v0.6.0"
	case "ts":
		return "github.com/webrpc/gen-typescript@v0.6.0"
	case "js":
		return "github.com/webrpc/gen-javascript@v0.6.0"
	}
	return target
}

func isGolangTarget(target string) bool {
	target, _, _ = strings.Cut(target, "@")

	if strings.HasSuffix(target, "gen-golang") {
		return true
	}

	if target == "go" {
		return true
	}

	return false
}

// Format Go source and & update imports.
func formatGoSource(source []byte) (string, error) {
	formatted, err := imports.Process("", source, &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 8,
	})
	if err != nil {
		return "", fmt.Errorf("failed to format generated Go source: %w", err)
	}
	return string(formatted), nil
}
