all:
	@echo "please read Makefile source or README to see available commands"

bootstrap:
	rm -rf server/node_modules && rm -rf webapp/node_modules
	cd server && yarn
	cd webapp && yarn

generate:
	../../bin/webrpc-gen -schema=service.ridl -target=ts -server -out=./server/server.gen.ts
	../../bin/webrpc-gen -schema=service.ridl -target=ts -client -out=./webapp/client.gen.ts

run-server:
	yarn --cwd ./server start

run-client:
	yarn --cwd ./webapp start
