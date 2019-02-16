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

tools:
	GO111MODULE=off go get -u github.com/goware/statik

build:
	go generate ./gen/...
	go build -o ./bin/webrpc-gen ./cmd/webrpc-gen
	go generate ./...

install: build
	go install ./cmd/webrpc-gen

test: generate
	go test -v ./...

generate:
	go generate ./...

clean:
	rm -rf ./bin

dep:
	@export GO111MODULE=on && \
		go mod tidy && \
		rm -rf ./vendor && go mod vendor

dep-upgrade-all:
	@GO111MODULE=on go get -u
	@$(MAKE) dep

