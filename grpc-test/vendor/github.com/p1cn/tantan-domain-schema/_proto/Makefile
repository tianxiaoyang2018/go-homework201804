# Disable GNU make built-in rules for building .c files from .y and .l files
%.c: %.y
%.c: %.l

all: dcl
	cp -r github.com/p1cn/tantan-domain-schema/golang/* ../golang/

domain:
	protoc --proto_path=${CURDIR} service/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} domain/*.proto --go_out=plugins=grpc:.

user: domain
	protoc --proto_path=${CURDIR} user/domain/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} user/service/*.proto --go_out=plugins=grpc:.

device: domain
	protoc --proto_path=${CURDIR} device/domain/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} device/service/*.proto --go_out=plugins=grpc:.

chat: domain
	protoc --proto_path=${CURDIR} chat/domain/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} chat/service/*.proto --go_out=plugins=grpc:.

moment: domain
	protoc --proto_path=${CURDIR} moment/domain/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} moment/service/*.proto --go_out=plugins=grpc:.

feedback: domain
	protoc --proto_path=${CURDIR} feedback/domain/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} feedback/service/*.proto --go_out=plugins=grpc:.

location: domain
	protoc --proto_path=${CURDIR} location/domain/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} location/service/*.proto --go_out=plugins=grpc:.

dcl: user device chat moment feedback
	protoc --proto_path=${CURDIR} dcl/*.proto --go_out=plugins=grpc:.

relationship: domain
	protoc --proto_path=${CURDIR} relationship/domain/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=${CURDIR} relationship/service/*.proto --go_out=plugins=grpc:.

clean:
	rm -fr github.com

.PHONY: domain user device chat moment feedback location dcl clean relationship
