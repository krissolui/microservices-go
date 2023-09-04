#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative --go_grpc_out=. --go_grpc_opt=paths=source_relative logs.proto

cd ../
go get google.golang.org/grpc