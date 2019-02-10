all:
	@echo "please read Makefile source to see available commands"

generate:
	@echo "generating es5 javascript client... or use 'make tsc-es6' if you want es6 output.." 
	@$(MAKE) tsc-es5

tsc-es5:
	tsc --target es5 --lib es5,es6,dom example.gen.ts

tsc-es6:
	tsc --target es6 --lib es6,dom example.gen.ts

start-go-server:
	go run .

run-node-client:
	node main.js
