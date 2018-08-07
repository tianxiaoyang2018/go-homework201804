#!/bin/sh

# putong-domain-tools
# Generate go code from schemas
# Golang gRPC install guide: https://grpc.io/docs/quickstart/go.html

cd $( dirname "${BASH_SOURCE[0]}")


protoc --proto_path=${DIR} user/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} relationship/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} device/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} chat/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} contact/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} moment/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} feedback/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} membership/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} location/*.proto --go_out=plugins=grpc:.

protoc --proto_path=${DIR} common/*.proto --go_out=plugins=grpc:.
protoc --proto_path=${DIR} event/*.proto --go_out=plugins=grpc:.

cp -r github.com/p1cn/tantan-domain-schema/golang/* ../golang/

rm -r github.com
