.PHONY: build

ONOS_BUILD_VERSION := v0.6.7
ONOS_PROTOC_VERSION := v0.6.7
BUF_VERSION := 0.27.1

all: protos golang

golang: # @HELP compile Golang sources
	cd go && go build ./...

test: # @HELP run the unit tests and source code validation
test: protos golang license_check linters
	cd go && go test -race github.com/onosproject/onos-api/go/...

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

jenkins-test: build-tools jenkins-tools deps test # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
	export TEST_PACKAGES=github.com/onosproject/onos-api/go/... && cd go && ./../../build-tools/build/jenkins/make-unit
	mv go/*.xml .

deps: # @HELP ensure that the required dependencies are in place
	cd go && go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: golang-ci # @HELP examines Go source code and reports coding problems
	cd go && golangci-lint run --timeout 5m

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.36.0

license_check: build-tools # @HELP examine and ensure license headers exist
	./../build-tools/licensing/boilerplate.py -v --rootdir=/go/src/github.com/onosproject/onos-api/proto

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-api \
		-w /go/src/github.com/onosproject/onos-api \
		bufbuild/buf:${BUF_VERSION} check lint

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos:
	docker run -it \
	    -v `pwd`:/onos-api \
		-w /onos-api \
		--entrypoint build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

mocks:
	./build/bin/generate-mocks.sh

publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION}
	./../build-tools/publish-version go/${VERSION}

clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
