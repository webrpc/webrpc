package gen

import (
	"context"
	"fmt"
	"hash/fnv"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/posener/gitfs"
	"github.com/shurcooL/httpfs/path/vfspath"
	"github.com/shurcooL/httpfs/text/vfstemplate"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/webrpc/webrpc/schema"
)

func loadTemplates(proto *schema.WebRPCSchema, target string, config *Config) (*template.Template, *TemplateSource, error) {
	s, err := NewTemplateSource(target, config)
	if err != nil {
		return nil, nil, err
	}
	tmpl, err := s.loadTemplates()
	if err != nil {
		return nil, nil, err
	}
	return tmpl, s, nil
}

// period of time before we attempt to refetch from git source again.
// in case of a failure, we will use the local cache.
const (
	templateCacheTTL               = 1 * time.Hour
	templateCacheTimestampFilename = ".webrpc-gen-timestamp"
	templateCacheInfoFilename      = ".webrpc-gen-info"
)

type TemplateSource struct {
	tmpl   *template.Template
	target string
	config *Config

	IsLocal         bool
	TmplDir         string
	TmplVersion     string // git url and hash, or the local dir same as TmplDir
	CacheAge        time.Duration
	CacheRefreshErr error
}

func NewTemplateSource(target string, config *Config) (*TemplateSource, error) {
	tmpl := template.New(target).Funcs(templateFuncMap(config.TemplateOptions))
	return &TemplateSource{
		tmpl:   tmpl,
		target: target,
		config: config,
	}, nil
}

func (s *TemplateSource) loadTemplates() (*template.Template, error) {
	if isLocalDir(s.target) {
		// from local directory
		s.IsLocal = true
		tmpl, err := s.tmpl.ParseGlob(filepath.Join(s.target, "/*.go.tmpl"))
		if err != nil {
			return nil, fmt.Errorf("failed to load templates from %s: %w", s.target, err)
		}
		s.TmplDir = s.target
		s.TmplVersion = s.target
		return tmpl, err
	} else {
		// from remote git or cache source
		s.IsLocal = false
		s.target = s.inferRemoteTarget(s.target)
		tmpl, err := s.loadRemote()
		if err != nil {
			return nil, fmt.Errorf("failed to load templates from %s: %w", s.target, err)
		}
		return tmpl, err
	}
}

func (s *TemplateSource) loadRemote() (*template.Template, error) {
	var err error
	var sourceFS http.FileSystem

	cacheAvailable, cacheDir, cacheFS, cacheAge := s.openCacheDir()
	s.TmplDir = cacheDir

	if !cacheAvailable || s.config.RefreshCache || cacheAge > templateCacheTTL {
		s.TmplVersion = s.target
		// fetch from remote git
		sourceFS, err = gitfs.New(context.Background(), s.target, gitfs.OptPrefetch(true), gitfs.OptGlob("*.go.tmpl"))
		if err != nil {
			s.CacheRefreshErr = err
			// load from cache, if available
			if cacheAvailable {
				sourceFS = cacheFS
				s.CacheAge = cacheAge
			} else {
				return nil, fmt.Errorf("failed to fetch from remote git: %w", err)
			}
		} else {
			// cache remote git
			if err := s.cacheTemplates(s.target, sourceFS, cacheDir); err != nil {
				s.CacheRefreshErr = err
			}
		}
	} else {
		// load from cache
		sourceFS = cacheFS
		s.CacheAge = cacheAge
	}

	// read template version info
	if cacheAvailable && s.TmplVersion == "" {
		tmplVersion, _ := os.ReadFile(filepath.Join(cacheDir, templateCacheInfoFilename))
		s.TmplVersion = strings.TrimSpace(string(tmplVersion))
	}

	// parse the template files from the source
	tmpl, err := vfstemplate.ParseGlob(sourceFS, s.tmpl, "/*.go.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	return tmpl, nil
}

func (s *TemplateSource) cacheTemplates(target string, remoteFS http.FileSystem, cacheDir string) error {
	// create empty cache directory
	if _, err := os.Stat(cacheDir); os.IsExist(err) {
		if err := os.RemoveAll(cacheDir); err != nil {
			return err
		}
	}

	filenames, err := vfspath.Glob(remoteFS, "/*.go.tmpl")
	if err != nil {
		return err
	}

	if len(filenames) == 0 {
		return fmt.Errorf("no template files were found in target %s", target)
	}

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return fmt.Errorf("unable to create directory for template cache at %s: %w", cacheDir, err)
	}

	for _, filename := range filenames {
		data, err := vfsutil.ReadFile(remoteFS, filename)
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(cacheDir, filename), data, 0755)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(filepath.Join(cacheDir, templateCacheInfoFilename), []byte(strings.TrimSpace(target)), 0755)
	if err != nil {
		return err
	}

	now := time.Now()
	data := []byte(fmt.Sprintf("%d", now.Unix()))

	err = os.WriteFile(filepath.Join(cacheDir, templateCacheTimestampFilename), data, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (s *TemplateSource) openCacheDir() (bool, string, http.FileSystem, time.Duration) {
	cacheDir, _ := s.getTmpCacheDir()
	if cacheDir == "" {
		// unable to find OS temp dir, but we don't error -- although
		// we probably should print a warning
		return false, "", nil, 0
	}

	// convert local fs to http filesystem
	cacheFS := http.Dir(cacheDir)

	// read cache timestamp file to determine availability
	available := false
	ts := int64(0)
	b, _ := vfsutil.ReadFile(cacheFS, templateCacheTimestampFilename)
	if len(b) > 0 {
		ts, _ = strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
		if ts > 0 {
			available = true
		}
	}

	return available, cacheDir, cacheFS, time.Since(time.Unix(ts, 0))
}

func (s *TemplateSource) getTmpCacheDir() (string, error) {
	dir := os.TempDir()
	if dir == "" {
		return "", fmt.Errorf("unable to determine OS temp dir")
	}

	// derive a deterministic folder for this template source
	hash := fnv.New32a()
	hash.Write([]byte(s.target))

	parts := strings.Split(s.target, "/")
	name := parts[len(parts)-1]

	return filepath.Join(dir, "webrpc-cache", fmt.Sprintf("%d-%s", hash.Sum32(), name)), nil
}

func (s *TemplateSource) inferRemoteTarget(target string) string {
	// extra check to ensure its not a local dir
	if isLocalDir(target) {
		return target
	}

	// determine if a url is passed or just a gen-XXX name
	parts := strings.Split(target, "/")

	// just a name, so by convention assume the default target of the webrpc org
	if len(parts) == 1 {
		return fmt.Sprintf("github.com/webrpc/gen-%s", strings.ToLower(target))
	}

	// accept the target as is
	return target
}

func isLocalDir(target string) bool {
	return strings.HasPrefix(target, "/") || strings.HasPrefix(target, ".")
}
