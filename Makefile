.PHONY: protoc-cli-install
protoc-cli-install:
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
	npm i ts-proto -g

.PHONY: protoc-gen-go
protoc-gen-go:
	protoc --proto_path= ./proto-files/js/go/*.proto \
	--plugin=${HOME}/go/bin/protoc-gen-go \
	--go_out="${PWD}/go-microservices/internals/schemas"

.PHONY: protoc-gen-js
protoc-gen-js:
	protoc --proto_path= ./proto-files/js/*.proto \
	--plugin=${HOME}/.nvm/versions/node/v16.15.0/bin/protoc-gen-ts_proto \
	--ts_proto_out=${PWD}/node-gateways/src/schemas
	--ts_proto_opt=nestJs=true,addNestjsRestParameter=true,addGrpcMetadata=true