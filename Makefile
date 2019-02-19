default: | help

build:  ## Build all of binaries
	@for d in `echo rpc`; do \
		(cd $$d && make build) \
	; done

buildstatics:  ## Build all of binaries as static building
	@for d in `echo rpc`; do \
		(cd $$d && make buildstatics) \
	; done

install: ## Install all of dependencies
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
	for d in `echo rpc`; do (cd $$d && make install) ; done && \
	\
	go get -v github.com/golang/dep/cmd/dep && \
	dep ensure -v

deps:  ## Install Golang dependencies
	dep ensure -v

modelgen:  ## Generate Golang model definition
	sqlboiler mysql

runner:  ## Launch a build runner
	realize start

protocol:  ## Generate Protocol buffers definition
	@for d in `echo rpc`; do \
		(cd $$d && make protocol) \
	; done && \
	\
	dep ensure -v

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

check:  ## Golang all of style checking
	@if [ "`gometalinter --fast --vendor --deadline=100s --exclude base/data/model --exclude rpc/protocol ./... | tee /dev/stderr`" ]; then \
		echo "^ gometalinter errors!" && echo && exit 1; \
	fi

check-completely:  ## Golang completely all of style checking
	@if [ "`gometalinter --vendor --deadline=100s --exclude base/data/model --exclude rpc/protocol ./... | tee /dev/stderr`" ]; then \
		echo "^ gometalinter completely check errors!" && echo && exit 1; \
	fi

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
	check \
	protocol \
	runner \
	travis \
	check-completely \
	modelgen
