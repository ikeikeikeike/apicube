default: | help

build:  ## Build all of binaries
	go build

gogo: ## Download gogoproto/gogo.proto
	go get "github.com/gogo/protobuf/gogoproto" && \
	mkdir -p ./apicube-proto/gogo/gogoproto && \
	cp "$$GOPATH/src/github.com/gogo/protobuf/gogoproto/gogo.proto" ./apicube-proto/gogo/gogoproto/

go-proto-validators: ## Download mwitkow/go-proto-validators
	go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators && \
	mkdir -p ./apicube-proto/protocol/github.com/mwitkow/go-proto-validators && \
	cp "$$GOPATH/src/github.com/mwitkow/go-proto-validators/validator.proto" ./apicube-proto/protocol/github.com/mwitkow/go-proto-validators/

install: | gogo go-proto-validators  ## Install all of dependencies

buildstatics:  ## Build all of binaries as static building
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"'

protocol: ## Generate Protocol buffers definition
	protoeasy \
		--gogo --grpc --grpc-gateway --out ./protocol apicube-proto/protocol && \
	\
	protoc \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		-I${GOPATH}/src/go.pedge.io/protoeasy/vendor/go.pedge.io/pb/proto \
		-I./apicube-proto/protocol/apicube \
		--govalidators_out=\
	Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
	Mgoogle/rpc/status.proto=go.pedge.io/pb/gogo/google/rpc,\
	Mgoogle/rpc/code.proto=go.pedge.io/pb/gogo/google/rpc,\
	Mgoogle/rpc/error_details.proto=go.pedge.io/pb/gogo/google/rpc,\
	plugins=gogoimport=true:protocol/apicube \
		apicube-proto/protocol/apicube/product/product.proto && \
	\
	go tool fix -force context protocol && \
	gofmt -l -w protocol && \
	\
	mkdir -p ./mockprotocol/mock_product && \
	mockgen github.com/ikeikeikeike/apicube/rpc/protocol/apicube/product \
		ProductServiceServer,ProductServiceClient \
			> mockprotocol/mock_product/product.go && \
	mkdir -p ./mockprotocol/mock_lifecycle && \
	mockgen github.com/ikeikeikeike/apicube/rpc/protocol/apicube/lifecycle \
		PingServiceServer,PingServiceClient \
			> mockprotocol/mock_lifecycle/ping.go && \
	\
	go test -v -cover -race ./protocol/... && \
	go test -v -cover -race ./handler/...

run:  ## Launch a grpc server
	realize start --server

test:  ## Test to all of directories
	go test -v -cover -race ./protocol/... && \
	go test -v -cover -race ./handler/...

help:  ## Show all of tasks
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: \
	all \
	build \
	buildstatics \
	run \
	help \
	test \
	validator \
	install \
	gogo \
	go-proto-validators \
	protocol
