package gen

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestInferLocalTemplateVersion(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed")
	}

	tests := []struct {
		name    string
		setup   [][]string // git commands run in the template dir; nil = not a git repo
		version string     // regexp matched against TmplVersion after the dir
	}{
		{
			name:    "not a git repo",
			version: `^$`,
		},
		{
			name:    "no tags",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}},
			version: `^$`,
		},
		{
			name:    "lightweight tag",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}, {"tag", "v0.7.0"}},
			version: `^@v0\.7\.0$`,
		},
		{
			name:    "annotated tag",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}, {"tag", "-a", "v0.7.0", "-m", "v0.7.0"}},
			version: `^@v0\.7\.0$`,
		},
		{
			name:    "commits after tag",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}, {"tag", "v0.7.0"}, {"commit", "--allow-empty", "-m", "next"}},
			version: `^@v0\.7\.0-1-g[0-9a-f]+$`,
		},
		{
			name:    "tag with slash is ignored",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}, {"tag", "release/v1.0.0"}},
			version: `^$`,
		},
		{
			name:    "tag with semicolon is ignored",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}, {"tag", "v1;2"}},
			version: `^$`,
		},
		{
			name:    "tag with quote is ignored",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}, {"tag", `v1"2`}},
			version: `^$`,
		},
		{
			name:    "tag with at sign is ignored",
			setup:   [][]string{{"init"}, {"add", "."}, {"commit", "-m", "init"}, {"tag", "v1@2"}},
			version: `^$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			if err := os.WriteFile(filepath.Join(dir, "main.go.tmpl"), []byte(`{{define "main"}}{{end}}`), 0644); err != nil {
				t.Fatal(err)
			}
			for _, args := range tt.setup {
				git(t, dir, args...)
			}
			if tt.setup == nil {
				// TMPDIR may live inside a repo; stop git from walking up to it.
				t.Setenv("GIT_CEILING_DIRECTORIES", filepath.Dir(dir))
			}

			s, err := NewTemplateSource(dir, &Config{})
			if err != nil {
				t.Fatal(err)
			}
			if _, err := s.loadTemplates(); err != nil {
				t.Fatal(err)
			}

			if !strings.HasPrefix(s.TmplVersion, dir) {
				t.Fatalf("TmplVersion = %q, want prefix %q", s.TmplVersion, dir)
			}
			got := strings.TrimPrefix(s.TmplVersion, dir)
			if !regexp.MustCompile(tt.version).MatchString(got) {
				t.Errorf("TmplVersion = %q, want %s%s", s.TmplVersion, dir, tt.version)
			}
		})
	}
}

func git(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-c", "user.name=test", "-c", "user.email=test@example.com", "-c", "tag.gpgSign=false", "-c", "commit.gpgSign=false"}, args...)...)
	cmd.Dir = dir
	// Isolate from the developer's global git config and any enclosing repo.
	cmd.Env = append(envWithout(os.Environ(), "GIT_DIR", "GIT_WORK_TREE", "GIT_INDEX_FILE"),
		"GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_NOSYSTEM=1", "GIT_CEILING_DIRECTORIES="+filepath.Dir(dir))
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
	}
}
