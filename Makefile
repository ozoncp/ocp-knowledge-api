PROJECT_NAME=$(shell basename "$(PWD)")

.PHONY: run
run:
	go run cmd/$(PROJECT_NAME)/main.go

.PHONY: test
test:
	go test ./... -v

.PHONY: build
build: .vendor-proto .generate .build

.PHONY: .vendor-proto
.vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api/ocp-knowledge-api
	cp api/ocp-knowledge-api/ocp-knowledge-api.proto vendor.protogen/api/ocp-knowledge-api
	@if [ ! -d vendor.protogen/google ]; then \
  		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-knowledge-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-knowledge-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-knowledge-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-knowledge-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-knowledge-api/ocp-knowledge-api.proto
		mv pkg/ocp-knowledge-api/github.com/ozoncp/ocp-knowledge-api/pkg/ocp-knowledge-api/* pkg/ocp-knowledge-api/
		rm -rf pkg/ocp-knowledge-api/github.com
		cd pkg/ocp-knowledge-api && ls go.mod || go mod init github.com/ozoncp/ocp-knowledge-api/pkg/ocp-knowledge-api && go mod tidy

.PHONY: .build
.build:
		go build -o bin/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go

.PHONY: deps
	ls go.mod || go mod init github.com/ozoncp/ocp-knowledge-api
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger