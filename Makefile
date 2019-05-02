help:
	@echo "make setup"
	@echo "make ganarate"

.PHONY: setup
setup:
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: generate
generate:
	protoc --go_out=plugins=grpc:./era/ proto/era.proto
	protoc --go_out=plugins=grpc:./suga/ proto/era.proto
