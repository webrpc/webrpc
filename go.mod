module github.com/webrpc/webrpc

// Webrpc supports the two most recent major versions of Go.
go 1.24.0

toolchain go1.25.0

// replace github.com/webrpc/gen-golang => ../gen-golang
// replace github.com/webrpc/gen-typescript => ../gen-typescript
// replace github.com/webrpc/gen-javascript => ../gen-javascript
// replace github.com/webrpc/gen-openapi => ../gen-openapi
// replace github.com/webrpc/gen-kotlin => ../gen-kotlin

require (
	github.com/Masterminds/sprig/v3 v3.3.0
	github.com/davecgh/go-spew v1.1.1
	github.com/golang-cz/textcase v1.2.1
	github.com/google/go-cmp v0.7.0
	github.com/posener/gitfs v1.2.2
	github.com/shurcooL/httpfs v0.0.0-20230704072500-f1e31cf0ba5c
	github.com/stretchr/testify v1.11.1
	github.com/webrpc/gen-dart v0.1.1
	github.com/webrpc/gen-golang v0.24.0
	github.com/webrpc/gen-javascript v0.13.0
	github.com/webrpc/gen-kotlin v0.1.0
	github.com/webrpc/gen-openapi v0.17.1
	github.com/webrpc/gen-typescript v0.23.1
	golang.org/x/tools v0.36.0
)

require (
	dario.cat/mergo v1.0.2 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ProtonMail/go-crypto v1.3.0 // indirect
	github.com/cloudflare/circl v1.6.1 // indirect
	github.com/cyphar/filepath-securejoin v0.6.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.6.2 // indirect
	github.com/go-git/go-git/v5 v5.16.5 // indirect
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v1.4.0 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pjbgf/sha1cd v0.5.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/posener/diff v0.0.1 // indirect
	github.com/sergi/go-diff v1.4.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/skeema/knownhosts v1.3.1 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/mod v0.27.0 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
