#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src"

protoc -I=$proto_imports --doc_out=docs/api  --doc_opt=markdown,topo.md --gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,import_path=topo,plugins=grpc:. api/topo/*.proto
