#!/bin/sh
protoc --proto_path=api/proto --go_out=pkg/profileGRPC --go_opt=paths=source_relative --go-grpc_out=pkg/profileGRPC --go-grpc_opt=paths=source_relative api/proto/*.proto