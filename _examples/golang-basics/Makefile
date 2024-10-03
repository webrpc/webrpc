all:
	@awk -F'[ :]' '!/^all:/ && /^([A-z_-]+):/ {print "make " $$1}' Makefile

generate:
	webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -service="ExampleService" -out=./example.gen.go
	webrpc-gen -schema=example.ridl -target=golang -pkg=admin -server -client -service="AdminService" -ignore="@deprecated" -out=./admin/admin.gen.go
	webrpc-gen -schema=example.ridl -target=golang -pkg=internal -server -client -service="ExampleService,AdminService" -ignore="@public,@deprecated" -out=./internal/internal.gen.go

dev-generate:
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -service="ExampleService" -out=./example.gen.go
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=admin -server -client -service="AdminService" -ignore="@deprecated" -out=./admin/admin.gen.go
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=internal -server -client -service="ExampleService,AdminService" -ignore="@public,@deprecated" -out=./internal/internal.gen.go

dev-generate-local-templates:
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=main -server -client -service="ExampleService" -ignore="@deprecated" -out=./example.gen.go
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=admin -server -client -service="AdminService" -ignore="@deprecated" -out=./admin/admin.gen.go
	../../bin/webrpc-gen -schema=example.ridl -target=golang -pkg=internal -server -client -service="ExampleService,AdminService" -ignore="@public,@deprecated" -out=./internal/internal.gen.go
