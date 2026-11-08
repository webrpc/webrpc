package gen

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
)

func TestInferLocalTemplateVersion(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed")
	}

	tests := []struct {
		name    string
		setup   []string // git commands run in the template dir; nil = not a git repo
		version string   // regexp matched against TmplVersion after the dir
	}{
		{
			name:    "not a git repo",
			version: `^$`,
		},
		{
			name:    "no tags",
			setup:   []string{"init", "add .", "commit -m init"},
			version: `^$`,
		},
		{
			name:    "lightweight tag",
			setup:   []string{"init", "add .", "commit -m init", "tag v0.7.0"},
			version: `^@v0\.7\.0$`,
		},
		{
			name:    "annotated tag",
			setup:   []string{"init", "add .", "commit -m init", "tag -a v0.7.0 -m v0.7.0"},
			version: `^@v0\.7\.0$`,
		},
		{
			name:    "commits after tag",
			setup:   []string{"init", "add .", "commit -m init", "tag v0.7.0", "commit --allow-empty -m next"},
			version: `^@v0\.7\.0-1-g[0-9a-f]+$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			if err := os.WriteFile(filepath.Join(dir, "main.go.tmpl"), []byte(`{{define "main"}}{{end}}`), 0644); err != nil {
				t.Fatal(err)
			}
			for _, args := range tt.setup {
				git(t, dir, args)
			}

			s, err := NewTemplateSource(dir, &Config{})
			if err != nil {
				t.Fatal(err)
			}
			if _, err := s.loadTemplates(); err != nil {
				t.Fatal(err)
			}

			got := s.TmplVersion[len(dir):]
			if !regexp.MustCompile(tt.version).MatchString(got) {
				t.Errorf("TmplVersion = %q, want %s%s", s.TmplVersion, dir, tt.version)
			}
		})
	}
}

func git(t *testing.T, dir string, args string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-c", "user.name=test", "-c", "user.email=test@example.com", "-c", "tag.gpgSign=false", "-c", "commit.gpgSign=false"}, splitArgs(args)...)...)
	cmd.Dir = dir
	// Isolate from the developer's global git config and any enclosing repo.
	cmd.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_NOSYSTEM=1", "GIT_CEILING_DIRECTORIES="+filepath.Dir(dir))
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %s: %v\n%s", args, err, out)
	}
}

func splitArgs(s string) []string {
	return regexp.MustCompile(`\s+`).Split(s, -1)
}
