#!/bin/bash

.PHONY:
gen: clean
	@ echo "generating protobuf"
	@ protoc --proto_path=proto proto/*.proto \
 	  --go_out=gen.pb.go --go-grpc_out=gen.pb.go \
 	  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

.PHONY:
fmt:
	@ echo "formatting go code"
	@ go fmt ./...

.PHONY:
server:
	@ echo "running server"
	@ go run cmd/server/main.go

.PHONY:
client:
	@ echo "running client"
	@ go run cmd/client/main.go

.PHONY:
clean:
	@ echo "cleaning-up"
	@ rm -rf gen.pb.go/*.pb.go