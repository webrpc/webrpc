package gen

import "golang.org/x/tools/imports"

// FormatGoSource is gofmt with addition of removing any unused imports.
func FormatGoSource(source []byte) ([]byte, error) {
	return imports.Process("", source, &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 8,
	})
}
