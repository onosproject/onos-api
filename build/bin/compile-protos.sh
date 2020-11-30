#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src":"${GOPATH}/src/github.com/onosproject/onos-api/api"

protoc -I=$proto_imports --doc_out=docs/api  --doc_opt=markdown,topo.md --gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,import_path=topo,plugins=grpc:. api/topo/*.proto

protoc -I=$proto_imports  --doc_out=docs/api  --doc_opt=markdown,headers.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=ricapi/e2,plugins=grpc:. api/ricapi/e2/headers/v1beta1/headers.proto
protoc -I=$proto_imports  --doc_out=docs/api  --doc_opt=markdown,ricapi.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=ricapi/e2,plugins=grpc:. api/ricapi/e2/v1beta1/ricapi_e2.proto

cp -r github.com/onosproject/onos-api/* .
rm -rf github.com
