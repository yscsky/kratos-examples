HELLO_PROTO_FILES=$(shell find helloworld -name *.proto)
ERRORS_PROTO_FILES=$(shell find errors -name *.proto)
VAILD_PROTO_FILES=$(shell find validate -name *.proto)

.PHONY: hello
hello:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--go-http_out=paths=source_relative:. \
		$(HELLO_PROTO_FILES)

.PHONY: errors
errors:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-errors_out=paths=source_relative:. \
		$(ERRORS_PROTO_FILES)

.PHONY: validate
validate:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--go-http_out=paths=source_relative:. \
		--validate_out=paths=source_relative,lang=go:. \
		$(VAILD_PROTO_FILES)

help:
	@echo 'usage: hello, errors, validate'

.DEFAULT_GOAL := help
