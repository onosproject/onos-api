.PHONY: build

ONOS_PROTOC_VERSION := v0.6.9
BUF_VERSION := 0.47.0

all: protos golang

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

golang: # @HELP compile Golang sources
	cd go && go build ./...

test: # @HELP run the unit tests and source code validation
test: protos golang license_check-proto linters-go deps-go
	cd go && go test -race github.com/onosproject/onos-api/go/...

jenkins-test: # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: jenkins-tools test deps-go
	export TEST_PACKAGES=github.com/onosproject/onos-api/go/... && cd go && ../build/build-tools/build/jenkins/make-unit
	mv go/*.xml .

deps-go: # @HELP ensure that the required dependencies are in place
	cd go && go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go/go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go/go.sum)"

linters-go: golang-ci # @HELP examines Go source code and reports coding problems
	cd go && golangci-lint run --timeout 15m

twine: # @HELP install twine if not present
	twine --version || pip install twine

license_check-proto: # @HELP examine and ensure license headers exist
	./build/build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}/proto

buflint: #@HELP run the "buf check lint" command on the proto files in 'api'
	docker run -v `pwd`:/go/src/github.com/onosproject/onos-api \
		-w /go/src/github.com/onosproject/onos-api \
		bufbuild/buf:${BUF_VERSION} check lint

protos: # @HELP compile the protobuf files (using protoc-go Docker)
protos:
	docker run \
	    -v `pwd`:/onos-api \
		-w /onos-api \
		--entrypoint build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

mocks:
	./build/bin/generate-mocks.sh

publish: twine # @HELP publish version on github, dockerhub, abd PyPI
	BASEDIR=. PYPI_INDEX=pypi ./build/build-tools/publish-python-version
	./build/build-tools/publish-version ${VERSION}
	./build/build-tools/publish-version go/${VERSION}

jenkins-publish: jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/build-tools/release-merge-commit

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor

