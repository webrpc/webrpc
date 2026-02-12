export PATH = $(shell echo $$PWD/bin:$$PATH)

all:
	@echo "****************************************"
	@echo "**              webrpc                **"
	@echo "****************************************"
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@awk -F'[ :]' '/^#+/ {comment=$$0; gsub(/^#+[ ]*/, "", comment)} !/^(_|all:)/ && /^([A-Za-z_-]+):/ && !seen[$$1]++ {printf "  %-24s %s\n", $$1, (comment ? "- " comment : ""); comment=""} !/^#+/ {comment=""}' Makefile

# Build webrpc-gen
build:
	go build -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe --tags)" -o ./bin/webrpc-gen ./cmd/webrpc-gen

# Build webrpc-test
build-test:
	go build -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe --tags)" -o ./bin/webrpc-test ./cmd/webrpc-test

# Install webrpc-gen and webrpc-test binaries locally
install:
	go install -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe --tags)" ./cmd/webrpc-gen
	go install -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe --tags)" ./cmd/webrpc-test

clean:
	rm -rf ./bin

# Regenerate examples and tests using latest templates (see go.mod)
generate: build
	go generate -v -x ./...
	for i in _examples/*; do echo $$i; make -C $$i generate || exit 1; done
	# Replace webrpc version in all generated files to avoid git conflicts.
	git grep -l "$$(git describe --tags)" | xargs sed -i -e "s/@$$(git describe --tags)//g"
	sed -i "/$$(git describe --tags)/d" tests/schema/test.debug.gen.txt

# Upgrade Go dependencies
dep-upgrade-all:
	go get -u go@1.23 ./...

dep-upgrade-templates:
	go get github.com/webrpc/gen-dart@latest
	go get github.com/webrpc/gen-golang@latest
	go get github.com/webrpc/gen-javascript@latest
	go get github.com/webrpc/gen-kotlin@latest
	go get github.com/webrpc/gen-openapi@latest
	go get github.com/webrpc/gen-typescript@latest

# Run git diff and fail on any local changes
diff:
	git diff --color --ignore-all-space --ignore-blank-lines --exit-code

# Run all tests
test: test-go test-interoperability

# Run Go tests
test-go: generate
	go test -v ./...

# Run interoperability test suite
test-interoperability: build-test
	echo "Running interoperability test suite"; \
		./bin/webrpc-test -server -port=9988 -timeout=2s & \
		until nc -z localhost 9988; do sleep 0.1; done; \
		./bin/webrpc-test -client -url=http://localhost:9988; \
		wait

# Update ridl golden examples
update-ridl-test-golden-examples:
	cd ./schema/ridl && go test -update=./_example/example1-golden.json
	cd ./schema/ridl && go test -update=./_example/example2-golden.json
