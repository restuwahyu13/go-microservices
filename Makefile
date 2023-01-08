.PHONY: protoc-cli-install
protoc-cli-install:
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	npm i ts-proto -g

.PHONY: protoc-gen-go
protoc-gen-go:
	rm -rf ${PWD}/go-microservices/internals/schemas/**/*
	\
	protoc --proto_path= ./protofiles/*.proto \
	--plugin=${HOME}/go/bin/protoc-gen-go \
	--go_out="${PWD}/go-microservices/internals/schemas"
	\
	protoc --proto_path= ./protofiles/*.proto \
	--plugin=${HOME}/go/bin/protoc-gen-go-grpc \
	--go-grpc_out="${PWD}/go-microservices/internals/schemas"

.PHONY: protoc-gen-js
protoc-gen-js:
	protoc --proto_path= ./protofiles/*.proto \
	--plugin=${HOME}/.nvm/versions/node/v16.15.0/bin/protoc-gen-ts_proto \
	--ts_proto_out=${PWD}/gateways/src/schemas
	--ts_proto_opt=nestJs=true,addNestjsRestParameter=true,addGrpcMetadata=true

.PHONY: gateway
gateway:
	cd gateways && npm run start:dev

.PHONY: service
service:
ifdef type
	cd go-microservices && nodemon -V -e go -w ./go-microservices -i main -x go run ./domains/${type} --signal SIGTERM
endif