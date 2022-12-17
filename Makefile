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
	@echo ""
	@echo " + Builds:"
	@echo "   - build"
	@echo "   - clean"
	@echo "   - generate"
	@echo ""
	@echo " + Dep management:"
	@echo "   - dep"
	@echo "   - dep-upgrade-all"
	@echo ""

build:
	go build -o ./bin/webrpc-gen ./cmd/webrpc-gen
	./bin/webrpc-gen -schema=./tests/schema/api.ridl -target=golang -pkg=client -client -out=./tests/client/client.go
	./bin/webrpc-gen -schema=./tests/schema/api.ridl -target=golang -pkg=server -server -out=./tests/server/server.gen.go
	go build -o ./bin/webrpc-server ./cmd/webrpc-server

clean:
	rm -rf ./bin

install: build
	go install ./cmd/webrpc-gen

test: build
	go test -v ./...

generate:
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
