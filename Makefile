#!/bin/bash

.PHONY:
gen:
	@ echo "generating protobuf"
	@ protoc --proto_path=proto proto/*.proto --go_out=generated.pb --go_opt=paths=source_relative

.PHONY:
fmt:
	@ echo "formatting go code"
	@ go fmt ./...

.PHONY:
clean:
	@ rm -rf pb/*.go