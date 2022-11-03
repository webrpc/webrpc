package gen

import "testing"

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
