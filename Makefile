default: | help

build:  ## Build all of binaries
	@for d in `echo rpc`; do \
		(cd $$d && make build) \
	; done

buildstatics:  ## Build all of binaries as static building
	@for d in `echo rpc`; do \
		(cd $$d && make buildstatics) \
	; done

gogo: ## Download gogoproto/gogo.proto
	go get "github.com/gogo/protobuf/gogoproto" && \
	mkdir -p ./proto/gogo/gogoproto && \
	cp "$$GOPATH/src/github.com/gogo/protobuf/gogoproto/gogo.proto" ./proto/gogo/gogoproto/

go-proto-validators: ## Download mwitkow/go-proto-validators
	go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators && \
	mkdir -p ./proto/pb/github.com/mwitkow/go-proto-validators && \
	cp "$$GOPATH/src/github.com/mwitkow/go-proto-validators/validator.proto" ./proto/pb/github.com/mwitkow/go-proto-validators/

install: | gogo go-proto-validators ## Install all of dependencies
	go get -v github.com/golang/protobuf/proto && \
	go get -v github.com/golang/protobuf/protoc-gen-go && \
	go get -v google.golang.org/grpc && \
	go get -v google.golang.org/genproto/... && \
	go get -v github.com/gogo/protobuf/... && \
	go get -v go.pedge.io/protoeasy/cmd/protoeasy && \
	go get -v github.com/mwitkow/go-proto-validators/protoc-gen-govalidators && \
	go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
	go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
	go get -v github.com/golang/mock/gomock && \
	go install github.com/golang/mock/mockgen && \
	go get -v -t github.com/volatiletech/sqlboiler && \
	go get -v github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql && \
	go get -v github.com/oxequa/realize && \
	go get -v github.com/alecthomas/gometalinter && \
	\
	gometalinter --install --update && \
	\
	go get -v github.com/golang/dep/cmd/dep && \
	dep ensure -v

protocol: ## Generate Protocol buffers definition
	protoeasy \
		--gogo --grpc --grpc-gateway --out ./rpc/pb proto/pb && \
	\
	protoc \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		-I${GOPATH}/src/go.pedge.io/protoeasy/vendor/go.pedge.io/pb/proto \
		-I./proto/pb/apicube \
		--govalidators_out=\
	Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
	Mgoogle/rpc/status.proto=go.pedge.io/pb/gogo/google/rpc,\
	Mgoogle/rpc/code.proto=go.pedge.io/pb/gogo/google/rpc,\
	Mgoogle/rpc/error_details.proto=go.pedge.io/pb/gogo/google/rpc,\
	plugins=gogoimport=true:rpc/pb/apicube \
		proto/pb/apicube/product/product.proto && \
	\
	go tool fix -force context rpc/pb && \
	gofmt -l -w rpc/pb && \
	\
	mkdir -p ./rpc/mockpb/mock_product && \
	mockgen github.com/ikeikeikeike/apicube/rpc/pb/apicube/product \
		ProductServiceServer,ProductServiceClient \
			> ./rpc/mockpb/mock_product/product.go && \
	mkdir -p ./rpc/mockpb/mock_lifecycle && \
	mockgen github.com/ikeikeikeike/apicube/rpc/pb/apicube/lifecycle \
		PingServiceServer,PingServiceClient \
			> ./rpc/mockpb/mock_lifecycle/ping.go && \
	\
	go test -v -cover -race ./rpc/pb/... && \
	go test -v -cover -race ./rpc/handler/... && \
	dep ensure -v

deps:  ## Install Golang dependencies
	dep ensure -v

modelgen:  ## Generate Golang model definition
	sqlboiler mysql

runner:  ## Launch a build runner
	realize start

travis:
	go get -v github.com/alecthomas/gometalinter && \
	gometalinter --install --update && \
	go get -v github.com/golang/dep/cmd/dep && \
	dep ensure -v

test:  ## Test to all of directories
	@for d in `echo rpc`; do \
		(cd $$d && make test) \
	; done

test-cover-html: ## Generate test coverage report
	go test -coverprofile=coverage.out -covermode=count
	go tool cover -func=coverage.out

fmt:  ## Run gofmt linter
	@for d in `go list ./... | grep -v 'migrations'` ; do \
		if [ "`gofmt -l -s $$GOPATH/src/$$d | tee /dev/stderr`" ]; then \
			echo "^ improperly formatted go files" && echo && exit 1; \
		fi \
	done

lint:  ## Run golint linter
	@for d in `go list ./... | grep -v 'migrations'` ; do \
		if [ "`golint $$d | tee /dev/stderr`" ]; then \
			echo "^ golint errors!" && echo && exit 1; \
		fi \
	done

vet:  ## Run go vet linter
	@for d in `go list ./... | grep -v 'migrations'` ; do \
		if [ "`go vet $$d | tee /dev/stderr`" ]; then \
			echo "^ go vet errors!" && echo && exit 1; \
		fi \
	done

misspell:  ## Check misspell
	@if [ "`find . -type f | \grep -v 'vendor/\|\.git' | xargs misspell -error | tee /dev/stderr`" ]; then \
		echo "^ misspell errors!" && echo && exit 1; \
	fi \

check:  ## Golang all of style checking
	@if [ "`gometalinter --fast --vendor --deadline=100s --exclude base/data/model --exclude rpc/pb ./... | tee /dev/stderr`" ]; then \
		echo "^ gometalinter errors!" && echo && exit 1; \
	fi

check-completely:  ## Golang completely all of style checking
	@if [ "`gometalinter --vendor --deadline=100s --exclude base/data/model --exclude rpc/pb ./... | tee /dev/stderr`" ]; then \
		echo "^ gometalinter completely check errors!" && echo && exit 1; \
	fi

reindex_products:  ## Reindex products
	go run cmd/loades/*.go -idxname products reindex

reindex: | reindex_products  ## Reindex all of task

help:  ## Show all of tasks
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: \
	help \
	build \
	buildstatics \
	test \
	deps \
	fmt \
	lint \
	vet \
	test-cover-html \
	install \
	protocol \
	runner \
	travis \
	misspell \
	check \
	check-completely \
	gogo \
	go-proto-validators \
	reindex \
	reindex_products \
	modelgen
