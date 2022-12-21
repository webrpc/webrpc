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
	go build -o ./bin/webrpc-gen ./cmd/webrpc-gen

build-test:
	go build -o ./bin/webrpc-test ./cmd/webrpc-test

clean:
	rm -rf ./bin

install:
	go install ./cmd/webrpc-gen

generate: build
	go generate ./...
	@for i in _examples/*/Makefile; do           \
		echo; echo $$ cd $$i \&\& make generate; \
		cd $$(dirname $$i);                      \
		make generate || exit 1;                 \
		cd ../../;                               \
	done

dep:
	go mod tidy

dep-upgrade-all:
	go get -u ./...
	@$(MAKE) dep

diff:
	git diff --color --ignore-all-space --ignore-blank-lines --exit-code

test: generate build-test
	go test -v ./...
	echo "Running imperoperability test"; \
		./bin/webrpc-test -server -port=9988 -timeout=2s & \
		sleep 0.5; \
		./bin/webrpc-test -client -url=http://localhost:9988; \
		wait
