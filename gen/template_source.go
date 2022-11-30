package gen

import (
	"context"
	"fmt"
	"hash/fnv"
	"log"
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
	s, err := NewTemplateSource(proto, target, config)
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
	templateCacheTime              = 1 * time.Hour
	templateCacheTimestampFilename = ".webrpc-gen-timestamp"
	templateCacheInfoFilename      = ".webrpc-gen-info"
)

type TemplateSource struct {
	tmpl   *template.Template
	proto  *schema.WebRPCSchema
	target string
	config *Config

	IsLocal     bool
	TmplDir     string
	TmplVersion string // git url and hash, or the local dir same as TmplDir
}

func NewTemplateSource(proto *schema.WebRPCSchema, target string, config *Config) (*TemplateSource, error) {
	tmpl := template.New(target).Funcs(templateFuncMap(proto, config.TemplateOptions))
	return &TemplateSource{
		tmpl:   tmpl,
		proto:  proto,
		target: target,
		config: config,
	}, nil
}

func (s *TemplateSource) loadTemplates() (*template.Template, error) {
	if isLocalDir(s.target) {
		// from local directory
		s.IsLocal = true
		tmpl, err := s.tmpl.ParseGlob(filepath.Join(s.target, "/*.tmpl"))
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
			return nil, fmt.Errorf("failed to load templates from git %s: %w", s.target, err)
		}
		return tmpl, err
	}
}

func (s *TemplateSource) loadRemote() (*template.Template, error) {
	var sourceFS http.FileSystem

	cacheDir, cacheFS, cacheAvailable, cacheTS, err := s.openCacheDir()
	if err != nil {
		return nil, err
	}
	s.IsLocal = false
	s.TmplDir = cacheDir

	// cache is new, so lets fetch from git
	if !cacheAvailable || s.config.RefreshCache || time.Now().Unix()-cacheTS > int64(templateCacheTime.Seconds()) {
		sourceFS, err = gitfs.New(context.Background(), s.target) //, gitfs.OptPrefetch(true), gitfs.OptGlob("/*.tmpl"))
		s.TmplVersion = s.target

		if err != nil {
			// error occured reading from git, if cache is available, use that instead
			if cacheAvailable {
				sourceFS = http.Dir(cacheDir)
			} else {
				return nil, fmt.Errorf("failed to load templates from remote git repository %s: %w", s.target, err)
			}

		} else {
			// using git remote source -- lets cache the files too
			err := s.syncTemplates(s.target, sourceFS, cacheFS, cacheDir)
			if err != nil {
				// in case of error, just print a warning and carry on
				log.Println("[warning] error syncing git templates to local cache:", err.Error())
			} else {
				// read from our cache
				sourceFS = cacheFS
			}
		}
	} else {
		// load from cache directory
		sourceFS = cacheFS
	}

	// read template version info
	if cacheAvailable && s.TmplVersion == "" {
		tmplVersion, _ := os.ReadFile(filepath.Join(cacheDir, templateCacheInfoFilename))
		s.TmplVersion = strings.TrimSpace(string(tmplVersion))
	}

	// parse the template files from the source
	tmpl, err := vfstemplate.ParseGlob(sourceFS, s.tmpl, "/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	return tmpl, nil
}

func (s *TemplateSource) syncTemplates(target string, remoteFS, cacheFS http.FileSystem, cacheDir string) error {
	filenames, err := vfspath.Glob(remoteFS, "/*.tmpl")
	if err != nil {
		return err
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

	now := time.Now().Unix()
	data := []byte(fmt.Sprintf("%d", now))

	err = os.WriteFile(filepath.Join(cacheDir, templateCacheTimestampFilename), data, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(cacheDir, templateCacheInfoFilename), []byte(strings.TrimSpace(target)), 0755)
	if err != nil {
		return err
	}

	return nil
}

func (s *TemplateSource) openCacheDir() (string, http.FileSystem, bool, int64, error) {
	cacheDir, _ := s.getTmpCacheDir()
	if cacheDir == "" {
		// unable to find OS temp dir, but we don't error -- although
		// we probably should print a warning
		return "", nil, false, 0, nil
	}

	// delete the directory if asked to refresh
	if s.config.RefreshCache {
		os.RemoveAll(cacheDir)
	}

	// create the directory if it doesn't exist
	_, err := os.Stat(cacheDir)
	if !os.IsExist(err) {
		err := os.MkdirAll(cacheDir, 0755)
		if err != nil {
			return "", nil, false, 0, fmt.Errorf("unable to create directory for template cache at %s: %w", cacheDir, err)
		}
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

	return cacheDir, cacheFS, available, ts, nil
}

func (s *TemplateSource) getTmpCacheDir() (string, error) {
	dir := os.TempDir()
	if dir == "" {
		return "", fmt.Errorf("unable to determine OS temp dir")
	}

	// derive a deterministic folder for this template source
	parts := strings.Split(s.target, "/")
	name := parts[len(parts)-1]

	hash := fnv.New32a()
	hash.Write([]byte(s.target))

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
