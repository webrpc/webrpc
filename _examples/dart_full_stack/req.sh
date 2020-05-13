#!/bin/bash

curl -d "{\"user\":{\"id\":0,\"USERNAME\":\"amigo12\",\"role\":\"Kind.user()\"}}" -X POST -H 'Content-Type: application/json' http://localhost:8080/rpc/ExampleService/AddUser
