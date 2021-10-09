PROTO_FILES=$(shell find helloworld -name *.proto)

.PHONY: proto
proto:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--go-http_out=paths=source_relative:. \
		$(PROTO_FILES)

help:
	@echo 'usage: proto'

.DEFAULT_GOAL := help
