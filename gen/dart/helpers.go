package dart

import (
	"golang.org/x/tools/imports"
)

// FormatSource is gofmt with addition of removing any unused imports.
func FormatSource(source []byte) ([]byte, error) {
	return imports.Process("", source, &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 8,
	})
}
