export PATH = $(shell echo $$PWD/bin:$$PATH)

all:
	@echo "*****************************************"
	@echo "**             WebRPC Dev             **"
	@echo "*****************************************"
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo ""
	@echo " + Testing:"
	@echo "   - test"
	@echo "   - generate"
	@echo "   - diff"
	@echo ""
	@echo " + Builds:"
	@echo "   - build"
	@echo "   - build-test"
	@echo "   - clean"
	@echo "   - install"
	@echo ""
	@echo " + Dep management:"
	@echo "   - dep"
	@echo "   - dep-upgrade-all"
	@echo ""

build:
	go build -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe)" -o ./bin/webrpc-gen ./cmd/webrpc-gen

build-test:
	go build -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe)" -o ./bin/webrpc-test ./cmd/webrpc-test

clean:
	rm -rf ./bin

install:
	go install -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe)" ./cmd/webrpc-gen
	go install -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=$$(git describe)" ./cmd/webrpc-test

generate: build
	go generate -v -x ./...
	cd _examples/ && go generate -x ./...
	@for i in _examples/*/Makefile; do           \
		echo; echo $$ cd $$(dirname $$i) \&\& make generate; \
		cd $$(dirname $$i);                      \
		make generate || exit 1;                 \
		cd ../../;                               \
	done
	# Replace webrpc version in all generated files to avoid git conflicts.
	git grep -l "$$(git describe)" | xargs sed -i -e "s/@$$(git describe)//g"
	sed -i "/$$(git describe)/d" tests/schema/test.debug.gen.txt

dep:
	go mod tidy

dep-upgrade-all:
	go get -u ./...
	@$(MAKE) dep

diff:
	git diff --color --ignore-all-space --ignore-blank-lines --exit-code

test: test-go test-interoperability

test-go: generate
	go test -v ./...

test-interoperability: build-test
	echo "Running interoperability test suite"; \
		./bin/webrpc-test -server -port=9988 -timeout=2s & \
		until nc -z localhost 9988; do sleep 0.1; done; \
		./bin/webrpc-test -client -url=http://localhost:9988; \
		wait
