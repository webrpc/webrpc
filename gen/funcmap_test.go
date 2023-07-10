package gen

import (
	"fmt"
	"testing"
)

func TestMinVersion(t *testing.T) {
	tt := []struct {
		Version    string
		MinVersion string
		Result     bool
	}{
		{"v1.0.0", "v1", true},
		{"v1.0.0", "v1.0", true},
		{"v1.0.0", "v1.0.0", true},
		{"v2.5.8", "v1", true},
		{"v2.5.8", "v2.0", true},
		{"v2.5.8", "v1.1", true},
		{"v2.5.8", "v1.5.8", true},
		{"v2.5.8", "v2", true},
		{"v2.5.8", "v2.5", true},
		{"v2.5.8", "v2.5.5", true},

		{"v1.0.0", "v2", false},
		{"v1.0.0", "v2.0", false},
		{"v1.0.0", "v1.1", false},
		{"v1.0.0", "v1.5.8", false},
		{"v1.0.0", "v2", false},
		{"v2.5.8", "v3", false},
		{"v2.5.8", "v2.6", false},
		{"v2.5.8", "v2.6.0", false},
		{"v2.5.8", "v2.6.6", false},
	}

	for _, tc := range tt {
		if result := minVersion(tc.Version, tc.MinVersion); result != tc.Result {
			t.Errorf("unexpected result: minVersion(%q, %q) = %v", tc.Version, tc.MinVersion, result)
		}
	}
}

func TestParseMajorMinorVersion(t *testing.T) {
	tt := []struct {
		Version string
		Major   int
		Minor   int
		Error   bool
	}{
		{"v0.7", 0, 7, false},
		{"v0.7.0", 0, 7, false},
		{"v0.7.5", 0, 7, false},
		{"v0.7.0-dev", 0, 7, false},
		{"v0.7.x-dev", 0, 7, false},

		{"v1", 1, 0, false},
		{"v1.0", 1, 0, false},
		{"v1.1", 1, 1, false},
		{"v1.0.1", 1, 0, false},
		{"v1.1.1", 1, 1, false},
		{"v2", 2, 0, false},
		{"v2.0", 2, 0, false},
		{"v2.2", 2, 2, false},
		{"v2.2.0", 2, 2, false},
		{"v2.2.2", 2, 2, false},

		{"1", 1, 0, false},
		{"1.0", 1, 0, false},
		{"1.1", 1, 1, false},
		{"1.0.1", 1, 0, false},
		{"1.1.1", 1, 1, false},
		{"2", 2, 0, false},
		{"2.0", 2, 0, false},
		{"2.2", 2, 2, false},
		{"2.2.0", 2, 2, false},
		{"2.2.2", 2, 2, false},

		// Errors:
		{"", 0, 0, true},
		{"err", 0, 0, true},
	}

	for _, tc := range tt {
		major, minor, err := parseMajorMinorVersion(tc.Version)
		if err != nil && !tc.Error {
			t.Errorf("unexpected error: parseMajorMinorVersion(%q): %v", tc.Version, err)
		}
		if err == nil && tc.Error {
			t.Errorf("expected error: parseMajorMinorVersion(%q)", tc.Version)
		}
		if major != tc.Major {
			t.Errorf("major, _, _ := parseMajorMinorVersion(%q): got %v, expected %v", tc.Version, major, tc.Major)
		}
		if minor != tc.Minor {
			t.Errorf("_, minor, _ := parseMajorMinorVersion(%q): got %v, expected %v", tc.Version, minor, tc.Minor)
		}
	}
}

func TestArray(t *testing.T) {
	var a, b, c, d interface{} = "a", "b", "c", "d"

	arr := array(c, d)
	if got := len(arr); got != 2 {
		t.Errorf("array: expected two elements, got %v", got)
	}

	arr = appendFn(arr, a, b)
	if got := len(arr); got != 4 {
		t.Errorf("append: expected four elements, got %v", got)
	}

	sorted := sortFn(arr)
	if got := fmt.Sprintf("%v", sorted); got != "[a b c d]" {
		t.Errorf("sort: expected sorted array, got %v", got)
	}

	if got, _ := first(sorted); got != "a" {
		t.Errorf("first: expected a, got %v", got)
	}

	if got, _ := last(sorted); got != "d" {
		t.Errorf("last: expected d, got %v", got)
	}

	if got, err := first([]interface{}{}); err == nil || got != "" {
		t.Errorf("first(empty): expected error")
	}

	if got, err := last([]interface{}{}); err == nil || got != "" {
		t.Errorf("last(empty): expected error")
	}
}
