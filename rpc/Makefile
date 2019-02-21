default: | help

build:  ## Build all of binaries
	go build


buildstatics:  ## Build all of binaries as static building
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"'

run:  ## Launch a grpc server
	realize start --server

test:  ## Test to all of directories
	go test -v -cover -race ./pb/... && \
	go test -v -cover -race ./handler/...

help:  ## Show all of tasks
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: \
	all \
	build \
	buildstatics \
	run \
	help \
	test
