help:
	@echo "make setup"
	@echo "make ganarate"
	@echo "make run-era-service"
	@echo "make run-suga"

.PHONY: setup
setup:
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: generate
generate:
	protoc --go_out=plugins=grpc:./era/ proto/era.proto
	protoc --go_out=plugins=grpc:./suga/ proto/era.proto

.PHONY: run
run-era-service:
	go run ./era/main.go

.PHONY: run
run-suga:
	go run ./suga/main.go
